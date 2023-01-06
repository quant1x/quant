package main

import (
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/index"
	"github.com/quant1x/quant/utils"
	"time"
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
	f = &index.K89{}
	f.Load(fullCode)

	//N := 3
	days := f.Len()
	if days < 100 {
		return
	}
	tmp := f.Data().(index.MaLine)
	if tmp.Close > 0 {
		//fmt.Printf("%s, %s, %.02f, %.02f\n", fullCode, hd.Date, hd.MA10, hd.MA10*1.05)
		//buy := fmt.Sprintf("%.3f", hd.MA10)
		//sell := fmt.Sprintf("%.3f", hd.MA10*1.05)
		now := time.Now()
		tt, _ := utils.ParseTime(tmp.Date)
		if utils.DifferDays(now, tt) < 5 {
			buy := tmp.Close
			sell := buy * 1.05
			*result = append(*result, ResultInfo{Code: fullCode,
				Name: info.Name,
				Date: tmp.Date,
				Buy:  buy,
				Sell: sell})
		}

	}
}
