package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestCDTD(t *testing.T) {
	code := "002528.sz"
	df := stock.KLine(code)
	fmt.Println(df)
	df1 := CDTD(df)
	fmt.Println(df1)
}
