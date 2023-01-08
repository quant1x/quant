package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/formula"
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
	*cache.DataFrame
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
	self.DataFrame = cache.LoadDataFrame(code)
	if self.DataFrame == nil {
		return ErrCode
	} else if self.Length < 1 {
		return ErrData
	}

	count := self.Length
	for i := 0; i < count; i++ {
		end := i + 1
		tmp := MaLine{
			MA5:       formula.MA(self.Close[:end], 5),
			MA5Volume: formula.MA(self.Volume[:end], 5),

			MA10:       formula.MA(self.Close[:end], 10),
			MA10Volume: formula.MA(self.Volume[:end], 10),

			MA20:       formula.MA(self.Close[:end], 20),
			MA20Volume: formula.MA(self.Volume[:end], 20),

			MA30:       formula.MA(self.Close[:end], 30),
			MA30Volume: formula.MA(self.Volume[:end], 30),
		}
		tmp.Date = self.Date[i]
		self.data = append(self.data, tmp)
		// 输出最后2组数据
		if category.DEBUG && count < i+3 {
			fmt.Printf("day: %s, MA5: %.2f, MA10: %.2f, MA20: %.2f, MA30: %.2f\n", self.Date[i], self.data[i].MA5, self.data[i].MA10, self.data[i].MA20, self.data[i].MA30)
		}
	}
	return nil
}
