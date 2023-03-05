package sample

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestConfidenceInterval(t *testing.T) {
	code := "002528.sz"
	df := stock.KLine(code)
	df = ConfidenceInterval(df, 5)
	fmt.Println(df)
}
