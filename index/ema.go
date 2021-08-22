package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/models/Cache"
)

const (
	EmaWeight float64 = 2.0000
	Ema12     string  = "Ema12"
	Ema50     string  = "Ema50"
)

type EmaLine struct {
	Cache.DayKLine
	Ema12 float64 `json:"ema12"`
	Ema50 float64 `json:"ema50"`
}

type EMA struct {
	data []EmaLine
	N1   int // 短周期
	N2   int // 长周期
}

func LoadEma(code string) *EMA {
	ema := EMA{}
	err := ema.Load(code)
	if err != nil {
		return nil
	}
	return &ema
}

func (self *EMA) Len() int {
	return len(self.data)
}

func (self *EMA) Data() interface{} {
	return self.data
}

// 短期12天 长期50天
func (self *EMA) Load(code string) error {
	kls := cache.LoadKLine(code)
	if kls == nil {
		return ErrCode
	} else if len(kls) < 1 {
		return ErrData
	}

	if self.N1 < 1 {
		self.N1 = 12
	}
	if self.N2 < 1 {
		self.N2 = 50
	}

	// Weighting factor, 计算加权因子
	factor12 := float64(EmaWeight / float64(self.N1+1))
	factor50 := float64(EmaWeight / float64(self.N2+1))
	count := len(kls)
	ema12 := 0.000
	ema50 := 0.000
	for i, v := range kls {
		tmp := EmaLine{DayKLine: v}
		//ema.Date = v.Date
		if i == 0 {
			// 第一天ema等于当天收盘价
			tmp.Ema12 = v.Close
			tmp.Ema50 = v.Close
		} else {
			// 第二天以后，当天收盘 收盘价乘以系数再加上昨天EMA乘以1-系数
			// 短期12天
			tmp.Ema12 = v.Close*factor12 + ema12*(1.0000-factor12)
			// 长期50天
			tmp.Ema50 = v.Close*factor50 + ema50*(1.0000-factor50)
		}
		ema12 = tmp.Ema12
		ema50 = tmp.Ema50
		self.data = append(self.data, tmp)
		if category.DEBUG && count < i+3 {
			fmt.Printf("12: %.2f, 50: %.2f\n", self.data[i].Ema12, self.data[i].Ema50)
		}
	}
	return nil
}

// EMA（Exponential Moving Average）是指数移动平均值。
// 也叫EXPMA指标，它也是一种趋向类指标，指数移动平均值是以指数式递减加权的移动平均。
// EXPMA=（当日或当期收盘价－上一日或上期EXPMA）/N+上一日或上期EXPMA，其中，首次上期EXPMA值为上一期收盘价，N为天数。

// 当天EMA=昨天的EMA+加权因子*（当天的收盘价-昨天的EMA）
// = 加权因子*当天的收盘价+（1-加权因子）*昨天的EMA
// 加权因子=2/(N+1);
// N就是上面所说的周期 ，比如周期12 则加权的因子就是 2/13；
// 当天EMA=2/13*当天的收盘价+11/13*昨天的EMA
// 计算过程：（每日你看到的EMA计算结果是从上市第一天就开始累积了）
// 股票上市第一天：当天EMA1 = 当天收盘价
// 第二天：EMA2 = 2/13 * 当天收盘价 + 11/13 * EMA1
// 第三天：EMA3 = 2/13 * 当天收盘价 + 11/13 * EMA2

// 若求X的N日指数平滑移动平均, 则表达式为: EMA(X, N)
// 算法是: 若Y=EMA(X, N), 则Y=[2*X+(N-1)*Y’]/(N+1), 其中Y’表示上一周期的Y值。
func EXMPA0(hds []Cache.DayKLine, n int, x int) float64 {
	count := len(hds)
	if n > x+1 || count < n {
		return 0.000
	}

	// Weighting factor, 计算加权因子
	factor := float64(2.0000 / (n + 1.0000))
	// 第一天ema等于当天收盘价
	ema := hds[0].Close
	for i := 1; i < n; i++ {
		hd := hds[i]
		// 第二天以后，当天收盘 收盘价乘以系数再加上昨天EMA乘以1-系数
		ema = hd.Close*factor + ema*(1.0000-factor)
	}
	return ema
}
