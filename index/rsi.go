package index

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/formula"
	"math"
)

type RSI struct {
	*cache.DataFrame
	RSI1 []float64
	N1   int
	RSI2 []float64
	N2   int
	M    int
}

func (this *RSI) Len() int {
	return this.Length
}

func (this *RSI) Data() interface{} {
	return this.DataFrame
}

// Load
// LC:=REF(CLOSE,1);
// RSI1:SMA(MAX(CLOSE-LC,0),N1,1)/SMA(ABS(CLOSE-LC),N1,1)*100;
// RSI2:SMA(MAX(CLOSE-LC,0),N2,1)/SMA(ABS(CLOSE-LC),N2,1)*100;
// RSI3:SMA(MAX(CLOSE-LC,0),N3,1)/SMA(ABS(CLOSE-LC),N3,1)*100;
func (this *RSI) Load(code string) error {
	{
		this.N1 = 6
		this.N2 = 12
		this.M = 1
	}

	df := cache.LoadDataFrame(code)
	this.DataFrame = df
	if df == nil {
		return ErrCode
	} else if this.Length < 1 {
		return ErrData
	}

	count := this.Length
	var arr1, arr2 []float64
	for i := 0; i < count; i++ {
		if i < 1 {
			arr1 = append(arr1, 0.00)
			arr2 = append(arr2, 0.00)
			//this.RSI1 = append(this.RSI1, 0.00)
			//this.RSI2 = append(this.RSI2, 0.00)
		}
		end := i + 1
		//_date, _open, _close, _high, _low, _volume := this.DataFrame.Offset(end)
		_close := this.Close[:end]

		lc := formula.REF(_close, 1)
		cc := formula.REF(_close, 0)
		cx := cc - lc
		e1 := math.Max(cx, 0)
		arr1 = append(arr1, e1)
		e2 := math.Abs(cx)
		arr2 = append(arr2, e2)

		t1 := 100 * formula.SMA(arr1, this.N1, this.M) / formula.SMA(arr2, this.N1, this.M)
		this.RSI1 = append(this.RSI1, t1)

		t2 := 100 * formula.SMA(arr1, this.N2, this.M) / formula.SMA(arr2, this.N2, this.M)
		this.RSI2 = append(this.RSI2, t2)

		// 输出最后2组数据
		if category.DEBUG && count < i+3 {
			fmt.Printf("day: %s, RSI1: %.2f, RSI2: %.2f\n", this.Date[i], this.RSI1[i], this.RSI2[i])
		}
	}
	return nil
}
