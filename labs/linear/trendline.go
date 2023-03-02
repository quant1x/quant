package linear

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/quant1x/quant/num"
)

const (
	//MaximumTrendPeriod = 144
	MaximumTrendPeriod = 89
	TrendDelta         = float64(0.0618)
	//TrendDelta = float64(0.00191)
)

// TrendLine 原则上是不抄底, 当底部上移时, 观察压力线和支撑线的相对关系
func TrendLine(raw pandas.DataFrame) pandas.DataFrame {
	df := raw.Subset(raw.Nrow()-MaximumTrendPeriod, raw.Nrow())
	fmt.Println(df)
	// 第一步: 最低价的波谷
	vl := df.Col("close")
	v := vl.DTypes()
	mini, minv, maxi, maxv := PeakDetect(v[:], TrendDelta)
	var zcX1, zcX2 int
	var zcY1, zcY2 float64
	var zcSlope, ylSlope float64
	if len(mini) >= 2 {
		w := len(minv)
		zcX1 = mini[w-2]
		zcY1 = minv[w-2]
		zcX2 = mini[w-1]
		zcY2 = minv[w-1]
		zcSlope = num.Slope(zcX1, zcY1, zcX2, zcY2)
		fmt.Println("波谷x =", mini)
		fmt.Println("波谷y =", minv)
		//fmt.Println("波峰x =", maxi)
		//fmt.Println("波峰y =", maxv)
	}
	// 第二步: 最高价的波峰
	vh := df.Col("high")
	v = vh.DTypes()
	mini, minv, maxi, maxv = PeakDetect(v[:], TrendDelta)
	var ylX1, ylX2 int
	var ylY1, ylY2 float64
	if len(maxi) >= 2 {
		w := len(maxi)
		ylX1 = maxi[w-2]
		ylY1 = maxv[w-2]
		ylX2 = maxi[w-1]
		ylY2 = maxv[w-1]
		ylSlope = num.Slope(ylX1, ylY1, ylX2, ylY2)
		//fmt.Println("波谷x =", mini)
		//fmt.Println("波谷y =", minv)
		fmt.Println("波峰x =", maxi)
		fmt.Println("波峰y =", maxv)
	}

	fmt.Println("支撑趋势线斜率 =", zcSlope, "压力趋势线斜率 =", ylSlope)
	CLOSE := df.Col("close")
	length := raw.Nrow()
	tZc := make([]stat.DType, length)
	tYl := make([]stat.DType, length)
	cross := make([]bool, length)
	pos := length - MaximumTrendPeriod
	CLOSE.Apply(func(idx int, v any) {
		//vf := stat.AnyToFloat64(v)
		if idx >= zcX1 {
			tZc[pos+idx] = num.TriangleBevel(zcSlope, zcX1, zcY1, idx)
		}
		if idx >= ylX1 {
			tYl[pos+idx] = num.TriangleBevel(ylSlope, ylX1, ylY1, idx)
		}
	})
	zc := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "zc", tZc)
	yl := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "yl", tYl)
	sc := pandas.NewSeries(stat.SERIES_TYPE_BOOL, "cross", cross)
	df = raw.Join(zc).Join(yl).Join(sc)
	return df
}

// CrossTrend 原则上是不抄底, 当底部上移时, 观察压力线和支撑线的相对关系
func CrossTrend(raw pandas.DataFrame) pandas.DataFrame {
	if raw.Nrow() < MaximumTrendPeriod {
		return pandas.DataFrame{}
	}
	df := raw.Subset(raw.Nrow()-MaximumTrendPeriod, raw.Nrow())
	//fmt.Println(df)
	vh := df.Col("high")
	v := vh.DTypes()
	mini, minv, maxi, maxv := PeakDetect(v[:], TrendDelta)
	var ylX1, ylX2 int
	var ylY1, ylY2 float64
	var ylSlope float64
	if len(maxi) >= 2 {
		w := len(maxi)
		ylX1 = maxi[w-2]
		ylY1 = maxv[w-2]
		ylX2 = maxi[w-1]
		ylY2 = maxv[w-1]
		ylSlope = num.Slope(ylX1, ylY1, ylX2, ylY2)
		//fmt.Println("波谷x =", mini)
		//fmt.Println("波谷y =", minv)
		//fmt.Println("波峰x =", maxi)
		//fmt.Println("波峰y =", maxv)
	}

	//fmt.Println("压力趋势线斜率 =", ylSlope)
	CLOSE := df.Col("close")
	length := raw.Nrow()
	tZc := make([]stat.DType, length)
	tYl := make([]stat.DType, length)
	cross := make([]bool, length)
	pos := length - MaximumTrendPeriod
	CLOSE.Apply(func(idx int, v any) {
		vf := stat.AnyToFloat64(v)
		if idx >= ylX1 {
			tYl[pos+idx] = num.TriangleBevel(ylSlope, ylX1, ylY1, idx)
		}
		if vf > tYl[pos+idx] {
			cross[pos+idx] = true
		}
	})
	zc := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "zc", tZc)
	yl := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "yl", tYl)
	sc := pandas.NewSeries(stat.SERIES_TYPE_BOOL, "cross", cross)
	df = raw.Join(zc).Join(yl).Join(sc)
	_ = mini
	_ = minv
	return df
}
