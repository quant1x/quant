package index

import (
	"testing"
)

func TestRSI_Load(t *testing.T) {
	code := "sh000001"
	var f Formula
	f = &RSI{}
	f.Load(code)
}
