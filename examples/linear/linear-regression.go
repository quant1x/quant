package linear

import (
	"context"
	"fmt"
	"github.com/erni27/regression"
	"github.com/erni27/regression/linear"
	"github.com/erni27/regression/options"
	"github.com/mymmsc/gox/logger"
	"log"
)

func LinearT1() {
	// Creates regression options with:
	// 1) Learning rate equals 1e-8.
	// 2) Batch gradient descent variant.
	// 3) Iterative convergance with number of iterations equal 1000.
	opt := options.WithIterativeConvergence(1e-8, options.Batch, 1000)
	// Initialize linear regression with gradient descent (numerical approach).
	r := linear.WithGradientDescent(opt)
	// Create design matrix as a 2D slice.
	x := [][]float64{
		{2104, 3},
		{1600, 3},
		{2400, 3},
		{1416, 2},
		{3000, 4},
		{1985, 4},
		{1534, 3},
		{1427, 3},
		{1380, 3},
		{1494, 3},
	}
	// Create target vector as a slice.
	y := []float64{399900, 329900, 369000, 232000, 539900, 299900, 314900, 198999, 212000, 242500}
	//y := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ctx := context.Background()
	// Run linear regression.
	m, err := r.Run(ctx, regression.TrainingSet{X: x, Y: y})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	acc := m.Accuracy()
	fmt.Printf("Accuracy: %f.\n", acc)
	coeffs := m.Coefficients()
	fmt.Printf("Coefficients: %v.\n", coeffs)
	// Do a predicition for a new input feature vector.
	in := []float64{2550, 4}
	p, err := m.Predict(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("For vector %v, predicted value equals %f.\n", in, p)
}

// 最简单的最小二乘法
func LeastSquares(x []float64, y []float64) (slope float64, intercept float64) {
	// x是横坐标数据,y是纵坐标数据
	// a是斜率，b是截距
	xi := float64(0)
	x2 := float64(0)
	yi := float64(0)
	xy := float64(0)

	if len(x) != len(y) {
		logger.Debugf("最小二乘时，两数组长度不一致!")
	} else {
		xLen := len(x)
		length := float64(xLen)
		window := 5
		if xLen <= window {
			window = xLen
		}
		for i := xLen - window; i < xLen; i++ {
			xi += x[i]
			x2 += x[i] * x[i]
			yi += y[i]
			xy += x[i] * y[i]
		}
		slope = (yi*xi - xy*length) / (xi*xi - x2*length)
		intercept = (yi*x2 - xy*xi) / (x2*length - xi*xi)
	}
	return
}

func Predict(y, slope, intercept float64) float64 {
	return y*slope + intercept
}
