package index

import "testing"

func TestLoad89k(t *testing.T) {
	code := "sz002423"

	var f Formula
	f = &K89{}
	f.Load(code)
}
