package linear

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"github.com/quant1x/quant/num"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"testing"
)

func TestFindPeaks(t *testing.T) {
	code := "sh000905"
	code = "sz002528"
	code = "sh600602"
	length := 89
	df := stock.KLine(code)
	df = df.Subset(df.Nrow()-length, df.Nrow())
	fmt.Println(df)
	vh := df.Col("close")
	//vh.Mean()
	v := vh.DTypes()
	mini, minv, maxi, maxv := PeakDetect(v[:], 0.0618)

	fmt.Println("波谷x =", mini)
	fmt.Println("波谷y =", minv)
	fmt.Println("波峰x =", maxi)
	fmt.Println("波峰y =", maxv)

	var x1, x2 int
	var y1, y2 float64
	if len(maxv) < 2 {
		return
	}
	w := len(maxv)
	x1 = maxi[w-2]
	y1 = maxv[w-2]
	x2 = maxi[w-1]
	y2 = maxv[w-1]

	w = len(minv)
	x1 = mini[w-2]
	y1 = minv[w-2]
	x2 = mini[w-1]
	y2 = minv[w-1]

	slope := num.Slope(x1, y1, x2, y2)
	fmt.Println("斜率 =", slope)
	CLOSE := df.Col("close")
	// slope*float64(xn-x1) + y1
	p1 := make([]stat.DType, CLOSE.Len())
	cross := make([]bool, CLOSE.Len())
	CLOSE.Apply(func(idx int, v any) {
		vf := stat.AnyToFloat64(v)
		if idx > x2 {
			p1[idx] = num.TriangleBevel(slope, x1, y1, idx)
			cross[idx] = vf > p1[idx]
		}
	})
	sp := pandas.NewSeries(stat.SERIES_TYPE_DTYPE, "p1", p1)
	sc := pandas.NewSeries(stat.SERIES_TYPE_BOOL, "cross", cross)
	df = df.Join(sp).Join(sc)
	fmt.Println(df)
	_ = df.WriteCSV(code + ".csv")

	p := plot.New()

	p.Title.Text = code
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"close", sliceToPoints(df.Col("close").DTypes()),
		"p1", sliceToPoints(df.Col("p1").DTypes()),
		"low", sliceToPoints(df.Col("low").DTypes()))
	if err != nil {
		log.Fatal(err)
	}
	//pngSize := 6
	if err = p.Save(10*vg.Inch, 10*vg.Inch, code+".png"); err != nil {
		log.Fatal(err)
	}
}
