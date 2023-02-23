package num

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestC_(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := stat.Mul(a, a)
	c := []int{1, 1, 1, 1, 1}
	fmt.Println(C_[int](b, a, c))
}

func TestT_(t *testing.T) {
	df := stock.KLine("002528")
	df = df.Subset(0, df.Nrow()-1)
	fmt.Println(df)
	V := df.Col("low")
	N := 3
	y := V.Select(stat.RangeFinite(-N)).DTypes()
	x := Arange[float64](1, float64(N)+1, 1)
	t1 := Pow[float64](x, 2)
	t2 := x
	t3 := Ones[float64](x)

	A := C_(t1, t2, t3)
	T := T_(A)
	fmt.Println("A =", A)
	fmt.Println("T =", T)

	w0 := Dot2D[float64](T, A)
	fmt.Println("w0 =", w0)
	w1 := __inv(w0)
	fmt.Println("w1 =", w1)

	w2 := Dot2D(w1, T)
	fmt.Println("w2 =", w2)

	w3 := Dot2D1[float64](w2, y)
	fmt.Println("w3 =", w3)

	W := w3
	d1 := Arange[float64](1, float64(N)+2, 1)

	d21 := Pow(d1, 2)
	d2 := stat.NDArray[float64](d21).Mul(W[0])

	d3 := stat.NDArray[float64](d1).Mul(W[1]).Add(W[2])

	D := d2.Add(d3)
	fmt.Println("D =", D)
}
