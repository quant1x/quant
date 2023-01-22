// Package linear Go语言实现线性回归
// https://blog.csdn.net/Deep___Learning/article/details/107432855
package linear

import (
	"fmt"
	G "gorgonia.org/gorgonia"
	. "gorgonia.org/tensor"
	"math/rand"
	"runtime"
	"time"
)

const vecSize = 1000000

// 实现线性回归
func Tmain1() {
	var m, c G.Value

	m, c = LinearRegression(Float32, 500)
	fmt.Printf("Float32: y = %3.2fx + %3.3f", m, c)

	fmt.Println()

	m, c = LinearRegression(Float64, 500)
	fmt.Printf("Float64: y = %3.2fx + %3.3f", m, c)
}

// 构造数据
func xy(dt Dtype) (x Tensor, y Tensor) {
	var xBack, yBack interface{}
	switch dt {
	case Float32:
		xBack = Range(Float32, 1, vecSize+1).([]float32)
		yBackC := Range(Float32, 1, vecSize+1).([]float32)
		for i, v := range yBackC {
			yBackC[i] = v*2 + rand.Float32()
		}
		yBack = yBackC
	case Float64:
		xBack = Range(Float64, 1, vecSize+1).([]float64)
		yBackC := Range(Float64, 1, vecSize+1).([]float64)
		for i, v := range yBackC {
			yBackC[i] = v*2 + rand.Float64()
		}
		yBack = yBackC
	}
	return New(WithBacking(xBack), WithShape(vecSize)), New(WithBacking(yBack), WithShape(vecSize))
}

// 产生随机数
func random(dt Dtype) interface{} {
	rand.Seed(time.Now().UnixNano())
	switch dt {
	case Float32:
		return rand.Float32()
	case Float64:
		return rand.Float64()
	default:
		panic("错误的类型")
	}
}

func LinearRegressionSetup(dt Dtype) (m, c *G.Node, machine G.VM) {
	var xT, yT G.Value
	xT, yT = xy(dt)

	g := G.NewGraph()
	x := G.NewVector(g, dt, G.WithShape(vecSize), G.WithName("x"), G.WithValue(xT))
	y := G.NewVector(g, dt, G.WithShape(vecSize), G.WithName("y"), G.WithValue(yT))

	m = G.NewScalar(g, dt, G.WithName("m"), G.WithValue(random(dt)))
	c = G.NewScalar(g, dt, G.WithName("c"), G.WithValue(random(dt)))

	// y = m * x + c
	pred := G.Must(G.Add(G.Must(G.Mul(x, m)), c))

	// 使得均方差最小
	se := G.Must(G.Square(G.Must(G.Sub(pred, y))))
	cost := G.Must(G.Mean(se))
	if _, err := G.Grad(cost, m, c); err != nil {
		fmt.Println(err)
	}
	machine = G.NewTapeMachine(g, G.BindDualValues(m, c))
	return m, c, machine
}

func LinearRegressionRun(m, c *G.Node, machine G.VM, iter int, autoClean bool) (retM, retC G.Value) {
	if autoClean {
		defer machine.Close()
	}
	model := []G.ValueGrad{m, c}
	solver := G.NewVanillaSolver(G.WithLearnRate(0.001), G.WithClip(5))
	var err error
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Println(i, err)
			break
		}
		if err = solver.Step(model); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d: m = %3.2f, c = %3.2f\n", i+1, m.Value(), c.Value())
		machine.Reset()
	}
	return m.Value(), c.Value()
}

// 线性回归
func LinearRegression(Float Dtype, iter int) (retM, retC G.Value) {
	m, c, machine := LinearRegressionSetup(Float)
	defer runtime.GC()
	return LinearRegressionRun(m, c, machine, iter, true)
}
