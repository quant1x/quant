package main

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestCheckBlock(t *testing.T) {
	df := TopBlock(0)
	fmt.Println(df)
}

func Test_scanBlock(t *testing.T) {
	data := scanBlock(0)
	df := pandas.LoadStructs(data)
	fmt.Println(df)
}
