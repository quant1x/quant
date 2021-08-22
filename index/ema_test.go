package index

import (
	"testing"
)

func TestEXMPA(t *testing.T) {
	stockCode := "sh000001"
	var f Formula
	f = &EMA{}
	f.Load(stockCode)
}

func TestLoadEma(t *testing.T) {
	stockCode := "sh000001"
	ema := LoadEma(stockCode)
	_ = ema
}
