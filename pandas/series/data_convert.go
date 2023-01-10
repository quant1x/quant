//go:build !series_f32

package series

import "github.com/quant1x/quant/pandas/series/vek"

// IndexAsInt32 returns copy of underlying index slice converted to int32 array.
func (d Data) IndexAsInt32() (index []int32) {
	index = make([]int32, len(d.index))
	for i, v := range d.index {
		index[i] = int32(v)
	}
	return index
}

// IndexAsFloat32 returns copy of underlying index slice converted to float32 array.
func (d Data) IndexAsFloat32() (index []float32) {
	index = make([]float32, len(d.index))
	for i, v := range d.index {
		index[i] = float32(v)
	}
	return index
}

// IndexAsFloat64 returns copy of underlying index slice converted to float64 array.
func (d Data) IndexAsFloat64() (index []float64) {
	index = make([]float64, len(d.index))
	for i, v := range d.index {
		index[i] = float64(v)
	}
	return index
}

// ValuesAsInt32 returns copy of underlying values slice converted to int32 array.
func (d Data) ValuesAsInt32() (values []int32) {
	values = make([]int32, len(d.values))

	switch {
	case EnabledAVX2:
		vek.ToInt32(values, d.values)
		return values

	default:
		values = make([]int32, len(d.values))
		for i, v := range d.values {
			values[i] = int32(v)
		}
		return values
	}
}

// ValuesAsInt64 returns copy of underlying values slice converted to float32 array.
func (d Data) ValuesAsInt64() (values []int64) {
	values = make([]int64, len(d.values))

	switch {
	case EnabledAVX2:
		vek.ToInt64(values, d.values)
		return values

	default:
		for i, v := range d.values {
			values[i] = int64(v)
		}
		return values
	}
}

// ValuesAsFloat32 returns copy of underlying values slice converted to float32 array.
func (d Data) ValuesAsFloat32() (values []float32) {
	values = make([]float32, len(d.values))

	switch {
	case EnabledAVX2:
		vek.ToFloat32(values, d.values)
		return values

	default:
		for i, v := range d.values {
			values[i] = float32(v)
		}
		return values
	}
}

// DataAsFloat64 returns copy of underlying values slice converted to float64 array.
func (d Data) ValuesAsFloat64() (values []float64) {
	return d.values
}
