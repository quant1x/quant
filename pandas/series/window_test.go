package series

import (
	"testing"

	"github.com/quant1x/quant/pandas/series/math"
)

var NaN = math.NaN()

func TestWindow_Sum(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 6, 9, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Sum(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Mean(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Mean(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Min(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 3, 2, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Min(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Max(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 5, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Max(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Std(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{5, 5, 6, 7, 5, 5, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{NaN, NaN, 0.57735026, 1, 1, 1.1547005, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}

			ma := tt.fields.data.Rolling(3).Mean()

			if got := w.Std(ma, 1); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Median(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{0, 1, 2, 3, 4}, []DType{0, 1, 2, 3, 4}),
			},
			MakeData(1, []int64{0, 1, 2, 3, 4}, []DType{NaN, NaN, 1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Median(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Window.Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Variance(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{5, 5, 6, 7, 5, 5, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{NaN, NaN, 3.333333e-01, 1.000000e+00, 1.000000e+00, 1.333333e+00, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}

			ma := tt.fields.data.Rolling(3).Mean()

			if got := w.Variance(ma, 1); !got.Equals(tt.want, 1e-4) {
				t.Errorf("Window.Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Skew(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"period = 4",
			fields{
				len:  4,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 1, 2, 2, 3, 3, 7}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{NaN, NaN, NaN, 0, 0, 0, 1.7198680301759643}),
		},
		{
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 1, 2, 2, 3, 3, 7}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{NaN, NaN, 1.7320508075688787, -1.7320508075688785, 1.732050807568875, -1.732050807568875, 1.7320508075688787}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Skew(tt.fields.data); !got.Equals(tt.want, 1e-4) {
				t.Errorf("Window.Skew() = %v, want %v", got, tt.want)
			}
		})
	}
}
