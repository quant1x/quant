package test

import (
	"fmt"
	"github.com/quant1x/quant/pandas/dataframe"
	"github.com/quant1x/quant/pandas/series"
	"io"
)

import _ "unsafe"

type Pandas struct {
	dataframe.DataFrame
	series.Series
}

func ReadCSV(r io.Reader, options ...dataframe.LoadOption) dataframe.DataFrame {
	return dataframe.ReadCSV(r, options...)
}

func EWM(d []float64, span int) []float64 {

	result := make([]float64, len(d))
	pervious := 0.00
	beta := float64(2) / float64(span+1)
	//beta = 0.6666667846851226 (span=5)
	// comass = (span - 1) / 2
	//beta := (span - 1.00) / 2.00
	comass := (span - 1) / 2
	beta = float64(comass) / float64(comass+1)
	//fmt.Println("beta=", beta)
	for i, v := range d {
		if i == 0 {
			//pervious = float64(1-beta) * v
			pervious = v
		} else {
			pervious = beta*pervious + float64(1-beta)*v
		}
		result[i] = pervious
		fmt.Println(pervious)
	}

	return result
}
