package utils

import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
	"math"
	"testing"
)

func TestComp(t *testing.T) {

	//支持比较,只支持 == 或 !=,比较并不是每一个元素都一样,2个数组比较,数组类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	//c := [5]int{1, 2, 3}
	fmt.Println("a == b", a == b)
	//fmt.Println("a == b", a >= c)

	//同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println("d = ", d)
}

func _slope(x1 int, y1 float64, x2 int, y2 float64) float64 {
	return (y2 - y1) / float64(x2-x1)
}

func TestSlope(t *testing.T) {
	a := math.Nextafter(2.0, 3.0)
	t.Logf("a = %f", a)

	x1 := 0
	y1 := 5.00
	x2 := 3
	y2 := 10.00
	xl := _slope(x1, y1, x2, y2)
	t.Logf("xl = %f", xl)

	x3 := 6
	y3 := xl*float64(x3-x2) + y2
	t.Logf("y3 = %f", y3)

	x_data := matrix.Zeros(1, 8)
	x_data.Set(0, 0, float64(1047656524340.0))
	x_data.Set(0, 1, float64(1058551622196.0))
	x_data.Set(0, 2, float64(1069443044404.0))
	x_data.Set(0, 3, float64(1080331486772.0))
	x_data.Set(0, 4, float64(1091218266164.0))
	x_data.Set(0, 5, float64(1102110036019.0))
	x_data.Set(0, 6, float64(1112998271027.0))
	x_data.Set(0, 7, float64(1123886124083.0))
	y_data := matrix.Zeros(1, 8)
	y_data.Set(0, 0, float64(139673476849.0))
	y_data.Set(0, 1, float64(139673468728.0))
	y_data.Set(0, 2, float64(139673460587.0))
	y_data.Set(0, 3, float64(139673452448.0))
	y_data.Set(0, 4, float64(139673444291.0))
	y_data.Set(0, 5, float64(139673436139.0))
	y_data.Set(0, 6, float64(139673427972.0))
	y_data.Set(0, 7, float64(139673419783.0))
	r := polyfit(x_data, y_data, 2)
	fmt.Println(r.Get(0, 0), r.Get(0, 1), r.Get(0, 2))
}

func pinv(a matrix.Matrix) matrix.Matrix {
	U, S, V, _ := a.DenseMatrix().SVD()
	ps := pinvofs(S)
	result := matrix.Product(V, ps)
	result = matrix.Product(result, matrix.Transpose(U))
	return result
}

func pinvofs(s matrix.MatrixRO) matrix.Matrix {
	result := matrix.Zeros(s.Rows(), s.Cols())
	for row := 0; row < s.Rows(); row++ {
		for col := 0; col < s.Cols(); col++ {
			num := s.Get(row, col)
			if num != 0.0 {
				result.Set(row, col, 1.0/num)
			}
		}
	}
	return result
}

func vander(x matrix.MatrixRO, order int) (V matrix.Matrix) {
	len := x.Cols()
	V = matrix.Zeros(len, order)
	for row := 0; row < V.Rows(); row++ {
		for col := 0; col < V.Cols(); col++ {
			V.Set(row, col, math.Pow(x.Get(0, row), float64(order-col-1)))
		}
	}
	return V
}

func polyfit(x, y matrix.Matrix, deg int) matrix.Matrix {
	order := deg + 1
	lhs := vander(x, order)
	rhs := matrix.Transpose(y)
	scale := matrix.Zeros(1, order)
	mir := scale.Arrays()
	for row := 0; row < lhs.Rows(); row++ {
		for col := 0; col < lhs.Cols(); col++ {
			mir[0][col] += math.Pow(lhs.Get(row, col), 2)
		}
	}

	for row := 0; row < scale.Rows(); row++ {
		for col := 0; col < scale.Cols(); col++ {
			mir[0][col] = math.Sqrt(mir[0][col])
		}
	}

	for row := 0; row < lhs.Rows(); row++ {
		for col := 0; col < lhs.Cols(); col++ {
			lhs.Set(row, col, lhs.Get(row, col)/scale.Get(0, col))
		}
	}
	c := matrix.Product(pinv(lhs), rhs)
	c = matrix.Transpose(c).DenseMatrix()
	for row := 0; row < c.Rows(); row++ {
		for col := 0; col < c.Cols(); col++ {
			c.Set(row, col, c.Get(row, col)/scale.Get(0, col))
		}
	}
	return c
}
