package num

import (
	"fmt"
	"strconv"
	"testing"
)

func Test___index_go(t *testing.T) {
	f1 := 1.7763568394002505e-15
	fmt.Println(f1)
	fmt.Printf("1.7763568394002505e-15: %f\n", f1)
	fmt.Println(__float_exp_go(f1))
	f2 := -6.21724894e-15
	fmt.Println(__float_exp_go(f2))
	fmt.Printf("-6.21724894e-15: %f\n", f2)

	f2 = -0.07142857142857112
	fmt.Println(__float_exp_go(f2))
	s2 := strconv.FormatFloat(f2, 'E', -1, 64)
	fmt.Println(s2)
}
