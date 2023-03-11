package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestPlatform(t *testing.T) {
	code := "600703.sh"
	code = "000001.sh"
	df := stock.KLine(code)
	df1 := Platform(df)
	fmt.Println(df1)
	df1.WriteCSV("t.csv")
}
