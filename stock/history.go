package stock

import (
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/models/Cache"
	"reflect"
	"strings"
)

const (
	Open       = "Open"
	Close      = "Close"
	High       = "High"
	Low        = "Low"
	Volume     = "Volume"
	MA5        = "MA5"
	MA5Volume  = "MA5Volume"
	MA10       = "MA10"
	MA10Volume = "MA10Volume"
	MA20       = "MA20"
	MA20Volume = "MA20Volume"
	MA30       = "MA30"
	MA30Volume = "MA30Volume"
)

type History struct {
	// day        时间
	Day string `json:"day"`
	// open       开盘价
	Open float64 `json:"open"`
	// close      收盘价
	Close float64 `json:"close"`
	// high       最高价
	High float64 `json:"high"`
	// low        最低价
	Low float64 `json:"low"`
	// volume     成交量, 单位股, 除以100为手
	Volume int64 `json:"volume"`
	// MA5        五日平均价
	MA5 float64 `json:"ma_price5"`
	// MA5Volume  五日平均交易量
	MA5Volume int64 `json:"ma_volume5"`
	// MA10       十日平均价
	MA10 float64 `json:"ma_price10"`
	// MA10Volume 十日平均交易量
	MA10Volume int64 `json:"ma_volume10"`
	// MA20       二十日平均价
	MA20 float64 `json:"ma_price20"`
	// MA20Volume 二十日平均交易量
	MA20Volume int64 `json:"ma_volume20"`
	// MA30       三十日平均价
	MA30 float64 `json:"ma_price30"`
	// MA30Volume 三十日平均交易量
	MA30Volume int64 `json:"ma_volume30"`
}

func KLinePath(fc string) (string, string, int) {
	fc = strings.TrimSpace(fc)
	fcLen := len(fc)
	if fcLen != 7 && fcLen != 8 {
		return fc, "", D_ERROR | D_ECODE
	}
	pos := len(fc) - 3
	fc = strings.ToLower(fc)
	// 组织存储路径
	filename := cache.GetDayPath() + "/" + fc[0:pos] + "/" + fc
	if cache.CACHE_DATA_CSV {
		filename += ".csv"
	}
	return fc, filename, D_OK
}

//
func LoadKLine(fullCode string) []Cache.DayKLine {
	return cache.LoadKLine(fullCode)
}

// 加载历史数据
func LoadHistory(kls []Cache.DayKLine) []History {
	klines := kls

	var hds []History
	for i, v := range klines {
		var hd History
		hd.Day = v.Date
		hd.Open = v.Open
		hd.Close = v.Close
		hd.High = v.High
		hd.Low = v.Low
		hd.Volume = v.Volume
		// 5日均线
		//hd.MA5, hd.MA5Volume = index.MAX(klines, 5, i)
		// 10日均线
		//hd.MA10, hd.MA10Volume = index.MAX(klines, 10, i)
		// 20日均线
		//hd.MA20, hd.MA20Volume = index.MAX(klines, 20, i)
		// 30日均线
		//hd.MA30, hd.MA30Volume = index.MAX(klines, 30, i)
		hds = append(hds, hd)
		_ = i
	}
	return hds
}

type StockHistory struct {
	Data []History
}

func (self StockHistory) ref(flag string, n int) float64 {
	hds := self.Data
	if len(hds) < n {
		return 0
	}
	hd := hds[n]

	t := reflect.ValueOf(hd)
	v := t.FieldByName(flag)
	return v.Float()
}

// Cross 金叉, a 上穿 b, a和b对调就是死叉
func (self StockHistory) Cross(a, b string) bool {
	f1 := self.ref(a, 1)
	f2 := self.ref(b, 1)

	if f1 >= f2 {
		return false
	} else {
		f1 = self.ref(a, 0)
		f2 = self.ref(b, 0)

		if f1 <= f2 {
			return false
		} else {
			return true
		}
	}
}
