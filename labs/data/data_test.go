package data

import (
	"fmt"
	"gitee.com/quant1x/data/category/trading"
	"gitee.com/quant1x/data/stock"
	"testing"
)

func TestTick(t *testing.T) {
	symbol := "sz002528"
	//symbol = "sz000638"
	//symbol = "sz000736"
	//symbol = "sz000615"
	symbol = "sh600030"
	dates := trading.TradeRange("2023-02-20", "2023-03-03")
	df := stock.Tick(symbol, dates)
	fmt.Println(df)
}
