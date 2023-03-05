package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestF89K(t *testing.T) {
	df := stock.KLine("sh600496")
	fmt.Println(df)
	df1 := F89K(df, 89)
	fmt.Println(df1)
}
