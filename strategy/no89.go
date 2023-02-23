package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/indicator"
)

type FormulaNo89 struct{}

func (this FormulaNo89) Name() string {
	return "89K策略"
}

func (this FormulaNo89) Code() int {
	return 89
}

func (this FormulaNo89) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	//fmt.Println(fullCode)
	filename := cache.GetKLineFilename(fullCode)
	df := pandas.ReadCSV(filename)
	if df.Err != nil {
		return
	}
	_ = df.SetNames("date", "open", "high", "low", "close", "volume")
	N := 89
	CLOSE := df.Col("close")
	days := CLOSE.Len()
	date := df.Col("date").Values().([]string)[days-1]
	ret := indicator.F89K(df, N)

	rLen := ret.Nrow()
	B := ret.Col("B").Values().([]bool)
	buy := ret.Col("close").DTypes()
	if rLen > 1 && B[rLen-1] {
		buy := buy[rLen-1]
		sell := buy * 1.05
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			Buy:          float64(buy),
			Sell:         float64(sell),
			StrategyCode: this.Code(),
			StrategyName: this.Name()})
	}
}
