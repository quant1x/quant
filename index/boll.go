package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/models/Cache"
	"math"
)

type BollLine struct {
	Cache.DayKLine
	up  float64 // 上轨
	mid float64 // 中轨
	low float64 // 下轨
}

type Boll struct {
	n    int
	k    float64
	data []BollLine
}

// NewBoll BOLL(20, 2)
func NewBoll(n int, k int32) *Boll {
	return &Boll{n: n, k: float64(k)}
}

func (self *Boll) Len() int {
	return len(self.data)
}

func (self *Boll) Data() interface{} {
	return self.data
}

func (self *Boll) Load(code string) error {
	kls := cache.LoadKLine(code)
	if kls == nil {
		return ErrCode
	} else if len(kls) < 1 {
		return ErrData
	}
	count := len(kls)

	//mid = make([]float64, l)
	//up = make([]float64, l)
	//low = make([]float64, l)
	if count < self.n {
		return ErrData
	}
	for i := 0; i < count; i++ {
		v := kls[i]
		tmp := BollLine{DayKLine: v}
		if i >= self.n-1 {
			ps := kls[(i - self.n + 1) : i+1]
			tmp.mid = self.sma(ps)
			dm := self.k * self.dma(ps, tmp.mid)
			tmp.up = tmp.mid + dm
			tmp.low = tmp.mid - dm
		}

		self.data = append(self.data, tmp)
		if category.DEBUG && count < i+3 {
			fmt.Printf("%+v\n", self.data[i])
		}
	}

	return nil
}

func (self *Boll) sma(lines []Cache.DayKLine) float64 {
	s := len(lines)
	var sum float64 = 0.00
	for i := 0; i < s; i++ {
		sum += lines[i].Close
	}
	return sum / float64(s)
}

func (self *Boll) dma(lines []Cache.DayKLine, ma float64) float64 {
	s := len(lines)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		//sum += (lines[i].Close - ma) * (lines[i].Close - ma)
		sum += math.Pow(lines[i].Close-ma, 2)
	}
	return math.Sqrt(sum / float64(self.n-1))
}

func (self *Boll) Boll(lines []Cache.DayKLine) (mid []float64, up []float64, low []float64) {
	l := len(lines)

	mid = make([]float64, l)
	up = make([]float64, l)
	low = make([]float64, l)
	if l < self.n {
		return
	}
	for i := l - 1; i > self.n-1; i-- {
		ps := lines[(i - self.n + 1) : i+1]
		mid[i] = self.sma(ps)
		dm := self.k * self.dma(ps, mid[i])
		up[i] = mid[i] + dm
		low[i] = mid[i] - dm
	}

	return
}

//在所有的指标计算中，BOLL指标的计算方法是最复杂的之一，其中引进了统计学中的标准差概念，
//涉及到中轨线（MB）、上轨线（UP）和下轨线（DN）的计算。另外，和其他指标的计算一样，由于选用的计算周期的不同，
//BOLL指标也包括日BOLL指标、周BOLL指标、月BOLL指标年BOLL指标以及分钟BOLL指标等各种类型。
//经常被用于股市研判的是日BOLL指标和周BOLL指标。虽然它们的计算时的取值有所不同，但基本的计算方法一样。
//以日BOLL指标计算为例，其计算方法如下：
//日BOLL指标的计算公式
//中轨线=N日的移动平均线
//上轨线=中轨线+两倍的标准差
//下轨线=中轨线－两倍的标准差
//日BOLL指标的计算过程
//1）计算MA
//MA=N日内的收盘价之和÷N
//2）计算标准差MD
//MD=平方根N日的（C－MA）的两次方之和除以N
//3）计算MB、UP、DN线
//MB=N日的MA
//UP=MB+2×MD
//DN=MB－2×MD
//各大股票交易软件默认N是20，所以MB等于当日20日均线值
