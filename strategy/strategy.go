package main

import (
	"flag"
	"fmt"
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/progressbar"
	"github.com/mymmsc/gox/util/treemap"
	tableView "github.com/olekukonko/tablewriter"
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

var (
	MinVersion string
)

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
	case 84:
		api = new(FormulaNo84)
	case 3:
		api = new(FormulaNo3)
	default:
		api = new(FormulaNo1)
	}
	stat.SetAvx2Enabled(avx2)
	runtime.GOMAXPROCS(cpuNum)
	// 获取全部证券代码
	ss := stock.GetCodeList()
	count := len(ss)
	var wg = sync.WaitGroup{}
	fmt.Printf("Quant1X 预警系统 %s: %s\n", MinVersion, api.Name())
	infos, _ := cpu.Info()
	cpuInfo := infos[0]
	memory, _ := mem.VirtualMemory()
	fmt.Printf("CPU: %s %dCores, AVX2: %t, Mem: total %dGB, free %dGB\n", cpuInfo.ModelName, cpuInfo.Cores, stat.GetAvx2Enabled(), memory.Total/(1024*1024*1024), memory.Free/(1024*1024*1024))
	bar := progressbar.NewBar(0, "执行["+api.Name()+"]", count)
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
	bar = progressbar.NewBar(1, "执行[曲线回归策略]", goals)
	mapStock.Each(func(key interface{}, value interface{}) {
		bar.Add(1)
		row := value.(ResultInfo)
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
	bar = progressbar.NewBar(2, "执行[检测能量转强]", count)
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
	fmt.Printf("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 命中: %d, 平均: %.3f/s\n", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, goals, float64(count)/(float64(elapsedTime)/1000))
	table = tableView.NewWriter(os.Stdout)
	count = len(rsValue)
	bar = progressbar.NewBar(3, "执行[置信区间检测]", count)
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
	wg.Add(1)
	api.Evaluate(code, info, result)
}

func output(strategyNo int, v []ResultInfo) {
	df := pandas.LoadStructs(v)
	filename := fmt.Sprintf("%s/%s/%s-%d.csv", cache.CACHE_ROOT_PATH, CACHE_STRATEGY_PATH, cache.Today(), strategyNo)
	_ = cache.CheckFilepath(filename)
	_ = df.WriteCSV(filename)

}
