package num

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
)

func Dot1D[T stat.Number](a, b []T) T {
	return __dot1d_go(a, b)
}

func __dot1d_go[T stat.Number](x, y []T) T {
	res := T(0)
	for i := 0; i < len(x); i++ {
		res += x[i] * y[i]
	}
	return res
}

// 3 x 3
func __dot2d_go[T stat.Number](a, b [][]T) [][]T {
	A := a
	B := b
	rLen := len(B)
	cLen := len(B[0])
	c := make([][]T, rLen)
	for i := 0; i < rLen; i++ {
		col := make([]T, cLen)
		for j := 0; j < cLen; j++ {
			for k := 0; k < rLen; k++ {
				col[j] += A[i][k] * B[k][j]
			}
		}
		c[i] = col
	}
	return c
}

func __align_go[T stat.Number](x [][]T, n int) [][]T {
	//if n < 3 {
	//	panic("lt 3")
	//}
	d := make([][]T, n)
	rLen := len(x)
	cLen := len(x[0])
	for i := 0; i < n; i++ {
		col := make([]T, n)
		for j := 0; j < n; j++ {
			col[j] = x[rLen-n+i][cLen-n+j]
		}
		d[i] = col
	}
	fmt.Println(d)
	return d
}

func Dot2D_V1[T stat.Number](a, b [][]T) [][]T {
	W := 3
	A := __align_go(a, W)
	B := __align_go(b, W)

	return __dot2d_go(A, B)
}

// Dot2D 二维矩阵点积
//
//	点积(dot)运算及简单应用 https://www.jianshu.com/p/482abac8798c
func Dot2D[T stat.Number](a, b [][]T) [][]T {
	A := a
	B := b
	rLen := len(A[0])
	cLen := len(B[0])
	xLen := __min_n_go(rLen, cLen)
	x := make([][]T, xLen)
	// 行
	for i := 0; i < xLen; i++ {
		col := make([]T, cLen)
		// 列
		for j := 0; j < cLen; j++ {
			for k := 0; k < rLen; k++ {
				col[j] += A[i][k] * B[k][j]
			}
		}
		x[i] = col
	}
	return x
}

// Dot 二维点积
func Dot[T stat.Number](a, b [][]T) [][]T {
	m, n := Shape[T](a)
	k, l := Shape[T](b)
	if n != k {
		panic("dot 2d a.rows<>b.cols")
	}
	//fmt.Println("m, n:", m, n)
	//fmt.Println("k, l:", k, l)
	x := make([][]T, m)
	// 行
	for i := 0; i < m; i++ {
		col := make([]T, l)
		// 列
		for c := 0; c < l; c++ {
			for r := 0; r < k; r++ {
				col[c] += a[i][r] * b[r][c]
			}
		}
		x[i] = col
	}
	return x
}

func Dot_v1[T stat.Number](a, b [][]T) [][]T {
	m, n := Shape[T](a)
	k, l := Shape[T](b)
	if n != k {
		panic("dot 2d vs 1d a.rows<>b.cols")
	}
	//fmt.Println("m, n:", m, n)
	//fmt.Println("k, l:", k, l)
	x := make([][]T, m)
	// 行
	for i := 0; i < m; i++ {
		col := make([]T, l)
		// 列
		for c := 0; c < l; c++ {
			for r := 0; r < k; r++ {
				col[c] += a[i][r] * b[r][c]
			}
		}
		x[i] = col
	}
	return x
}

// Dot2D1 二维矩阵和一维矩阵计算点积
func Dot2D1[T stat.Number](a [][]T, b []T) []T {
	B := [][]T{b}
	b1 := Transpose2D(B)
	//fmt.Println("Dot2D1: b1 =", b1)
	x1 := Dot[T](a, b1)
	//fmt.Println("Dot2D1: x1 =", x1)
	x2 := Transpose2D(x1)
	//fmt.Println("Dot2D1: x2 =", x2)
	return x2[0]
}

func Dot2D1_v2[T stat.Number](a [][]T, b []T) []T {
	m, n := Shape[T](a)
	k, l := Shape[T](b)
	if l < 1 {
		l = 1
	}

	fmt.Println("m, n:", m, n)
	fmt.Println("k, l:", k, l)
	x := make( /*[]*/ []T, m)
	// 行
	for i := 0; i < m; i++ {
		col := T(0)
		// 列
		for c := 0; c < l; c++ {
			for r := 0; r < k; r++ {
				col += a[i][r] * b[r]
			}
		}
		x[i] = col
	}
	return x
}

func Dot2D1_v1[T stat.Number](a [][]T, b []T) []T {
	A := a
	B := b
	rLen := len(A)
	cLen := len(B)
	xLen := __min_n_go(rLen, cLen)
	x := make([]T, xLen)
	pos := 0
	// 行
	for i := pos; i < rLen; i++ {
		// 列
		for k := 0; k < cLen; k++ {
			x[i] += A[i][k] * B[k]
		}
	}
	return x
}
