package main

import (
	"flag"
	"fmt"
	termTable "github.com/olekukonko/tablewriter"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/index"
	"github.com/schollz/progressbar/v3"
	"os"
)

//3天内5天线上穿10天线，10天线上穿20天线的个股
//count(cross(**(c,5),**(c,10)),3)>=1 and count(cross(**(c,10),**(c,20)),3)>=1
func main() {
	var (
		path string
	)
	flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.Parse()
	cache.Init(path, cache.CACHE_DATA_CSV)
	// 获取全部证券代码
	ss := data.GetCodeList()
	count := len(ss)
	//bar := progressbar.DefaultSilent(int64(count))
	//doneCh := make(chan struct{})
	bar := progressbar.NewOptions(count,
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(80),
		progressbar.OptionSetDescription("[cyan][1/3][reset]执行1号策略..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[red]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	//fmt.Printf("计划买入, 信号日期, 委托价格, 目标价位\n")
	result := make([][]string, 0)
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

		//fmt.Printf("%s\n", fullCode)
		var f index.Formula
		f = &index.MA{}
		f.Load(fullCode)

		N := 3
		days := f.Len()
		if days < 100 {
			continue
		}
		hds := f.Data().([]index.MaLine)
		for i := 0; i < N; i++ {
			hd := hds[days-i-1]
			a := index.CompVal{
				Data:  hds,
				Flag:  index.MA5,
				Cycle: i,
			}
			b := index.CompVal{
				Data:  hds,
				Flag:  index.MA10,
				Cycle: i,
			}
			c := index.CompVal{
				Data:  hds,
				Flag:  index.MA20,
				Cycle: i,
			}
			// 过滤 超过10.00的股票
			//if hd.Close > 9.00 || hd.Close < 4.00 {
			//	continue
			//}
			//sh := stock.StockHistory{Data:hds[i:]}
			//ma5 := float64(hd.MA5 * float64(hd.MA5Volume))
			//ma10 := float64(hd.MA10 * float64(hd.MA10Volume))
			//ma20 := float64(hd.MA20 * float64(hd.MA20Volume))
			//b1 := sh.Cross(stock.MA5, stock.MA10)
			b1 := index.Cross(a, b)
			//b2 := sh.Cross(stock.MA10, stock.MA20)
			b2 := index.Cross(b, c)
			if b1 && b2 {
				//fmt.Printf("%s, %s, %.02f, %.02f\n", fullCode, hd.Date, hd.MA10, hd.MA10*1.05)
				buy := fmt.Sprintf("%.3f", hd.MA10)
				sell := fmt.Sprintf("%.3f", hd.MA10*1.05)
				result = append(result, []string{fullCode, basicInfo.Name, hd.Date, buy, sell})
				break
			}
		}
	}
	//<-doneCh
	fmt.Println("\n ======= progress bar completed ==========\n")
	table := termTable.NewWriter(os.Stdout)
	table.SetHeader([]string{"证券代码", "证券名称", "信号日期", "委托价格", "目标价位"})

	for _, v := range result {
		table.Append(v)
	}
	table.Render() // Send output
}
