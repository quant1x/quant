package num

import (
	"fmt"
	"testing"
)

func Test___slice_shape_go(t *testing.T) {
	fmt.Println(__slice_shape_go[int](0))
	fmt.Println(__slice_shape_go[int]([]int{1, 2, 3}))
	fmt.Println(__slice_shape_go[int]([][]int{{1, 2, 3}, {4, 5, 6}}))
}
