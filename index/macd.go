package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/models/Cache"
)

type MacdLine struct {
	Cache.DayKLine
	Ema1 float64 `json:"ema1"` // 短周期, 12
	Ema2 float64 `json:"ema2"` // 长周期, 26
	Dif  float64 `json:"dif"`
	Dea  float64 `json:"dea"`
	Macd float64 `json:"macd"`
}

const (
	Ema1 = "Ema1"
	Ema2 = "Ema2"
	Dif  = "Dif"
	Dea  = "Dea"
	Macd = "Macd"
)

// MACD, 异同移动平均线
type MACD struct {
	data []MacdLine
	N1   int // 短周期
	N2   int // 长周期
	N    int // 差值周期
}

func LoadMacd(code string) *MACD {
	macd := MACD{}
	err := macd.Load(code)
	if err != nil {
		return nil
	}
	return &macd
}

func (self *MACD) Len() int {
	return len(self.data)
}

func (self *MACD) Data() interface{} {
	return self.data
}

func (self *MACD) Load(code string) error {
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
		self.N2 = 26
	}
	if self.N < 1 {
		self.N = 9
	}

	// Weighting factor, 计算加权因子
	factor1 := float64(EmaWeight / float64(self.N1+1))
	factor2 := float64(EmaWeight / float64(self.N2+1))
	//factor := float64(EmaWeight / float64(self.N+1))
	count := len(kls)
	ema1 := 0.000
	ema2 := 0.000
	dif := 0.000
	dea := 0.000
	//macd := 0.000
	for i, v := range kls {
		tmp := MacdLine{DayKLine: v}
		//ema.Date = v.Date
		if i == 0 {
			// 第一天ema等于当天收盘价
			tmp.Ema1 = v.Close
			tmp.Ema2 = v.Close
		} else {
			// 第二天以后，当天收盘 收盘价乘以系数再加上昨天EMA乘以1-系数
			// 短期12天
			tmp.Ema1 = v.Close*factor1 + ema1*(1.0000-factor1)
			// 长期50天
			tmp.Ema2 = v.Close*factor2 + ema2*(1.0000-factor2)
		}
		ema1 = tmp.Ema1
		ema2 = tmp.Ema2
		dif = ema1 - ema2
		// previous, Previous value
		//dea = (dea * 8 + dif * 2) / 10
		//dea = dif*factor + dea*(1.0000-factor)
		dea = ExpMA(dea, dif, 9)
		tmp.Dif = dif
		tmp.Dea = dea
		tmp.Macd = (tmp.Dif - tmp.Dea) * 2
		self.data = append(self.data, tmp)
		if category.DEBUG && count < i+3 {
			fmt.Printf("%+v\n", self.data[i])
		}
	}
	return nil
}
