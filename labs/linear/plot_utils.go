package linear

import "gonum.org/v1/plot/plotter"

func sliceToPoints(x []float64) plotter.XYs {
	n := len(x)
	points := make(plotter.XYs, n)
	for i := range points {
		points[i].X = float64(i)
		points[i].Y = x[i]
	}

	return points
}
