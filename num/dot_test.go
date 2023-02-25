package num

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestDot2D(t *testing.T) {
	// https://blog.csdn.net/llittleSun/article/details/115045660
	//
	//a := [][]int{{1, 2}, {3, 7}}
	//b := [][]int{{4, 3}, {5, 0}}
	//
	//c := Dot2D(a, b)
	//fmt.Println(c)

	A := [][]int{{1, 4, 9}, {1, 2, 3}, {1, 1, 1}}
	fmt.Println("A =", A)
	B := [][]int{{1, 1, 1}, {4, 2, 1}, {9, 3, 1}}
	fmt.Println("B =", B)
	C := Dot2D(A, B)
	fmt.Println("C =", C)
}

func TestDot2D1(t *testing.T) {
	df := stock.KLine("002528")
	df = df.Subset(0, df.Nrow()-1)
	fmt.Println(df)
	V := df.Col("low")
	N := 5
	y := V.Select(stat.RangeFinite(-N)).DTypes()
	fmt.Println("y =", y)
	x := Arange[float64](1, float64(N)+1, 1)
	t1 := Pow[float64](x, 2)
	t2 := x
	t3 := Ones[float64](x)

	A := Concat1D(t1, t2, t3)
	T := Transpose2D(A)
	fmt.Println("A =", A)
	fmt.Println("T =", T)

	w0 := Dot2D[float64](T, A)
	fmt.Println("w0 =", w0)
	w1 := Inverse(w0)
	fmt.Println("w1 =", w1)

	w2 := Dot2D(w1, T)
	fmt.Println("w2 =", w2)

	W := Dot2D1[float64](w2, y)
	fmt.Println("W =", W)

	d1 := Arange[float64](1, float64(N)+2, 1)
	fmt.Println("d1 =", d1)

	d21 := Pow(d1, 2)
	d2 := stat.NDArray[float64](d21).Mul(W[0])

	d3 := stat.NDArray[float64](d1).Mul(W[1]).Add(W[2])

	D := d2.Add(d3)
	fmt.Println("D =", D)
}

func TestDot(t *testing.T) {
	A := [][]int{{1, 4, 9}, {1, 2, 3}, {1, 1, 1}}
	A = [][]int{{1, 4, 9, 16, 25}, {1, 2, 3, 4, 5}, {1, 1, 1, 1, 1}}
	B := Transpose2D(A)
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	C := Dot(A, B)
	fmt.Println("C =", C)
}
