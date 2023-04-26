package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/progressbar"
	"github.com/quant1x/quant/internal"
	"sort"
)

// 板块常量
const (
	BlockTopN = 5 // 板块排行前几名
	StockTopN = 3 // 板块个股前几名
)

// BlockInfo 板块信息
type BlockInfo struct {
	BlockCode  string   // 板块代码
	BlockName  string   // 板块名称
	BlockType  string   // 板块类型
	ZhangDieFu float64  // 板块涨幅
	StockCodes []string // 股票代码
}

// 板块排序
func blockSort(a, b internal.QuoteSnapshot) bool {
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
func stockSort(a, b internal.QuoteSnapshot) bool {
	if a.ZhangDieFu > b.ZhangDieFu {
		return true
	}
	if a.ZhangDieFu == b.ZhangDieFu && a.TurnZ > b.TurnZ {
		return true
	}
	return false
}

func scanBlock(pbarIndex int) []internal.QuoteSnapshot {
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
		if bt != security.BK_GAINIAN {
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
	bar := progressbar.NewBar(pbarIndex, "执行[扫描板块指数]", blockCount)
	pbarIndex++
	snapshots := []internal.QuoteSnapshot{}
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
		shots := internal.BatchSnapShot(codes)
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

		return blockSort(a, b)
	})

	return snapshots
}

// TopBlock 板块排行
func TopBlock(pbarIndex int) []BlockInfo {
	blocks := scanBlock(pbarIndex)
	// 涨幅榜前N
	bs := []BlockInfo{}
	for i := 0; i < len(blocks) && i < BlockTopN; i++ {
		v := blocks[i]
		bi := BlockInfo{
			BlockCode: v.Code,
			BlockName: v.Name,
		}
		bs = append(bs, bi)
	}
	return bs
}

var (
	kMapBlockType = map[string]string{}
)

func getBlockList() []string {
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

func blockTypeName(blockCode string) string {
	name, _ := kMapBlockType[blockCode]
	return name
}
