package linear

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestCurveRegression(t *testing.T) {
	df := stock.KLine("sz002528")
	df = df.Subset(0, df.Nrow()-1)
	fmt.Println(df)
	V := df.Col("open")
	N := 5
	d := CurveRegression(V, N)
	fmt.Println(d)
}
