package main

import (
	"fmt"
	"testing"
)

func TestResultInfo_Headers(t *testing.T) {
	var t1 ResultInfo
	headers := t1.Headers()
	fmt.Println(headers)
	t1.Date = "2022-12-30"
	t1.Buy = 1.1
	t1.Code = "600600"
	t1.Name = "青岛啤酒"
	t1.Sell = 2.1
	fmt.Println(t1.Values())
}
