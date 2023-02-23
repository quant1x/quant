package num

import (
	"gitee.com/quant1x/pandas/stat"
)

// Ones v -> shape
func Ones[T stat.Number](v []T) []T {
	return stat.Repeat[T](T(1), len(v))
}
