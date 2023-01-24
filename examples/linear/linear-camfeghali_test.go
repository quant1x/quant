package linear

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 最初的
func TestCamfeghaliT001(t *testing.T) {
	assert := assert.New(t)

	var x_train = []float64{1.0, 2.0}
	var y_train = []float64{300.0, 500.0}
	w_init := 0.0
	b_init := 0.0
	iterations := 10000
	tmp_alpha := 0.01

	w_final, b_final, _, _ := gradient_descent_single_var(x_train, y_train, w_init, b_init, tmp_alpha, iterations)

	assert.Equal(199.99285075131766, w_final, "")
	assert.Equal(100.011567727362, b_final, "")

}

// 进阶
func TestCamfeghaliT002(t *testing.T) {
	assert := assert.New(t)
	rl := LinearRegressionT01{}
	var x_train = []float64{1.0, 2.0}
	var y_train = []float64{300.0, 500.0}
	//w_init := 0.0
	//b_init := 0.0
	//iterations := 10000
	//tmp_alpha := 0.01
	w_final, b_final, _, _ := rl.Train(x_train, y_train)
	assert.Equal(199.99285075131766, w_final, "")
	assert.Equal(100.011567727362, b_final, "")
	a := rl.Predict(1)
	fmt.Println(a)
}

// 再进阶
func TestCamfeghaliT003(t *testing.T) {
	assert := assert.New(t)
	rl := LinearRegressionT01{}
	var x_train = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	var y_train = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	w_final, b_final, _, _ := rl.Train(x_train, y_train)
	assert.Equal(199.99285075131766, w_final, "")
	assert.Equal(100.011567727362, b_final, "")
	a := rl.Predict(7.0)
	fmt.Println("a =", a)
}

// 再进阶
func TestCamfeghaliT004(t *testing.T) {
	assert := assert.New(t)
	rl := LinearRegressionT01{}
	var x_train = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	var y_train = []float64{10.0, 20.0, 30.0, 40.0, 50.0}
	w_final, b_final, _, _ := rl.Train(x_train, y_train)
	assert.Equal(199.99285075131766, w_final, "")
	assert.Equal(100.011567727362, b_final, "")
	a := rl.Predict(70.0)
	fmt.Println("a =", a)
}

// 再进阶
func TestCamfeghaliT005(t *testing.T) {
	assert := assert.New(t)
	rl := LinearRegressionT01{}
	var x_train = []float64{10, 20, 30}
	var y_train = []float64{30.0, 40.0, 50}
	w_final, b_final, _, _ := rl.Train(x_train, y_train)
	assert.Equal(199.99285075131766, w_final, "")
	assert.Equal(100.011567727362, b_final, "")
	a := rl.Predict(10.0)
	fmt.Println("a =", a)
}
