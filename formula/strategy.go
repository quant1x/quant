package main

import (
	"flag"
	"fmt"
	termTable "github.com/olekukonko/tablewriter"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data"
	"github.com/quant1x/quant/data/security"
	"github.com/schollz/progressbar/v3"
	"os"
	"sync"
)

// DataApi 数据接口
type FormulaApi interface {
	Name() string
	// 评估 日线数据
	Evaluate(fullCode string, info *security.StaticBasic, result *[]ResultInfo)
}

func main() {
	var (
		path     string
		strategy int
	)
	flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.IntVar(&strategy, "strategy", 1, "formula serial number")
	flag.Parse()
	cache.Init(path)
	var api FormulaApi
	switch strategy {
	case 89:
		api = new(FormulaNo89)
	default:
		api = new(FormulaNo1)
	}
	// 获取全部证券代码
	ss := data.GetCodeList()
	count := len(ss)
	var wg = sync.WaitGroup{}
	wg.Add(count)
	//bar := progressbar.DefaultSilent(int64(count))
	//doneCh := make(chan struct{})
	bar := progressbar.NewOptions(count,
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
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
		}),
	)
	//fmt.Printf("计划买入, 信号日期, 委托价格, 目标价位\n")
	result := make([]ResultInfo, 0)
	for _, v := range ss {
		bar.Add(1)
		fullCode := v
		basicInfo, err := security.GetBasicInfo(fullCode)
		if err == security.ErrCodeNotExist {
			// 证券代码不存在
			continue
		}
		if err != nil {
			// 其它错误 输出错误信息
			continue
		}

		go api.Evaluate(fullCode, basicInfo, &result)
	}
	//<-doneCh
	wg.Done()
	fmt.Println("\n ======= ["+api.Name()+"] progress bar completed ==========\n", "")
	table := termTable.NewWriter(os.Stdout)
	var t1 ResultInfo
	table.SetHeader(t1.Headers())

	for _, v := range result {
		table.Append(v.Values())
	}
	table.Render() // Send output
}
