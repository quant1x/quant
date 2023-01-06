package index

import "testing"

func TestLoad89k(t *testing.T) {
	code := "sh000001"

	var f Formula
	f = &K89{}
	f.Load(code)
}
