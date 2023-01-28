package linear

import (
	"gitee.com/quant1x/pandas"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/plot/plotter"
	"math"
)

// 根据条件修改原先值
func getTotal(s pandas.Series) pandas.Series {

	loadTime, _ := s.Val(3).(int)
	searchTime, _ := s.Val(4).(int)
	rAsTime, _ := s.Val(5).(int)

	res := loadTime + searchTime + rAsTime
	resF := float64(res) / float64(60)
	return pandas.Floats(resF)
}

func getDoc(s pandas.Series) pandas.Series {
	document, _ := s.Val(1).(float64)
	resF := float64(2*document) / float64(1000)
	return pandas.Floats(resF)
}

// 最小二乘法的线性拟合

// dataOptimize 数据优化和拟合函数
func dataOptimize(clsDF *pandas.DataFrame) (actPoints, expPoints plotter.XYs, fa, fb float64) {
	// 开始数据拟合

	// 实际观测点
	actPoints = plotter.XYs{}
	// N行数据产生N个点
	for i := 0; i < clsDF.Nrow(); i++ {
		document := clsDF.Elem(i, 1).Val().(float64)
		machine := clsDF.Elem(i, 2).Val().(int)
		val := clsDF.Elem(i, 3).Val().(float64)

		actPoints = append(actPoints, plotter.XY{
			X: float64(document) / float64(machine),
			Y: val,
		})
	}

	result, err := optimize.Minimize(optimize.Problem{
		Func: func(x []float64) float64 {
			if len(x) != 2 {
				panic("illegal x")
			}
			a := x[0]
			b := x[1]
			var sum float64
			for _, point := range actPoints {
				y := a*point.X + b
				sum += math.Abs(y - point.Y)
			}
			return sum
		},
	}, []float64{1, 1}, &optimize.Settings{}, &optimize.NelderMead{})
	if err != nil {
		panic(err)
	}

	// 最小二乘法拟合出来的k和b值
	fa, fb = result.X[0], result.X[1]
	expPoints = plotter.XYs{}
	for i := 0; i < clsDF.Nrow(); i++ {
		document := clsDF.Elem(i, 1).Val().(float64)
		machine := clsDF.Elem(i, 2).Val().(int)
		x := float64(document) / float64(machine)
		expPoints = append(expPoints, plotter.XY{
			X: x,
			Y: fa*float64(x) + fb,
		})
	}

	return
}
