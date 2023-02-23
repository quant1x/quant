package num

import (
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

// Dot2D 二维矩阵点积
//
//	点积(dot)运算及简单应用 https://www.jianshu.com/p/482abac8798c
func Dot2D[T stat.Number](a, b [][]T) [][]T {
	A := a[len(a)-2:]
	A = a
	B := b
	rLen := len(B)
	cLen := len(B[0])
	c := make([][]T, rLen)
	for i := 0; i < rLen; i++ {
		col := make([]T, cLen)
		for j := 0; j < cLen; j++ {
			//col[j] = A[i][cLen-2]*B[0][j] + A[i][cLen-1]*B[1][j]
			for k := 0; k < rLen; k++ {
				col[j] += A[i][k] * B[k][j]
			}
		}
		c[i] = col
	}
	return c
}

// Dot2D1 二维矩阵和一维矩阵计算点积
func Dot2D1[T stat.Number](a [][]T, b []T) []T {
	A := a[len(a)-1:]
	A = a
	B := b
	rLen := len(B)
	c := make([]T, rLen)
	for i := 0; i < rLen; i++ {
		//c[i] += A[0][i] * B[i]
		for k := 0; k < rLen; k++ {
			c[i] += A[i][k] * B[k]
		}
	}
	return c
}
