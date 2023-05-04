package models

import (
	"fmt"
	"gitee.com/quant1x/data/security"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestCheckBlock(t *testing.T) {
	pbarIndex := 0
	data := TopBlock(&pbarIndex)
	df := pandas.LoadStructs(data)
	fmt.Println(df)
}

func Test_scanBlock(t *testing.T) {
	pbarIndex := 0
	data := scanBlock(&pbarIndex, security.BK_HANGYE)
	df := pandas.LoadStructs(data)
	fmt.Println(df)
}
