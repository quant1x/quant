package num

import (
	"gitee.com/quant1x/pandas/stat"
)

// Arange Return evenly spaced values within a given interval.
//
//	返回给定间隔内的等间距值
func Arange[T stat.Number](start T, end T, argv ...T) []T {
	step := T(1)
	if len(argv) > 0 {
		step = argv[0]
	}
	x := make([]T, 0)
	for i := start; i < end; i += step {
		x = append(x, start)
		start += T(step)
	}

	return x
}
