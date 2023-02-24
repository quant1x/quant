package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestF89K(t *testing.T) {
	df := stock.KLine("sh600090")
	//csv := "~/.quant1x/data/cn/600520.csv"
	//df = pandas.ReadCSV(csv)
	////df.SetNames("日期/开盘/收盘/最高/最低/成交量/成交额/振幅/涨跌幅/涨跌额/换手率")
	//_ = df.SetNames("date", "open", "close", "high", "low", "volume", "amount", "zf", "zdf", "zde", "hsl")
	fmt.Println(df)
	df1 := F89K(df, 89)
	fmt.Println(df1)
}
