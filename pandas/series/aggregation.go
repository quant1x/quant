package series

import (
	"github.com/quant1x/quant/pandas/series/math"
)

// AggregateFunc is applied aggregation function.
type AggregateFunc func(data Data) DType

// Mean returns mean of data's values.
func Mean(data Data) DType {
	var (
		count int
		sum   DType
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		sum += v
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return sum / DType(len(items))
}

// Sum returns sum of data's values.
func Sum(data Data) DType {
	var (
		sum   DType
		count int
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		sum += v
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return sum
}

// Min returns minimum value.
func Min(data Data) DType {
	var (
		min   DType = maxFloat
		count int
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		if v < min {
			min = v
		}
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return min
}

// Max returns maximum value.
func Max(data Data) DType {
	var (
		max   DType = -maxFloat
		count int
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		if v > max {
			max = v
		}
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return max
}

// Median returns median value of series.
// Linear interpolation is used for odd length.
func Median(data Data) DType {
	values := data.Values()

	if len(values) == 0 {
		return math.NaN()
	}

	if len(values) == 1 {
		return values[0]
	}

	if len(values)%2 == 0 {
		i := len(values) / 2
		return (values[i-1] + values[i]) / 2
	}

	return values[len(values)/2]
}

// Argmin returns offset of the smallest value of series data.
// If the minimum is achieved in multiple locations, the first row position is returned.
func Argmin(data Data) int {
	var (
		min   DType = maxFloat
		pos   int
		items = data.Values()
	)
	for i, v := range items {
		if IsNA(v) {
			continue
		}
		if v < min {
			min = v
			pos = i
		}
	}
	return pos
}

// Argmax returns offset of the biggest value of series data.
// If the maximum is achieved in multiple locations, the first row position is returned.
func Argmax(data Data) int {
	var (
		max   DType = -maxFloat
		pos   int   = -1
		items       = data.Values()
	)
	for i, v := range items {
		if IsNA(v) {
			continue
		}
		if v > max {
			max = v
			pos = i
		}
	}
	return pos
}

// Variance returns variance of values.
// Ddof - Delta Degrees of Freedom. The divisor used in calculations is N - ddof,
// where N represents the number of elements.
func Variance(data Data, mean DType, ddof int) DType {
	var (
		dev   DType
		count int

		values = data.Values()
	)
	for _, v := range values {
		if IsNA(v) {
			continue
		}
		d := v - mean
		dev += (d * d)
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return dev / DType(len(values)-ddof)
}

// Std returns standard deviation.
// Ddof - Delta Degrees of Freedom. The divisor used in calculations is N - ddof,
// where N represents the number of elements.
func Std(data Data, mean DType, ddof int) DType {
	return math.Sqrt(Variance(data, mean, ddof))
}

func First(data Data) DType {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[0]
}

func Last(data Data) DType {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[len(items)-1]
}

func Skew(data Data) DType {
	count := countNotNA(data)

	mean := Sum(data) / count

	m2 := sumAdjustedPow2(data, mean)
	m3 := sumAdjustedPow3(data, mean)

	// fix floating point error.
	m2 = fpZero(m2, Eps)
	m3 = fpZero(m3, Eps)

	if m2 == 0 || m3 == 0 {
		return 0
	}

	if count < 3 {
		return math.NaN()
	}

	g1 := m3 / (math.Sqrt(m2) * m2)

	G1 := ((count * math.Sqrt(count-1)) * g1) / (count - 2)

	return G1
}

func countNotNA(data Data) DType {
	count := 0
	items := data.values
	for _, v := range items {
		if !IsNA(v) {
			count++
		}
	}
	return DType(count)
}

func sumAdjustedPow2(data Data, mean DType) DType {
	var (
		sum   DType
		count int
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		v -= mean
		sum += v * v
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return sum
}

func sumAdjustedPow3(data Data, mean DType) DType {
	var (
		sum   DType
		count int
		items = data.Values()
	)
	for _, v := range items {
		if IsNA(v) {
			continue
		}
		v -= mean
		sum += v * v * v
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return sum
}
