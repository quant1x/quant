package index

import "testing"

func TestLoadMacd(t *testing.T) {
	stockCode := "sh000001"
	macd := LoadMacd(stockCode)
	_ = macd
}
