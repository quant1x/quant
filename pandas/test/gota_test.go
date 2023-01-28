package test

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"github.com/quant1x/quant/cache"
	"strings"
	"testing"
)

func TestCsv(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
	df.SetNames("a", "b", "c", "d", "e")
	s1 := df.Col("d")
	fmt.Println(s1)

	fp := cache.GetCache("sz000002")
	df = pandas.ReadCSV(fp)
	fmt.Println(df)
	closes := df.Col("Close")
	closes.Median()
	ma5 := closes.Rolling(5).Mean()
	pandas.NewSeries(closes, pandas.Float, "")
	fmt.Println(ma5)
}
