package fft

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"math"
	"testing"
	"time"
)

func Test_fft(t *testing.T) {
	code := "002483.sz"
	df := stock.KLine(code)
	fmt.Println(df)
	t1 := df.Col("close").DTypes()
	const n0 = 1024
	var i int
	var pr [n0]float64
	for i = 0; i < n0; i++ { //生成输入信号
		t1[i] = t1[i] * 0.001
		pr[i] = 1.2 + 2.7*math.Cos(2*math.Pi*33*t1[i]) + 5*math.Cos(2*math.Pi*200*t1[i]+math.Pi/2)

	}

	now := time.Now()

	v1_fft(pr[0:n0], n0) //调用FFT函数
	for i = 0; i < n0; i++ {
		fmt.Printf("%v\t%f\n", i, pr[i]) //输出结果
	}
	fmt.Print(now, "\n", time.Now())
}

func Test_FFT(t *testing.T) {
	data := make([]complex128, 32)
	for i := range data {
		// Fill data
		data[i] = complex(float64(i*2)/float64(32), 0)
	}
	fmt.Println(data)
	Fft(data)
	fmt.Println(data)
}

func Test_FFT2(t *testing.T) {
	x0 := []float64{
		5,
		32,
		38,
		-33,
		-19,
		-10,
		1,
		-8,
		-20,
		10,
		-1,
		4,
		11,
		-1,
		-7,
		-2,
	}
	n := len(x0)
	x := make([]complex128, n)
	for k := 0; k < n; k++ {
		x[k] = complex(x0[k], 0.0)
	}

	y := FFT(x, n)
	z := IFFT(y, n)
	f := FFTFreq(n, 1)

	fmt.Println(" K   DATA  FOURIER TRANSFORM  INVERSE TRANSFORM  FREQUENCY")
	for k := 0; k < n; k++ {
		fmt.Printf("%2d %6.1f  %8.3f%8.3f   %8.3f%8.3f   %8.3f\n",
			k, x0[k], real(y[k]), imag(y[k]), real(z[k]), imag(z[k]), f[k])
	}
}
