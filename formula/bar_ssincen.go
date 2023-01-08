package formula

// 指标计算接口
type BarHandler = func(a, b float64) bool

func slice_ssincen[V int64 | float64](slice []V, n int, iterator func(a, b V) bool) int {
	count := len(slice)
	if count < n {
		return -1
	}
	var (
		ret int = -1
		//inited bool = false
	)
	for i := 0; i < n; i++ {
		// 对比成交量至少需要2天的数据
		if i < 1 {
			continue
		}
		// CompVal
		v1 := slice[i-1]
		v2 := slice[i+0]

		bRet := iterator(v1, v2)
		if bRet {
			//fmt.Printf("a=%v, b=%v, ret=%v\n", v1, v2, bRet)
			ret = i
			break
		}
	}
	return ret
}

// BARSSINCEN N周期内第一次X不为0到现在的周期数,N为常量
func BARSSINCEN[V int64 | float64](slice []V, n int, iterator func(a, b V) bool) int {
	return slice_ssincen(slice, n, iterator)
}
