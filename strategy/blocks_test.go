package main

import (
	"fmt"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestCheckBlock(t *testing.T) {
	data := TopBlock(0)
	df := pandas.LoadStructs(data)
	fmt.Println(df)
}

func Test_scanBlock(t *testing.T) {
	data := scanBlock(0, security.BK_HANGYE)
	df := pandas.LoadStructs(data)
	fmt.Println(df)
}
