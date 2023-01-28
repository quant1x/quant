package tdx

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestGetKLine(t *testing.T) {
	data := GetKLine("000002", 0, 10)
	df := pandas.LoadStructs(data.List)
	df = df.Select([]string{"Open", "Close", "High", "Low", "Vol", "Amount", "DateTime"})
	//df = df.Select([]string{"Open", "Close", "High", "Low", "Vol", "DateTime"})
	_ = df.SetNames("open", "close", "high", "low", "vol", "amount", "date")
	df = df.Select([]string{"open", "close", "high", "low", "vol", "date"})
	date := df.Col("date")
	t1 := date.Map(func(element pandas.Element) pandas.Element {
		e := element.String()[0:10]
		element.Set(e)
		return element
	})
	df = df.Mutate(t1)
	fmt.Println(df)
}
