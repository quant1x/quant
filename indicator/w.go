package indicator

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"gitee.com/quant1x/t89k/labs/linear"
)

func MAX_GO[T stat.Ordered](a, b T) (x, y T) {
	if a >= b {
		return a, b
	}
	return b, a
}

// W W底突破颈线
func W(raw pandas.DataFrame, argv ...bool) (stat.DType, bool) {
	const (
		//__delta = float64(0.0191)
		__delta = linear.TrendDelta
	)
	var (
		__debug        bool = false
		__ignoreSignal      = false
	)
	if raw.Nrow() < linear.MaximumTrendPeriod {
		return 0.00, false
	}
	if len(argv) > 0 {
		__debug = argv[0]
	}
	if len(argv) > 1 {
		__ignoreSignal = argv[1]
	}
	df := raw.Subset(raw.Nrow()-linear.MaximumTrendPeriod, raw.Nrow())
	//fmt.Println(df)
	// 获取全部的波峰的数据
	vh := df.Col("high")
	v := vh.DTypes()
	_, _, maxi, maxv := linear.PeakDetect(v[:], __delta)
	if len(maxi) < 1 {
		return 0.00, false
	}
	// 获取全部的波谷的数据
	vl := df.Col("low")
	v = vl.DTypes()
	mini, minv, _, _ := linear.PeakDetect(v[:], __delta)
	var zcX1, zcX2 int
	var zcY1, zcY2 float64
	//var zcSlope float64
	if len(mini) >= 2 {
		w := len(mini)
		zcX1 = mini[w-2]
		zcY1 = minv[w-2]
		zcX2 = mini[w-1]
		zcY2 = minv[w-1]
		//zcSlope = num.Slope(zcX1, zcY1, zcX2, zcY2)
	} else {
		return 0.00, false
	}
	// 找到最近的一个波峰
	maxX := -1
	maxY := float64(-1) // 颈线
	for i := 0; i < len(maxi); i++ {
		if maxi[i] > zcX1 && maxi[i] < zcX2 {
			maxX = maxi[i]
			maxY = maxv[i]
			break
		}
	}

	if maxX > 0 {
		dbH, dbL := MAX_GO(zcY1, zcY2)
		pLow := maxY*2 - dbH
		pHigh := maxY*2 - dbL
		dates := df.Col("date").Strings()
		closes := df.Col("close").DTypes()
		cl := len(closes)
		//if closes[cl-2] < maxY && closes[cl-1] > maxY && pHigh/closes[cl-1] >= 1.10 {
		//if closes[cl-2] < maxY && closes[cl-1] > maxY && pHigh/closes[cl-1] >= 1.10 {
		if __ignoreSignal || (closes[cl-2] < closes[cl-1] && closes[cl-2] < zcY1 && closes[cl-1] > zcY1) {
			if __debug {
				fmt.Printf("W底突破颈线: 颈线(%s)=%f, %s\n", dates[maxX], maxY, dates[cl-1])
				fmt.Printf("       W底: 左低(%s)=%f, 右低(%s)=%f, 颈线(%s)=%f\n", dates[zcX1], zcY1, dates[zcX2], zcY2, dates[maxX], maxY)
				fmt.Printf("预计目标位置: 最低=%f, 最高=%f\n", pLow, pHigh)
			}
			return pHigh, true
		}
	}

	return 0.00, false
}
