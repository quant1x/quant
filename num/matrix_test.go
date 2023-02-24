package num

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestConcat1D(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := stat.Mul(a, a)
	c := []int{1, 1, 1, 1, 1}
	fmt.Println(Concat1D[int](b, a, c))
}

func TestTranspose2D(t *testing.T) {
	N := 5
	x := Arange[float64](1, float64(N)+1, 1)

	t1 := Pow[float64](x, 2)
	t2 := x
	t3 := Ones[float64](x)

	A := Concat1D(t1, t2, t3)
	T := Transpose2D(A)
	fmt.Println("A =", A)
	fmt.Println("T =", T)
}
