package formula

// CROSS 金叉判断
// 当期 a > b 并且 前一期 a < b
// a和b对调就是死叉
func CROSS[V int64 | float64](a, b []V, n int) bool {
	ca := len(a)
	cb := len(b)
	if ca <= n || cb <= n {
		return false
	}
	// 首先判断前一周期是否 a < b
	fa1 := a[n+1]
	fb1 := b[n+1]
	if fa1 >= fb1 {
		return false
	} else {
		fa0 := a[n+0]
		fb0 := b[n+0]

		if fa0 <= fb0 {
			return false
		} else {
			return true
		}
	}
}
