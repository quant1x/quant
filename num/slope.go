package num

// Slope 计算斜率
func Slope(x1 int, y1 float64, x2 int, y2 float64) float64 {
	return (y2 - y1) / float64(x2-x1)
}

// TriangleBevel 三角形斜边
func TriangleBevel(slope float64, x1 int, y1 float64, xn int) float64 {
	return slope*float64(xn-x1) + y1
}
