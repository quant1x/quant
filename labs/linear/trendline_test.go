package linear

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"testing"
)

func TestTrendLine(t *testing.T) {
	code := "sh000905"
	code = "sz002528"
	//code = "sz002322"
	df := stock.KLine(code)
	df = TrendLine(df)
	fmt.Println(df)

	p := plot.New()
	p.Title.Text = code + "  /  " + df.Col("date").IndexOf(-1).(string)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	df = df.Subset(df.Nrow()-MaximumTrendPeriod, df.Nrow())
	err := plotutil.AddLinePoints(p,
		"close", sliceToPoints(df.Col("close").DTypes()),
		"support", sliceToPoints(df.Col("zc").DTypes()),
		"high", sliceToPoints(df.Col("high").DTypes()),
		"pressure", sliceToPoints(df.Col("yl").DTypes()))
	if err != nil {
		log.Fatal(err)
	}
	//pngSize := 6
	if err = p.Save(10*vg.Inch, 10*vg.Inch, code+".png"); err != nil {
		log.Fatal(err)
	}
}

func TestCrossTrend(t *testing.T) {
	code := "sh000905"
	code = "sz002528"
	//code = "sz002322"
	code = "sh600018"
	code = "sh603130"
	df := stock.KLine(code)
	df = CrossTrend(df)
	fmt.Println(df)

	p := plot.New()
	p.Title.Text = code + "  /  " + df.Col("date").IndexOf(-1).(string)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	df = df.Subset(df.Nrow()-MaximumTrendPeriod, df.Nrow())
	err := plotutil.AddLinePoints(p,
		"close", sliceToPoints(df.Col("close").DTypes()),
		"support", sliceToPoints(df.Col("zc").DTypes()),
		"high", sliceToPoints(df.Col("high").DTypes()),
		"pressure", sliceToPoints(df.Col("yl").DTypes()))
	if err != nil {
		log.Fatal(err)
	}
	//pngSize := 6
	if err = p.Save(10*vg.Inch, 10*vg.Inch, code+".png"); err != nil {
		log.Fatal(err)
	}
}
