package index

import "testing"

func TestLoadMa(t *testing.T) {
	code := "sh000001"

	var f Formula
	f = &MA{}
	f.Load(code)
}