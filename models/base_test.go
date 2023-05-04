package models

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"testing"
)

func TestResultInfo_Headers(t *testing.T) {
	var t1 ResultInfo
	headers := t1.Headers()
	fmt.Println(headers)
	t1.Date = "2022-12-30"
	t1.Buy = 1.1
	t1.Code = "600600"
	t1.Name = "青岛啤酒"
	t1.Sell = 2.1
	fmt.Println(t1.Values())
}

func TestVolume(t *testing.T) {
	code := "sz000506"
	//code = "sh600477"
	//code = "sz002280"
	code = "sz002528"
	code = "sz002292"
	//code = "sz002792"
	//code = "sh880473"
	//code = "sz000686"
	//code = "sh688041"
	code = "sz002665"
	N := 100
	DIFF := 0
	df := stock.KLine(code)
	fmt.Println(df)
	dates := df.Col("date").Select(stat.RangeFinite(-N-DIFF, -1-DIFF)).Values().([]string)
	df = stock.Tick(code, dates)
	fmt.Println(df)
	plt := plot.New()
	start := df.Col("date").IndexOf(0).(string)
	end := df.Col("date").IndexOf(-1).(string)
	plt.Title.Text = fmt.Sprintf("%s", code)
	plt.X.Label.Text = "X: " + start + "-" + end
	plt.Y.Label.Text = "Y"
	length := N
	df2 := df.Subset(df.Nrow()-length, df.Nrow())
	err2 := plotutil.AddLinePoints(plt,
		"buy.inflow", SliceToPoints(df2.Col("iv").DTypes()),
		"buy.incr", SliceToPoints(df2.Col("bv").DTypes()),
		"sell.incr", SliceToPoints(df2.Col("sv").DTypes()))
	if err2 != nil {
		log.Fatal(err2)
	}
	//pngSize := 6
	if err3 := plt.Save(6*vg.Inch, 6*vg.Inch, code+".png"); err3 != nil {
		log.Fatal(err3)
	}
}

func TestResultInfo_DetectVolume(t *testing.T) {
	result := ResultInfo{Code: "sh600641"}
	ok := result.DetectVolume()
	fmt.Println(ok)
}
