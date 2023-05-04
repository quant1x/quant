package models

import (
	"fmt"
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/progressbar"
	"github.com/quant1x/quant/labs/trade"
	"sort"
)

// 板块常量
const (
	BlockTopN = 3 // 板块排行前几名
	StockTopN = 3 // 板块个股前几名
)

// BlockInfo 板块信息
type BlockInfo struct {
	BlockCode  string   // 板块代码
	BlockName  string   // 板块名称
	BlockType  string   // 板块类型
	ZhangDieFu float64  // 板块涨幅
	BlockTop   int      // 板块排名
	TopCode    string   // 领涨个股
	TopName    string   // 领涨个股名称
	TopRate    float64  // 领涨个股涨幅
	ZhanTing   int      // 涨停数
	Ling       int      // 平盘数
	Count      int      // 总数
	Up         int      // 上涨家数
	Down       int      // 下定家数
	LiuTongPan float64  // 流通盘
	FreeGuBen  float64  // 自由流通股本
	TurnZ      float64  // 开盘换手
	StockCodes []string `dataframe:"-"` // 股票代码

}

// 板块排序
func BlockSort(a, b trade.QuoteSnapshot) bool {
	if a.ZhangDieFu > b.ZhangDieFu {
		return true
	}
	if a.ZhangDieFu == b.ZhangDieFu && a.Amount > b.Amount {
		return true
	}
	if a.ZhangDieFu == b.ZhangDieFu && a.Amount == b.Amount && a.TurnZ > b.TurnZ {
		return true
	}
	return false
}

// 个股排序
func StockSort(a, b trade.QuoteSnapshot) bool {
	if a.ZhangDieFu > b.ZhangDieFu {
		return true
	}
	if a.ZhangDieFu == b.ZhangDieFu && a.TurnZ > b.TurnZ {
		return true
	}
	return false
}

func scanBlock(pbarIndex int, blockType security.BlockType) []trade.QuoteSnapshot {
	// 执行板块指数的检测
	dfBlock := stock.BlockList()
	var blockCodes []string
	for i := 0; i < dfBlock.Nrow(); i++ {
		m := dfBlock.IndexOf(i)
		var bt security.BlockType
		v, ok := m["type"]
		if ok {
			bt = security.BlockType(stat.AnyToInt32(v))
		} else {
			continue
		}
		// 只保留行业和概念
		//if bt != security.BK_HANGYE && bt != security.BK_GAINIAN && bt != security.BK_YJHY {
		if bt != blockType {
			continue
		}
		var bc string
		v, ok = m["code"]
		if ok {
			bc = stat.AnyToString(v)
		} else {
			continue
		}
		blockCodes = append(blockCodes, bc)
	}

	blockCount := len(blockCodes)
	fmt.Println()
	btn, ok := security.BlockTypeNameByTypeCode(blockType)
	if !ok {
		btn = stat.AnyToString(blockType)
	}
	bar := progressbar.NewBar(pbarIndex, "执行[扫描"+btn+"板块指数]", blockCount)
	pbarIndex++
	snapshots := []trade.QuoteSnapshot{}
	mapBlockName := make(map[string]string)
	for start := 0; start < blockCount; start += quotes.TDX_SECURITY_QUOTES_MAX {
		codes := []string{}
		length := blockCount - start
		if length >= quotes.TDX_SECURITY_QUOTES_MAX {
			length = quotes.TDX_SECURITY_QUOTES_MAX
		}
		for i := 0; i < length; i++ {
			code := blockCodes[start+i]
			basicInfo, err := security.GetBasicInfo(code)
			if err == security.ErrCodeNotExist {
				// 证券代码不存在
				bar.Add(1)
				continue
			}
			if err != nil {
				// 其它错误 输出错误信息
				logger.Errorf("%s => %v", code, err)
				bar.Add(1)
				continue
			}
			if basicInfo.Delisting {
				logger.Errorf("%s => 已退市", code)
				bar.Add(1)
				continue
			}
			bar.Add(1)
			codes = append(codes, code)
			mapBlockName[code] = basicInfo.Name
		}
		logger.Infof("%+v", codes)
		if len(codes) == 0 {
			continue
		}
		shots := trade.BatchSnapShot(codes)
		if len(shots) == 0 {
			continue
		}
		for _, v := range shots {
			v.Name = mapBlockName[v.Code]
			v.LiuTongPan = cache.GetLiuTongPan(v.Code)
			v.FreeGuBen = cache.GetFreeGuBen(v.Code)
			kpVol := cache.GetKaipanVol(v.Code)
			kpVol = kpVol * 100
			v.TurnZ = kpVol / v.FreeGuBen * 100

			snapshots = append(snapshots, v)
		}
	}
	sort.Slice(snapshots, func(i, j int) bool {
		a := snapshots[i]
		b := snapshots[j]

		return BlockSort(a, b)
	})
	return snapshots
}

func getBlockByType(pbarIndex int, blockType security.BlockType) []BlockInfo {
	bs := []BlockInfo{}
	blocks := scanBlock(pbarIndex, blockType)
	// 涨幅榜前N
	top := 0
	for i := 0; i < len(blocks) && i < BlockTopN; i++ {
		v := blocks[i]
		// 获取板块内个股列表
		fn := cache.BlockFilename(v.Code[2:])
		dfStock := pandas.ReadCSV(fn)
		stockCount := dfStock.Nrow()
		if stockCount == 0 {
			continue
		}
		top++
		stockCodes := dfStock.Col("code").Strings()
		bi := BlockInfo{
			BlockCode:  v.Code,
			BlockName:  v.Name,
			BlockType:  BlockTypeName(v.Code),
			ZhangDieFu: v.ZhangDieFu,
			BlockTop:   top,
			StockCodes: stockCodes,
		}
		bs = append(bs, bi)
	}
	return bs
}

// TopBlock 板块排行
func TopBlock(pbarIndex int) []BlockInfo {
	bs := []BlockInfo{}
	blockTypes := []security.BlockType{security.BK_HANGYE, security.BK_GAINIAN}
	for _, bt := range blockTypes {
		pbarIndex += 1
		blocks := getBlockByType(pbarIndex, bt)
		bs = append(bs, blocks...)
	}
	fmt.Println()
	return bs
}

var (
	kMapBlockType = map[string]string{}
)

func init() {
	_ = GetBlockList()
}

func GetBlockList() []string {
	// 执行板块指数的检测
	dfBlock := stock.BlockList()
	var blockCodes []string
	for i := 0; i < dfBlock.Nrow(); i++ {
		m := dfBlock.IndexOf(i)
		var bt security.BlockType
		v, ok := m["type"]
		if ok {
			bt = security.BlockType(stat.AnyToInt32(v))
		} else {
			continue
		}
		// 只保留行业和概念
		//if bt != security.BK_HANGYE && bt != security.BK_GAINIAN && bt != security.BK_YJHY {
		if bt != security.BK_HANGYE && bt != security.BK_GAINIAN {
			continue
		}
		var bc string
		v, ok = m["code"]
		if ok {
			bc = stat.AnyToString(v)
		} else {
			continue
		}
		blockCodes = append(blockCodes, bc)
		it := int(bt)
		btn, _ := security.BlockTypeNameByCode(it)
		kMapBlockType[bc] = btn
	}

	return blockCodes
}

func BlockTypeName(blockCode string) string {
	name, _ := kMapBlockType[blockCode]
	return name
}
