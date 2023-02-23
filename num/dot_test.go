package num

import (
	"fmt"
	"testing"
)

func TestDot2D(t *testing.T) {
	//a := [][]int{{1, 2}, {3, 7}}
	//b := [][]int{{4, 3}, {5, 0}}
	//
	//c := Dot2D(a, b)
	//fmt.Println(c)

	A := [][]int{{1, 4, 9}, {1, 2, 3}, {1, 1, 1}}
	fmt.Println("A =", A)
	B := [][]int{{1, 1, 1}, {4, 2, 1}, {9, 3, 1}}
	fmt.Println("B =", B)
	C := Dot2D(A, B)
	fmt.Println("C =", C)
}
