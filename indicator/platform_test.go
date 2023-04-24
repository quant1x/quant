package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestPlatform(t *testing.T) {
	code := "600703.sh"
	code = "603789.sh"
	code = "sz000506"
	code = "sh603367"
	//code = "sz002275"
	code = "sz002665"
	code = "sz002528"
	//code = "sz000892"
	//code = "sz000905"
	code = "sh600641"
	code = "sh688031"
	code = "sz000988"
	code = "sh600105"
	code = "sz002292"
	code = "sh600354"
	df := stock.KLine(code)
	df1 := Platform(df)
	fmt.Println(df1)
	_ = df1.WriteCSV("t02.csv")
}
