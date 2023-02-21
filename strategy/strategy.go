package main

import (
	"flag"
	"fmt"
	"gitee.com/quant1x/data/category"
	"gitee.com/quant1x/data/security"
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

func main() {
	var (
		//path     string // 数据路径
		strategy int  // 策略编号
		avx2     bool // AVX2加速状态
		cpuNum   int  // cpu数量
	)
	//flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.IntVar(&strategy, "strategy", 1, "strategy serial number")
	flag.BoolVar(&avx2, "avx2", false, "Avx2 acceleration")
	flag.IntVar(&cpuNum, "cpu", runtime.NumCPU()/2, "sets the maximum number of CPUs")
	flag.Parse()
	//cache.Init(path)
	var api Strategy
	switch strategy {
	case 89:
		api = new(FormulaNo89)
	default:
		api = new(FormulaNo1)
	}
	stat.SetAvx2Enabled(avx2)
	//numCPU := runtime.NumCPU() / 2
	runtime.GOMAXPROCS(cpuNum)
	// 获取全部证券代码
	ss := category.GetCodeList()
	count := len(ss)
	var wg = sync.WaitGroup{}
	fmt.Println("Quant1X 预警系统: " + api.Name())
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
	elapsedTime := time.Since(mainStart) / time.Millisecond
	fmt.Printf("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 平均: %.3f/s\n", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, float64(count)/(float64(elapsedTime)/1000))
	logger.Infof("CPU: %d, AVX2: %t, 总耗时: %.3fs, 总记录: %d, 平均: %.3f/s", cpuNum, stat.GetAvx2Enabled(), float64(elapsedTime)/1000, count, float64(count)/(float64(elapsedTime)/1000))
	table := tableView.NewWriter(os.Stdout)
	var row ResultInfo
	table.SetHeader(row.Headers())

	mapStock.Each(func(key interface{}, value interface{}) {
		row := value.(ResultInfo)
		table.Append(row.Values())
	})
	table.Render() // Send output
}

func evaluate(api Strategy, wg *sync.WaitGroup, code string, info *security.StaticBasic, result *treemap.Map) {
	defer wg.Done()
	wg.Add(1)
	api.Evaluate(code, info, result)
}
