package index

import (
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/models/Cache"
)

type MaLine struct {
	Cache.DayKLine
	MA5        float64 `json:"ma5_price"`
	MA5Volume  int64   `json:"MA5_volume"`
	MA10       float64 `json:"ma10_price"`
	MA10Volume int64   `json:"ma10_volume"`
	MA20       float64 `json:"ma20_price"`
	MA20Volume int64   `json:"ma20_volume"`
	MA30       float64 `json:"ma30_price"`
	MA30Volume int64   `json:"ma30_volume"`
}

const (
	MA5        = "MA5"
	MA5Volume  = "MA5Volume"
	MA10       = "MA10"
	MA10Volume = "MA10Volume"
	MA20       = "MA20"
	MA20Volume = "MA20Volume"
	MA30       = "MA30"
	MA30Volume = "MA30Volume"
)

type MA1X struct {
	data []MaLine
}

func (self *MA1X) Len() int {
	return len(self.data)
}

func (self *MA1X) Data() interface{} {
	return self.data
}

// 5, 10, 20, 30
func (self *MA1X) Load(code string) error {
	kls := cache.LoadKLine(code)
	if kls == nil {
		return ErrCode
	} else if len(kls) < 1 {
		return ErrData
	}

	count := len(kls)
	for i, v := range kls {
		var (
			price  float64
			volume int64
		)
		tmp := MaLine{DayKLine: v}

		if i+1 >= 5 {
			x := 5
			hds := kls[:i+1]
			//tmpArrs := self.data[:i]
			//price = Ref(tmpArrs, MA5, 1)
			//volume = RefInt(tmpArrs, MA5Volume, 1)
			if price == 0 {
				//price, volume = MAX(kls[i - 5  + 1:i + 1], 5, 1)
				price = SUM(hds, Close, x) / float64(x)
				volume = int64(SUM(hds, Volume, x)) / int64(x)
			} else {
				//price = (price * float64(x) + v.Close - Ref(tmpArrs, Close, 1)) / float64(x)
				//volume = (volume * int64(x) + v.Volume - int64(Ref(tmpArrs, Volume, 1))) / int64(x)
				price = SUM(hds, Close, x) / float64(x)
				volume = int64(SUM(hds, Volume, x)) / int64(x)
			}
			tmp.MA5 = price
			tmp.MA5Volume = volume
		}
		if i+1 >= 10 {
			x := 10
			hds := kls[:i+1]
			price = SUM(hds, Close, x) / float64(x)
			volume = int64(SUM(hds, Volume, x)) / int64(x)
			tmp.MA10 = price
			tmp.MA10Volume = volume
		}
		if i+1 >= 20 {
			x := 20
			hds := kls[:i+1]
			price = SUM(hds, Close, x) / float64(x)
			volume = int64(SUM(hds, Volume, x)) / int64(x)
			tmp.MA20 = price
			tmp.MA20Volume = volume
		}
		if i+1 >= 30 {
			x := 30
			hds := kls[:i+1]
			price = SUM(hds, Close, x) / float64(x)
			volume = int64(SUM(hds, Volume, x)) / int64(x)
			tmp.MA30 = price
			tmp.MA30Volume = volume
		}
		self.data = append(self.data, tmp)
		// 输出最后2组数据
		if category.DEBUG && count < i+3 {
			//fmt.Printf("day: %s, MA5: %.2f, MA10: %.2f, MA20: %.2f, MA30: %.2f\n", self.data[i].Date, self.data[i].MA5, self.data[i].MA10, self.data[i].MA20, self.data[i].MA30)
		}
	}
	return nil
}
