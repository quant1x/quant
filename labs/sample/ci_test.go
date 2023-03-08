package sample

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestConfidenceInterval(t *testing.T) {
	code := "688351.sh"
	df := stock.KLine(code)
	fmt.Println(df)
	df = ConfidenceInterval(df, 5)
	fmt.Println(df)
}
