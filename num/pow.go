package num

import (
	"gitee.com/quant1x/pandas/stat"
	"math"
)

func Pow[T stat.Number](v []T, n int) []T {
	x := make([]T, len(v))
	for i := 0; i < len(v); i++ {
		x[i] = __pow_go(v[i], n)
	}
	return x
}

func __pow_go[T stat.Number](x T, n int) T {
	y := math.Pow(float64(x), float64(n))
	return T(y)
}
