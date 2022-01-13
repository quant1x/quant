package index

import "testing"

func TestBoll(t *testing.T) {
	stockCode := "sh000001" // 上证指数
	//stockCode := "sh601600" // 中国铝业
	var f Formula
	f = NewBoll(20, 2)
	f.Load(stockCode)
}
