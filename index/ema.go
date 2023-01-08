package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
)

const (
	EmaWeight float64 = 2.0000
	Ema12     string  = "Ema12"
	Ema50     string  = "Ema50"
)

type EmaLine struct {
	Ema12 float64 `json:"ema12"`
	Ema50 float64 `json:"ema50"`
}

type EMA struct {
	*cache.DataFrame
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
	self.DataFrame = cache.LoadDataFrame(code)
	if self.DataFrame == nil {
		return ErrCode
	} else if self.Length < 1 {
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
	count := self.Length
	ema12 := 0.000
	ema50 := 0.000
	for i, v := range self.Close {
		_close := v
		tmp := EmaLine{}
		//ema.Date = v.Date
		if i == 0 {
			// 第一天ema等于当天收盘价
			tmp.Ema12 = _close
			tmp.Ema50 = _close
		} else {
			// 第二天以后，当天收盘 收盘价乘以系数再加上昨天EMA乘以1-系数
			// 短期12天
			tmp.Ema12 = _close*factor12 + ema12*(1.0000-factor12)
			// 长期50天
			tmp.Ema50 = _close*factor50 + ema50*(1.0000-factor50)
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
