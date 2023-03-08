package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/labs/linear"
)

type FormulaNo3 struct{}

func (this FormulaNo3) Name() string {
	return "W底"
}

func (this FormulaNo3) Code() int {
	return 3
}

func (this FormulaNo3) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	N := 89
	filename := cache.KLineFilename(fullCode)
	df := pandas.ReadCSV(filename)
	if df.Err != nil {
		return
	}
	if df.Nrow() < N {
		return
	}
	if p, ok := linear.W(df); ok {
		rLen := df.Nrow()
		date := df.Col("date").Values().([]string)[rLen-1]
		closes := df.Col("close").DTypes()
		buy := closes[rLen-1]
		sell := p
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			Buy:          buy,
			Sell:         sell,
			StrategyCode: this.Code(),
			StrategyName: this.Name()})

	}
}