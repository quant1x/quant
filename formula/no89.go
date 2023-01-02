package main

import (
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/index"
)

// 89K策略
type FormulaNo89 struct {
}

func (this *FormulaNo89) Name() string {
	return "89K策略"
}

// Evaluate 评估K线数据
func (this *FormulaNo89) Evaluate(fullCode string, info *security.StaticBasic, result *[]ResultInfo) {
	//fmt.Printf("%s\n", fullCode)
	var f index.Formula
	f = &index.MA{}
	f.Load(fullCode)

	N := 3
	days := f.Len()
	if days < 100 {
		return
	}
	var (
		k5  float64 = 0.00
		n5  int     = -1
		k7  float64 = 0.00
		n7  int     = -1
		k8  float64 = 0.00
		n8  int     = -1
		k9  float64 = 0.00
		n9  int     = -1
		k10 float64 = 0.00
		n10 int     = -1
	)
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
		// 第一步, 找最低价
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
			*result = append(*result, ResultInfo{Code: fullCode,
				Name: info.Name,
				Date: hd.Date,
				Buy:  buy,
				Sell: sell})
			break
		}
	}
}
