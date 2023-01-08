package formula

// HHV 计算n周期内的flag的最大值
func HHV[V int64 | float64](slice []V, n int) V {
	return slice_universal(slice, n, func(a, b V) V {
		if a < b {
			return b
		}
		return a
	})
}

// LLV 计算n周期内的flag的最小值
func LLV[V int64 | float64](slice []V, n int) V {
	return slice_universal(slice, n, func(a, b V) V {
		if a > b {
			return b
		}
		return a
	})
}
