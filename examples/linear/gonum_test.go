package linear

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestLinearRegression(t *testing.T) {
	u := mat.NewVecDense(3, []float64{1, 2, 3})
	fmt.Println("u: ", u)
	v := mat.NewVecDense(3, []float64{4, 5, 6})
	fmt.Println("v: ", v)
	w := mat.NewVecDense(3, nil)
	w.AddVec(u, v)
	fmt.Println("u + v: ", w)
	// Add u + alpha * v for some scalar alpha
	w.AddScaledVec(u, 2, v)
	fmt.Println("u + 2 * v: ", w)
	// Subtract v from u
	w.SubVec(u, v)
	fmt.Println("v - u: ", w)
	// Scale u by alpha
	w.ScaleVec(23, u)
	fmt.Println("u * 23: ", w)
	// Compute the dot product of u and v
	// Since float64’s don’t have a dot method, this is not done
	//inplace
	d := mat.Dot(u, v)
	fmt.Println("u dot v: ", d)
	// element-wise product
	w.MulElemVec(u, v)
	fmt.Println("u element-wise product v: ", w)
	// Find length of v
	l := v.Len()
	fmt.Println("Length of v: ", l)
}
