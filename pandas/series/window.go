package series

import (
	"sort"

	"github.com/quant1x/quant/pandas/series/math"
)

type Window struct {
	len  int
	data Data
}

func (w Window) Sum() Data {
	return w.Apply(Sum)
}

func (w Window) Mean() Data {
	return w.Apply(Mean)
}

func (w Window) Min() Data {
	return w.Apply(Min)
}

func (w Window) Max() Data {
	return w.Apply(Max)
}

func (w Window) Skew(ma Data) Data {
	return w.Apply(Skew)
}

func (w Window) Median() Data {
	return w.applyMedian()
}

func (w Window) Variance(ma Data, ddof int) Data {
	return w.applyVar(Variance, ma, ddof)
}

func (w Window) Std(ma Data, ddof int) Data {
	return w.applyVar(Std, ma, ddof)
}

func (w Window) Apply(agg AggregateFunc) Data {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		values[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		slice := w.data.Slice(l, r)
		values[r-1] = agg(slice)
	})

	return clone
}

func (w Window) applyVar(varfn func(data Data, mean DType, ddof int) DType, ma Data, ddof int) Data {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		period = w.len
	)

	total := period - 1

	for i := total; i < len(values); i++ {
		p := i + 1
		v := w.data.Slice(p-period, p)
		values[i] = varfn(v, ma.values[p-1], ddof)
	}

	for i := 0; i < total; i++ {
		values[i] = math.NaN()
	}

	return clone
}

func (w Window) applyMedian() Data {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		tmp    = make([]DType, 0, w.len)
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		values[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		slice := w.data.Slice(l, r)

		tmp = append(tmp[:0], slice.values...)
		sort.Sort(DTypeSlice(tmp))

		values[r-1] = Median(Data{values: tmp})
	})

	return clone
}
