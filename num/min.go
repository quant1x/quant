package num

import (
	"gitee.com/quant1x/pandas/stat"
	"math"
)

func __min_n_go[T stat.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func __float_exp_go(f float64) (frac float64, exp int) {
	return math.Frexp(f)
}
