package main

import (
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/indicator"
)

type FormulaNo85 struct{}

func (this FormulaNo85) Name() string {
	return "抄底逃顶"
}

func (this FormulaNo85) Code() int {
	return 85
}

func (this FormulaNo85) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	N := 89
	df := stock.KLine(fullCode)
	if df.Nrow() < N {
		return
	}
	df1 := indicator.CDTD(df)

	rLen := df.Nrow()
	b := df1.Col("B").IndexOf(-1).(bool)
	if b {
		date := df.Col("date").Values().([]string)[rLen-1]
		closes := df.Col("close").DTypes()
		buy := closes[rLen-1]
		sell := buy * 1.05
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			Buy:          buy,
			Sell:         sell,
			StrategyCode: this.Code(),
			StrategyName: this.Name()})
	}
}
