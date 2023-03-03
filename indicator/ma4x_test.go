package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestMA4X(t *testing.T) {
	code := "000736.sz"
	df := stock.KLine(code)
	fmt.Println(df)
	df1 := MA4X(df, 5)
	fmt.Println(df1)
}
