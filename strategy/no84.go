package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/indicator"
)

type FormulaNo84 struct{}

func (this FormulaNo84) Name() string {
	return "缠论时间窗口"
}

func (this FormulaNo84) Code() int {
	return 84
}

func (this FormulaNo84) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	defer func() {
		// 解析失败以后输出日志, 以备检查
		if err := recover(); err != nil {
			logger.Errorf("FormulaNo89.Evaluate code=%s, error=%+v\n", fullCode, err)
		}
	}()
	//fmt.Println(fullCode)
	N := 6
	filename := cache.KLineFilename(fullCode)
	df := pandas.ReadCSV(filename)
	if df.Err != nil {
		return
	}
	if df.Nrow() < N {
		return
	}
	_ = df.SetNames("date", "open", "high", "low", "close", "volume")
	CLOSE := df.Col("close")
	days := CLOSE.Len()
	date := df.Col("date").Values().([]string)[days-1]
	ret := indicator.MA4X(df, N)
	if ret.Nrow() < 1 {
		return
	}
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
