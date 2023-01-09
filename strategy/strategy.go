package main

import (
	"flag"
	"fmt"
	"github.com/mymmsc/gox/util/treemap"
	termTable "github.com/olekukonko/tablewriter"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/utils/progressbar"
	"os"
	"runtime"
	"sync"
)

// FormulaApi 公式指标(features)接口
type FormulaApi interface {
	// Name 策略名称
	Name() string
	// Code 策略编号
	Code() int
	// Evaluate 评估 日线数据
	Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map)
}

func main() {
	var (
		path     string
		strategy int
	)
	flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.IntVar(&strategy, "strategy", 1, "strategy serial number")
	flag.Parse()
	cache.Init(path)
	var api FormulaApi
	switch strategy {
	case 89:
		api = new(FormulaNo89)
	default:
		api = new(FormulaNo1)
	}
	numCPU := runtime.NumCPU() / 2
	runtime.GOMAXPROCS(numCPU)
	// 获取全部证券代码
	ss := data.GetCodeList()
	count := len(ss)
	var wg = sync.WaitGroup{}
	doneCh := make(chan struct{})
	bar := progressbar.NewOptions(count,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(80),
		progressbar.OptionSetDescription("[cyan][1/3][reset]执行["+api.Name()+"]..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[red]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
			//SaucerPadding: "[white]•",
			//BarStart:      "[blue]|[reset]",
			//BarEnd:        "[blue]|[reset]",
		}),
		progressbar.OptionOnCompletion(func() {
			doneCh <- struct{}{}
		}),
	)
	//fmt.Printf("计划买入, 信号日期, 委托价格, 目标价位\n")
	var mapStock *treemap.Map
	mapStock = treemap.NewWithStringComparator()
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
		go evaluate(bar, api, &wg, fullCode, basicInfo, mapStock)
		_ = i

	}
	// got notified that progress bar is complete.
	<-doneCh
	wg.Wait()
	fmt.Println("\n ======= [" + api.Name() + "] progress bar completed ==========\n")
	table := termTable.NewWriter(os.Stdout)
	var row ResultInfo
	table.SetHeader(row.Headers())

	mapStock.Each(func(key interface{}, value interface{}) {
		row := value.(ResultInfo)
		table.Append(row.Values())
	})
	table.Render() // Send output
}

func evaluate(bar *progressbar.ProgressBar, api FormulaApi, wg *sync.WaitGroup, code string, info *security.StaticBasic, result *treemap.Map) {
	defer wg.Done()
	defer bar.Add(1)

	wg.Add(1)
	api.Evaluate(code, info, result)
}
