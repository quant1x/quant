package main

import (
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/data/stock"
	"github.com/mymmsc/gox/util/treemap"
)

// FormulaNo0 0号策略
type FormulaNo0 struct{}

func (this *FormulaNo0) Name() string {
	return "0号策略"
}

func (this *FormulaNo0) Code() int {
	return 0
}

func (this *FormulaNo0) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	df := stock.KLine(fullCode)
	if df.Err != nil {
		//fmt.Println(fullCode)
		return
	}
	days := df.Nrow()
	if days < 1 {
		//fmt.Println(fullCode)
		return
	}
	turnZ := stock.GetTurnZ(fullCode)
	//fmt.Println(fullCode, turnZ, kpVol, freeGuBen)

	CLOSE := df.ColAsNDArray("close")
	ZF := df.ColAsNDArray("zf")
	if days > 1 {
		buy := CLOSE.IndexOf(-1).(float64)
		zf := ZF.IndexOf(-1).(float64)
		sell := buy * 1.05
		date := df.Col("date").Values().([]string)[days-1]
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			TurnZ:        turnZ,
			Rate:         zf,
			Buy:          buy,
			Sell:         sell,
			StrategyCode: this.Code(),
			StrategyName: this.Name()})
	} else {
		//fmt.Println(fullCode)
	}
}
