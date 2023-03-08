package linear

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestW(t *testing.T) {
	code := "sh000905"
	code = "sz002528"
	//code = "sz000151"
	//code = "sz002564"
	//code = "sz002209"
	//code = "sz002951"
	//code = "sh000001"
	code = "sh600703"
	df := stock.KLine(code)
	//df = df.SelectRows(stat.RangeFinite(0, -5))
	fmt.Println(df)
	fmt.Printf("   证券代码: %s\n", code)
	W(df, true)
}
