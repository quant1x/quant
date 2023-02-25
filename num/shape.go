package num

import "gitee.com/quant1x/pandas/stat"

func Shape[T stat.Number](x any) (r, c int) {
	return __slice_shape_go[T](x)
}

func __slice_shape_go[T stat.Number](x any) (r, c int) {
	switch vs := x.(type) {
	case T:
		return 0, 0
	case []T:
		return len(vs), 0
	case [][]T:
		r = len(vs)
		if r > 0 {
			c = len(vs[0])
		}
		return
	default:
		return -1, -1
	}

}
