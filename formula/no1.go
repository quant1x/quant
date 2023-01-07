package main

import (
	"github.com/mymmsc/gox/util/arraylist"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/index"
)

// 3天内5天线上穿10天线，10天线上穿20天线的个股
// count(cross(**(c,5),**(c,10)),3)>=1 and count(cross(**(c,10),**(c,20)),3)>=1
type FormulaNo1 struct {
}

func (this *FormulaNo1) Name() string {
	return "1号策略"
}

// Evaluate 评估K线数据
func (this *FormulaNo1) Evaluate(fullCode string, info *security.StaticBasic, result *arraylist.List) {
	//fmt.Printf("%s\n", fullCode)
	var f index.Formula
	f = &index.MA{}
	f.Load(fullCode)

	N := 3
	days := f.Len()
	if days < 100 {
		return
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
			//buy := fmt.Sprintf("%.3f", hd.MA10)
			//sell := fmt.Sprintf("%.3f", hd.MA10*1.05)

			buy := hd.MA10
			sell := hd.MA10 * 1.05

			result.Add(ResultInfo{Code: fullCode,
				Name: info.Name,
				Date: hd.Date,
				Buy:  buy,
				Sell: sell})
			break
		}
	}
}
