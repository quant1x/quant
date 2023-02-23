package linear

import "gonum.org/v1/gonum/stat"

// 这一组功能里面会收敛一部分gonum.org/v1/gonum (https://github.com/gonum/gonum.git)的功能

// LinearRegression 线性回归
func LinearRegression(x, y, weights []float64, origin bool) (alpha, beta float64) {
	return stat.LinearRegression(x, y, weights, origin)
}
