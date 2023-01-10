package series

import (
	"sort"
)

type DTypeSlice []DType

func (x DTypeSlice) Len() int { return len(x) }
func (x DTypeSlice) Less(i, j int) bool {
	return x[i] < x[j] || (IsNA(x[i]) && !IsNA(x[j]))
}
func (x DTypeSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type sortable Data

func (x sortable) Len() int { return len(x.values) }

func (x sortable) Less(i, j int) bool {
	return x.values[i] < x.values[j] || (IsNA(x.values[i]) && !IsNA(x.values[j]))
}

func (x sortable) Swap(i, j int) {
	x.values[i], x.values[j] = x.values[j], x.values[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

type argSortable Data

func (x argSortable) Len() int { return len(x.index) }

func (x argSortable) Less(i, j int) bool { return x.index[i] < x.index[j] }

func (x argSortable) Swap(i, j int) {
	x.values[i], x.values[j] = x.values[j], x.values[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

// IndexSort sorts data's index.
func (d Data) IndexSort() Data {
	sort.Sort(argSortable(d))
	return d
}

// Sort sorts data.
func (d Data) Sort() Data {
	sort.Sort(sortable(d))
	return d
}

// IndexSortStable sorts data's index using stable sort algorithm.
func (d Data) IndexSortStable() Data {
	sort.Stable(argSortable(d))
	return d
}

// SortStable sorts data's index using stable sort algorithm.
func (d Data) SortStable() Data {
	sort.Stable(sortable(d))
	return d
}
