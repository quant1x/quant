package main

import (
	"flag"
	"fmt"
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/category"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/progressbar"
	"github.com/mymmsc/gox/util/lambda"
	"github.com/mymmsc/gox/util/treemap"
	tableView "github.com/olekukonko/tablewriter"
	"github.com/quant1x/quant/internal"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"runtime"
	"sync"
	"time"
)

// Strategy 策略/公式指标(features)接口
type Strategy interface {
	// Name 策略名称
	Name() string
	// Code 策略编号
	Code() int
	// Evaluate 评估 日线数据
	Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map)
}

const (
	batchMax = quotes.TDX_SECURITY_QUOTES_MAX // 批量最大80条
	TopN     = 10
	kUnknown = "unknown"
)

var (
	MinVersion string
)

// 策略入口
func main() {
	var (
		//path     string // 数据路径
		strategy int  // 策略编号
		avx2     bool // AVX2加速状态
		cpuNum   int  // cpu数量
		version  bool // 显示版本号
	)
	flag.IntVar(&strategy, "strategy", 1, "strategy serial number")
	flag.BoolVar(&avx2, "avx2", false, "Avx2 acceleration")
	flag.IntVar(&cpuNum, "cpu", runtime.NumCPU()/2, "sets the maximum number of CPUs")
	flag.BoolVar(&version, "version", false, "print quant version")
	flag.Parse()

	if version {
		fmt.Println(MinVersion)
		os.Exit(0)
	}

	var api Strategy
	switch strategy {
	case 89:
		api = new(FormulaNo89)
	case 85:
		api = new(FormulaNo85)
	case 84:
		api = new(FormulaNo84)
	case 3:
		api = new(FormulaNo3)
	default:
		api = new(FormulaNo1)
	}
	stat.SetAvx2Enabled(avx2)
	runtime.GOMAXPROCS(cpuNum)
	var wg = sync.WaitGroup{}
	fmt.Printf("Quant1X 预警系统 %s: %s\n", MinVersion, api.Name())
	infos, _ := cpu.Info()
	cpuInfo := infos[0]
	memory, _ := mem.VirtualMemory()
	fmt.Printf("CPU: %s %dCores, AVX2: %t, Mem: total %dGB, free %dGB\n", cpuInfo.ModelName, cpuInfo.Cores, stat.GetAvx2Enabled(), memory.Total/(1024*1024*1024), memory.Free/(1024*1024*1024))

	// progress bar index
	pbarIndex := 0

	// 执行板块指数的检测
	dfBlock := stock.BlockList()
	blockCodes := dfBlock.Col("code").Strings()
	blockCount := len(blockCodes)
	bar := progressbar.NewBar(pbarIndex, "执行[扫描板块指数]", blockCount)
	pbarIndex++
	snapshots := []internal.QuoteSnapshot{}
	mapBlockName := make(map[string]string)
	for start := 0; start < blockCount; start += batchMax {
		codes := []string{}
		length := blockCount - start
		if length >= batchMax {
			length = batchMax
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
		snapshots = append(snapshots, shots...)
	}

	blocks := lambda.LambdaArray(snapshots).Sort(func(a internal.QuoteSnapshot, b internal.QuoteSnapshot) bool {
		aZf := a.Price / a.LastClose
		bZf := b.Price / b.LastClose
		aSpeed := a.Rate
		bSpeed := b.Rate
		zf := aZf > bZf
		speed := aSpeed > bSpeed
		return zf && speed || zf
	}).Pointer().([]internal.QuoteSnapshot)
	// 涨幅榜前N
	mapBlockData := make(map[string]internal.QuoteSnapshot)
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		bc := blocks[i].Code
		bn := mapBlockName[bc]
		block.Name = bn
		mapBlockData[bc] = block
	}
	//ssdf := pandas.LoadStructs(blocks[:TopN])
	ssdf := pandas.LoadStructs(blocks)
	ssdf = ssdf.Select([]string{"Code", "Price", "LastClose", "Rate"})
	codes := ssdf.Col("Code").Strings()
	names := []string{}
	topCodes := []string{}
	topNames := []string{}
	topRates := []float64{}
	// 板块代码映射板块数据
	// 个股到板块的映射
	stock2Block := make(map[string][]string)
	stock2Rank := make(map[string]internal.QuoteSnapshot)
	blockCount = len(codes)
	fmt.Println("")
	bar = progressbar.NewBar(pbarIndex, "执行[板块个股涨幅扫描]", blockCount)
	pbarIndex++
	for _, blockCode := range codes {
		bar.Add(1)
		name, ok := mapBlockName[blockCode]
		if !ok {
			name = kUnknown
		}
		names = append(names, name)
		fn := cache.BlockFilename(blockCode[2:])
		dfStock := pandas.ReadCSV(fn)
		stockCount := dfStock.Nrow()
		if stockCount == 0 {
			continue
		}
		topCode := kUnknown
		topName := kUnknown
		topRate := float64(0.00)
		stockCodes := dfStock.Col("code").Strings()
		stockSnapshots := []internal.QuoteSnapshot{}
		for start := 0; start < stockCount; start += batchMax {
			codes := []string{}
			length := stockCount - start
			if length >= batchMax {
				length = batchMax
			}
			for i := 0; i < length; i++ {
				code := stockCodes[start+i]
				_, mname, mcode := category.DetectMarket(code)
				code = security.GetStockCode(mname, mcode)
				basicInfo, err := security.GetBasicInfo(code)
				if err == security.ErrCodeNotExist {
					// 证券代码不存在
					continue
				}
				if err != nil {
					// 其它错误 输出错误信息
					logger.Errorf("%s => %v", code, err)
					continue
				}
				if basicInfo.Delisting {
					logger.Errorf("%s => 已退市", code)
					continue
				}
				tmpBlocks, _ := stock2Block[code]
				if len(tmpBlocks) == 0 {
					tmpBlocks = make([]string, 0)
				}
				tmpBlocks = append(tmpBlocks, blockCode)
				stock2Block[code] = tmpBlocks
				codes = append(codes, code)
				mapBlockName[code] = basicInfo.Name
			}
			if len(codes) == 0 {
				continue
			}
			stockShots := internal.BatchSnapShot(codes)
			if len(stockShots) == 0 {
				continue
			}
			stockSnapshots = append(stockSnapshots, stockShots...)
		}
		if len(stockSnapshots) == 0 {
			continue
		}
		tops := lambda.LambdaArray(stockSnapshots).Sort(func(a internal.QuoteSnapshot, b internal.QuoteSnapshot) bool {
			aZf := a.Price / a.LastClose
			bZf := b.Price / b.LastClose
			aSpeed := a.Rate
			bSpeed := b.Rate
			zf := aZf > bZf
			speed := aSpeed > bSpeed
			return zf && speed || zf
		}).Pointer().([]internal.QuoteSnapshot)
		topCode = tops[0].Code
		info, err := security.GetBasicInfo(topCode)
		if err == nil {
			topName = info.Name
		}
		topRate = (tops[0].Price/tops[0].LastClose - 1.00) * 100.00
		topCodes = append(topCodes, topCode)
		topNames = append(topNames, topName)
		topRates = append(topRates, topRate)
		total := 0
		limits := 0
		ling := 0
		up := 0
		down := 0
		for j := 0; j < len(tops); j++ {
			gp := tops[j]
			total += 1
			zf := gp.Price/gp.LastClose - 1.00
			zfLimit := category.MarketLimit(gp.Code)
			if zf >= zfLimit {
				limits += 1
			}
			if zf > 0.0000 {
				up++
			} else if zf < 0.0000 {
				down++
			} else {
				ling += 1
			}
			gp.TopNo = j
			_, ok := stock2Rank[gp.Code]
			if !ok {
				stock2Rank[gp.Code] = gp
			}
		}
		for j, v := range blocks {
			if v.Code == blockCode {
				blocks[j].Name = name
				blocks[j].TopCode = topCode
				blocks[j].TopName = topName
				blocks[j].TopRate = topRate
				blocks[j].TopNo = j
				blocks[j].Count = total
				blocks[j].ZhanTing = limits
				blocks[j].BidVol1 = up
				blocks[j].Ling = ling
				blocks[j].AskVol1 = down
				mapBlockData[blockCode] = blocks[j]
			}
		}
	}
	//oc := ssdf.Col("Code")
	//op := ssdf.Col("Price")
	//olc := ssdf.Col("LastClose")
	//or := ssdf.Col("Rate")
	//on := pandas.NewSeries(stat.SERIES_TYPE_STRING, "Name", names)
	//tmpZf := ssdf.Col("Price").Div(ssdf.Col("LastClose")).Sub(1.00)
	//ozf := pandas.NewSeries(stat.SERIES_TYPE_FLOAT64, "ZhangFu", tmpZf)
	//osc := pandas.NewSeries(stat.SERIES_TYPE_STRING, "Top1Code", topCodes)
	//osn := pandas.NewSeries(stat.SERIES_TYPE_STRING, "Top1Name", topNames)
	//osr := pandas.NewSeries(stat.SERIES_TYPE_FLOAT64, "Top1ZF", topRates)
	//ssdf = pandas.NewDataFrame(oc, on, op, olc, ozf, or, osc, osn, osr)
	//fmt.Println(ssdf)
	// 执行策略
	// 获取全部证券代码
	ss := stock.GetCodeList()
	count := len(ss)
	fmt.Println("")
	bar = progressbar.NewBar(pbarIndex, "执行["+api.Name()+"]", count)
	pbarIndex++
	var mapStock *treemap.Map
	mapStock = treemap.NewWithStringComparator()
	mainStart := time.Now()
	for i, v := range ss {
		fullCode := v
		basicInfo, err := security.GetBasicInfo(fullCode)
		if err == security.ErrCodeNotExist {
			// 证券代码不存在
			bar.Add(1)
			continue
		}
		if err != nil {
			// 其它错误 输出错误信息
			bar.Add(1)
			continue
		}
		bar.Add(1)
		wg.Add(1)
		go evaluate(api, &wg, fullCode, basicInfo, mapStock)
		_ = i
	}
	wg.Wait()
	fmt.Println("\n ======= [" + api.Name() + "] progress bar completed ==========\n")
	table := tableView.NewWriter(os.Stdout)
	var row ResultInfo
	table.SetHeader(row.Headers())

	elapsedTime := time.Since(mainStart) / time.Millisecond
	goals := mapStock.Size()
	fmt.Printf("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 命中: %d, 平均: %.3f/s\n", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, goals, float64(count)/(float64(elapsedTime)/1000))
	logger.Infof("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 命中: %d, 平均: %.3f/s", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, goals, float64(count)/(float64(elapsedTime)/1000))
	rs := make([]ResultInfo, 0)
	fmt.Println("")
	// 执行曲线回归
	bar = progressbar.NewBar(pbarIndex, "执行[曲线回归策略]", goals)
	pbarIndex++
	mapStock.Each(func(key interface{}, value interface{}) {
		bar.Add(1)
		row := value.(ResultInfo)
		sc := row.Code
		bs, ok := stock2Block[sc]
		if ok {
			btop := bs[0]
			shot, ok1 := mapBlockData[btop]
			if ok1 {
				//row.BlockCode = shot.Code
				row.BlockName = shot.Name
				row.BlockRate = (shot.Price/shot.LastClose - 1.00) * 100
				//row.BlockTopCode = shot.TopCode
				row.BlockTop = shot.TopNo
				row.BlockZhangTing = fmt.Sprintf("%d/%d", shot.ZhanTing, shot.Count)
				row.BlockDescribe = fmt.Sprintf("%d/%d/%d", shot.BidVol1, shot.AskVol1, shot.Ling)
				row.BlockTopName = shot.TopName
				row.BlockTopRate = shot.TopRate
			}
			shot, ok1 = stock2Rank[sc]
			if ok1 {
				row.BlockRank = shot.TopNo
			}
		}
		row.Predict()
		rs = append(rs, row)
		table.Append(row.Values())
	})
	output(api.Code(), rs)
	table.Render() // Send output

	// 过滤未有效突破的信号
	table = tableView.NewWriter(os.Stdout)
	count = mapStock.Size()
	fmt.Println("")
	// 执行检测趋势突破
	//bar = progressbar.NewBar(2, "执行[检测趋势突破]", count)
	//rsCross := make([]ResultInfo, 0)
	//mainStart = time.Now()
	//for _, v := range rs {
	//	bar.Add(1)
	//	if v.Cross() {
	//		rsCross = append(rsCross, v)
	//		table.Append(v.Values())
	//	}
	//}

	// 执行检测能量转强
	fmt.Println("")
	bar = progressbar.NewBar(pbarIndex, "执行[检测能量转强]", count)
	pbarIndex++
	rsValue := make([]ResultInfo, 0)
	mainStart = time.Now()
	for _, v := range rs {
		bar.Add(1)
		if v.DetectVolume() {
			rsValue = append(rsValue, v)
			table.Append(v.Values())
		}
	}
	fmt.Println("")
	goals = len(rsValue)
	fmt.Printf("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 命中: %d, 平均: %.3f/s\n", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, goals, float64(count)/(float64(elapsedTime)/1000))
	table = tableView.NewWriter(os.Stdout)
	count = len(rsValue)

	// 执行置信区间检测
	fmt.Println("")
	bar = progressbar.NewBar(pbarIndex, "执行[置信区间检测]", count)
	pbarIndex++
	rsCib := make([]ResultInfo, 0)
	mainStart = time.Now()
	for _, v := range rsValue {
		bar.Add(1)
		if v.Sample() {
			rsCib = append(rsCib, v)
			table.Append(v.Values())
		}
	}
	elapsedTime = time.Since(mainStart) / time.Millisecond
	goals = len(rsCib)
	fmt.Println("")
	fmt.Printf("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 命中: %d, 平均: %.3f/s\n", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, goals, float64(count)/(float64(elapsedTime)/1000))
	table.SetHeader(row.Headers())
	table.Render()
	output(api.Code()+10000, rsCib)
}

func evaluate(api Strategy, wg *sync.WaitGroup, code string, info *security.StaticBasic, result *treemap.Map) {
	defer wg.Done()
	api.Evaluate(code, info, result)
}

func output(strategyNo int, v []ResultInfo) {
	df := pandas.LoadStructs(v)
	filename := fmt.Sprintf("%s/%s/%s-%d.csv", cache.CACHE_ROOT_PATH, CACHE_STRATEGY_PATH, cache.Today(), strategyNo)
	_ = cache.CheckFilepath(filename)
	_ = df.WriteCSV(filename)

}
