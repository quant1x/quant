package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	. "gitee.com/quant1x/pandas/formula"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/treemap"
)

// FormulaNo1 3天内5天线上穿10天线，10天线上穿20天线的个股
//
//	count(cross(MA(c,5),MA(c,10)),3)>=1 and count(cross(MA(c,10),MA(c,20)),3)>=1
type FormulaNo1 struct {
}

func (this *FormulaNo1) Code() int {
	return 1
}

func (this *FormulaNo1) Name() string {
	return "1号策略"
}

// Evaluate 评估K线数据
func (this *FormulaNo1) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	N := MaximumResultDays

	//fmt.Printf("%s\n", fullCode)
	filename := cache.GetKLineFilename(fullCode)
	df := pandas.ReadCSV(filename)
	if df.Err != nil {
		return
	}
	_ = df.SetNames("date", "open", "high", "low", "close", "volume")
	// 收盘价序列
	CLOSE := df.Col("close")
	days := CLOSE.Len()
	if days < 1 {
		return
	}

	// 取5、10、20日均线
	ma5 := MA(CLOSE, 5)
	ma10 := MA(CLOSE, 10)
	ma20 := MA(CLOSE, 20)
	if ma5.Len() != days || ma10.Len() != days || ma20.Len() != days {
		logger.Errorf("均线, 数据没对齐")
	}
	// 两个金叉
	c1 := CROSS(ma5, ma10)
	c2 := CROSS(ma10, ma20)
	if c1.Len() != days || c2.Len() != days {
		logger.Errorf("金叉, 数据没对齐")
	}
	// 两个统计
	r1 := COUNT(c1.Bools(), N)
	r2 := COUNT(c2.Bools(), N)
	if r1.Len() != days || r2.Len() != days {
		logger.Errorf("统计, 数据没对齐")
	}
	// 横向对比
	d := r1.And(r2)
	if d.Len() != days {
		logger.Errorf("横向对比, 数据没对齐")
	}

	//cc1 := CompareGte(r1, 1)
	rLen := d.Len()
	//fmt.Printf("%s, len = %d\n", fullCode, rLen)
	b := d.IndexOf(rLen - 1).(bool)
	if rLen > 1 && b {
		buy := stat.Any2DType(ma10.IndexOf(days - 1))
		sell := buy * 1.05
		date := df.Col("date").Values().([]string)[days-1]
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			Buy:          buy,
			Sell:         sell,
			StrategyCode: this.Code(),
			StrategyName: this.Name()})
	}
}
