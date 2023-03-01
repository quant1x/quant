package num

import (
	"math"
	"testing"
)

func TestSlope(t *testing.T) {
	a := math.Nextafter(2.0, 3.0)
	t.Logf("a = %f", a)

	x1 := 0
	y1 := 5.00
	x2 := 3
	y2 := 10.00
	xl := Slope(x1, y1, x2, y2)
	t.Logf("xl = %f", xl)

	x3 := 6
	y3 := xl*float64(x3-x2) + y2
	t.Logf("y3 = %f", y3)
}
