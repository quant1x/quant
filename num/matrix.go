package num

import "gitee.com/quant1x/pandas/stat"

// C_ Translates slice objects to concatenation along the second axis.
//
//	沿第二个轴将切片对象转换为串联
func C_[T stat.Number](a, b, c []T) [][]T {
	length := len(a)
	cLen := 3 // a,b,c
	rows := make([][]T, length)
	for i := 0; i < length; i++ {
		col := make([]T, cLen)
		col[0] = a[i]
		col[1] = b[i]
		col[2] = c[i]
		rows[i] = col
	}
	return rows
}

// T_ 矩阵转置
func T_[T stat.Number](x [][]T) [][]T {
	length := len(x[0])
	cLen := len(x)
	rows := make([][]T, length)
	for i := 0; i < length; i++ {
		col := make([]T, cLen)
		for j := 0; j < cLen; j++ {
			col[j] = x[j][i]
		}
		rows[i] = col
	}
	return rows
}

// 计算矩阵的（乘法）逆
//
//	Compute the (multiplicative) inverse of a matrix.
//
//	Given a square matrix `a`, return the matrix `ainv` satisfying
//	``dot(a, ainv) = dot(ainv, a) = eye(a.shape[0])``.
func __inv(a [][]float64) [][]float64 {
	var n = len(a)

	// Create augmented matrix
	var augmented = make([][]float64, n)
	for i := range augmented {
		augmented[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			augmented[i][j] = a[i][j]
		}
	}
	for i := 0; i < n; i++ {
		augmented[i][i+n] = 1
	}

	// Perform Gaussian elimination
	for i := 0; i < n; i++ {
		var pivot = augmented[i][i]
		for j := i + 1; j < n; j++ {
			var factor = augmented[j][i] / pivot
			for k := i; k < 2*n; k++ {
				augmented[j][k] -= factor * augmented[i][k]
			}
		}
	}

	// Perform back-substitution
	for i := n - 1; i >= 0; i-- {
		var pivot = augmented[i][i]
		for j := i - 1; j >= 0; j-- {
			var factor = augmented[j][i] / pivot
			for k := i; k < 2*n; k++ {
				augmented[j][k] -= factor * augmented[i][k]
			}
		}
	}

	// Normalize rows
	for i := 0; i < n; i++ {
		var pivot = augmented[i][i]
		for j := 0; j < 2*n; j++ {
			augmented[i][j] /= pivot
		}
	}

	// Extract inverse from augmented matrix
	var inverse = make([][]float64, n)
	for i := range inverse {
		inverse[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			inverse[i][j] = augmented[i][j+n]
		}
	}

	return inverse
}
