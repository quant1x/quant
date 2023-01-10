package series

import (
	"testing"
	"time"

	"github.com/quant1x/quant/pandas/series/math"
)

func TestData_Resample_Interpolate(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		freq   int64
		origin ResampleOrigin
		method InterpolationMethod
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"freq=2 to 1 length upsample",
			fields{2, []int64{2, 4, 6, 8}, []DType{2, 4, 6, 8}},
			args{
				freq:   1,
				origin: OriginStart,
				method: InterpolationNone,
			},
			MakeData(1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, NaN, 4, NaN, 6, NaN, 8}),
		},
		{
			"freq=4 to 1 length upsample",
			fields{4, []int64{0, 4, 8, 12, 16}, []DType{0, 4, 8, 12, 16}},
			args{
				freq:   1,
				origin: OriginStart,
				method: InterpolationNone,
			},
			MakeData(
				1,
				[]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				[]DType{0, NaN, NaN, NaN, 4, NaN, NaN, NaN, 8, NaN, NaN, NaN, 12, NaN, NaN, NaN, 16},
			),
		},
		{
			"freq=4 to 1 length upsample lerp",
			fields{4, []int64{0, 4, 8, 12, 16}, []DType{0, 4, 8, 12, 16}},
			args{
				freq:   1,
				origin: OriginStart,
				method: InterpolationLinear,
			},
			MakeData(
				1,
				[]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				[]DType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Resample(tt.args.freq, OriginStart).Interpolate(tt.args.method); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Resample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Resample(t *testing.T) {
	const (
		second = int64(time.Second)
		minute = int64(time.Minute)
	)

	dayStart := time.Date(2022, 5, 7, 0, 0, 0, 0, time.UTC).UnixNano()

	freq := 2 * minute
	point := dayStart + 1*minute
	nearestFrameBegin := freq * (point / freq)

	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		freq   int64
		origin ResampleOrigin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				freq:   2,
				origin: OriginStart,
			},
			MakeData(2, []int64{1, 3, 5}, []DType{3, 7, 11}),
		},
		{
			"odd length",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 2, 3, 4, 5, 6, 7}},
			args{
				freq:   2,
				origin: OriginStart,
			},
			MakeData(2, []int64{1, 3, 5, 7}, []DType{3, 7, 11, 7}),
		},
		{
			"even length minutes freq",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
				},
				[]DType{1, 2, 3, 4, 5, 6},
			},
			args{
				freq:   2 * minute,
				origin: OriginStart,
			},
			MakeData(
				2*minute,
				[]int64{
					1 * minute,
					3 * minute,
					5 * minute,
				},
				[]DType{3, 7, 11},
			),
		},
		{
			"odd length minutes freq",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
					7 * minute,
				},
				[]DType{1, 2, 3, 4, 5, 6, 7},
			},
			args{
				freq:   2 * minute,
				origin: OriginStart,
			},
			MakeData(
				2*minute,
				[]int64{
					1 * minute,
					3 * minute,
					5 * minute,
					7 * minute,
				},
				[]DType{3, 7, 11, 7},
			),
		},
		{
			"even length minutes freq origin epoch",
			fields{
				1 * minute,
				[]int64{
					dayStart + 1*minute + 1*second,
					dayStart + 2*minute,
					dayStart + 3*minute,
					dayStart + 4*minute,
					dayStart + 5*minute,
					dayStart + 6*minute,
				},
				[]DType{1, 2, 3, 4, 5, 6},
			},
			args{
				freq:   2 * minute,
				origin: OriginEpoch,
			},
			MakeData(
				2*minute,
				[]int64{
					nearestFrameBegin,
					nearestFrameBegin + 2*minute,
					nearestFrameBegin + 4*minute,
					nearestFrameBegin + 6*minute,
				},
				[]DType{1, 5, 9, 6},
			),
		},
		{
			"not aligned values",
			fields{
				1 * minute,
				[]int64{
					dayStart + 1*minute + 1*second,
					dayStart + 2*minute,
					dayStart + 3*minute,
					dayStart + 5*minute,
					dayStart + 6*minute,
				},
				[]DType{1, 2, 3, 5, 6},
			},
			args{
				freq:   2 * minute,
				origin: OriginEpoch,
			},
			MakeData(
				2*minute,
				[]int64{
					nearestFrameBegin,
					nearestFrameBegin + 2*minute,
					nearestFrameBegin + 4*minute,
					nearestFrameBegin + 6*minute,
				},
				[]DType{1, 5, 5, 6},
			),
		},
		{
			"even length minutes freq origin start of the day",
			fields{
				1 * minute,
				[]int64{
					dayStart + 1*minute + 1*second,
					dayStart + 2*minute,
					dayStart + 3*minute,
					dayStart + 4*minute,
					dayStart + 5*minute,
					dayStart + 6*minute,
				},
				[]DType{1, 2, 3, 4, 5, 6},
			},
			args{
				freq:   2 * minute,
				origin: OriginStartDay,
			},
			MakeData(
				2*minute,
				[]int64{
					dayStart,
					dayStart + 2*minute,
					dayStart + 4*minute,
					dayStart + 6*minute,
				},
				[]DType{1, 5, 9, 6},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Resample(tt.args.freq, tt.args.origin).Sum(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Resample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_ResampleMedian(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		freq   int64
		origin ResampleOrigin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length reversed",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{6, 5, 4, 3, 2, 1}},
			args{
				freq:   3,
				origin: OriginStart,
			},
			MakeData(3, []int64{1, 4}, []DType{5, 2}),
		},
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				freq:   3,
				origin: OriginStart,
			},
			MakeData(3, []int64{1, 4}, []DType{2, 5}),
		},
		{
			"odd length",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 2, 3, 4, 5, 6, 7}},
			args{
				freq:   3,
				origin: OriginStart,
			},
			MakeData(3, []int64{1, 4, 7}, []DType{2, 5, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Resample(tt.args.freq, tt.args.origin).Median(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Resample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Add(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 4, 6, 8, 10, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Add(tt.args.r); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Sub(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{0, 0, 0, 0, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Sub(tt.args.r); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Mul(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 4, 9, 16, 25, 36}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Mul(tt.args.r); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Div(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 1, 1, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Div(tt.args.r); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_AddScalar(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		s DType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{5, 6, 7, 8, 9, 10}),
		},
		{
			"8 length",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{1, 2, 3, 4, 5, 6, 7, 8}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{5, 6, 7, 8, 9, 10, 11, 12}),
		},
		{
			"16 length",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []DType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []DType{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.AddScalar(tt.args.s); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.AddScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SubScalar(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		s DType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{-3, -2, -1, 0, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.SubScalar(tt.args.s); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.SubScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MulScalar(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		s DType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{4, 8, 12, 16, 20, 24}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.MulScalar(tt.args.s); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.MulScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_DivScalar(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		s DType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 4, 8, 12, 14, 16}},
			args{
				2,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 4, 6, 7, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.DivScalar(tt.args.s); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.DivScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Fillna(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		value   DType
		inplace bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			name: "simple fillna",
			fields: fields{
				freq:   1,
				index:  []int64{1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 5, 2, NaN},
			},
			args: args{
				value: 0,
			},
			want: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{0, 0, 5, 2, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Fillna(tt.args.value); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Fillna() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Pad(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			name: "complex pad",
			fields: fields{
				freq:   1,
				index:  []int64{-2, -1, 0, 1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 0, NaN, NaN, 5, 2, NaN},
			},
			want: MakeData(1, []int64{-2, -1, 0, 1, 2, 3, 4, 5}, []DType{0, 0, 0, 0, 0, 5, 2, 2}),
		},
		{
			name: "all NaN",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, NaN, NaN, NaN, NaN},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{NaN, NaN, NaN, NaN, NaN, NaN}),
		},
		{
			name: "between NaNs",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 1, 9, NaN, NaN},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{1, 1, 1, 9, 9, 9}),
		},
		{
			name: "NaN at mid",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{0, 1, 2, NaN, 4, 5},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{0, 1, 2, 2, 4, 5}),
		},
		{
			name: "NaN at begin",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 2, 3, 4, 5},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{2, 2, 2, 3, 4, 5}),
		},
		{
			name: "NaN at begin",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{0, 1, 2, 3, NaN, NaN},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{0, 1, 2, 3, 3, 3}),
		},
		{
			name: "NaN last",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{0, 1, 2, 3, 4, NaN},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{0, 1, 2, 3, 4, 4}),
		},
		{
			name: "without NaN",
			fields: fields{
				freq:   1,
				index:  []int64{0, 1, 2, 3, 4, 5},
				values: []DType{0, 1, 2, 3, 4, 5},
			},
			want: MakeData(1, []int64{0, 1, 2, 3, 4, 5}, []DType{0, 1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}

			d.Pad()

			if !tt.want.Equals(d, EpsFp32) {
				t.Fatalf("Data.Pad() = %v, want %v", d.values, tt.want.values)
			}
		})
	}
}

func TestData_Sort(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			name: "simple Sort",
			fields: fields{
				freq:   1,
				index:  []int64{1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 5, 2, NaN},
			},
			want: MakeData(1, []int64{1, 2, 5, 4, 3}, []DType{NaN, NaN, NaN, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}

			d.Sort()

			if len(d.values) != len(tt.want.values) {
				t.Fatalf("Data.Sort() = %v, want %v", d.values, tt.want.values)
			}

			for i, v := range tt.want.values {
				if v != d.values[i] && (!IsNA(v) || !IsNA(d.values[i])) {
					t.Fatalf("Data.Sort() = %v, want %v", d.values, tt.want.values)
				}
			}
		})
	}
}

func TestData_SortStable(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			name: "simple SortStable",
			fields: fields{
				freq:   1,
				index:  []int64{1, 2, 3, 4, 5},
				values: []DType{NaN, NaN, 5, 2, NaN},
			},
			want: MakeData(1, []int64{1, 2, 5, 4, 3}, []DType{NaN, NaN, NaN, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}

			d.SortStable()

			if len(d.values) != len(tt.want.values) {
				t.Fatalf("Data.SortStable() = %v, want %v", d.values, tt.want.values)
			}

			for i, v := range tt.want.values {
				if v != d.values[i] && (!IsNA(v) || !IsNA(d.values[i])) {
					t.Fatalf("Data.SortStable() = %v, want %v", d.values, tt.want.values)
				}
			}
		})
	}
}

func TestData_IndexSort(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			name: "simple IndexSort",
			fields: fields{
				freq:   1,
				index:  []int64{4, 1, 3, 2, 5},
				values: []DType{2, NaN, 5, NaN, NaN},
			},
			want: MakeData(
				1,
				[]int64{1, 2, 3, 4, 5},
				[]DType{NaN, NaN, 5, 2, NaN}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}

			d.IndexSort()

			if !d.Equals(tt.want, EpsFp32) {
				t.Fatalf("Data.IndexSort() = %v, want %v", d.values, tt.want.values)
			}
		})
	}
}

func TestData_IndexSortStable(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			name: "simple IndexSortStable",
			fields: fields{
				freq:   1,
				index:  []int64{4, 1, 3, 2, 5},
				values: []DType{2, NaN, 5, NaN, NaN},
			},
			want: MakeData(
				1,
				[]int64{1, 2, 3, 4, 5},
				[]DType{NaN, NaN, 5, 2, NaN}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}

			d.IndexSortStable()

			if !d.Equals(tt.want, EpsFp32) {
				t.Fatalf("Data.IndexSortStable() = %v, want %v", d.values, tt.want.values)
			}
		})
	}
}

func TestData_Reverse(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"len == 1",
			fields{1, []int64{1}, []DType{1}},
			MakeData(1, []int64{1}, []DType{1}),
		},
		{
			"len == 2",
			fields{1, []int64{1, 2}, []DType{1, 2}},
			MakeData(1, []int64{2, 1}, []DType{2, 1}),
		},
		{
			"len == 3",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			MakeData(1, []int64{3, 2, 1}, []DType{3, 2, 1}),
		},
		{
			"even",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			MakeData(1, []int64{6, 5, 4, 3, 2, 1}, []DType{6, 5, 4, 3, 2, 1}),
		},
		{
			"odd",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 2, 3, 4, 5, 6, 7}},
			MakeData(1, []int64{7, 6, 5, 4, 3, 2, 1}, []DType{7, 6, 5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Reverse(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_IndexReverse(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"len == 1",
			fields{1, []int64{1}, []DType{1}},
			MakeData(1, []int64{1}, []DType{1}),
		},
		{
			"len == 2",
			fields{1, []int64{1, 2}, []DType{1, 2}},
			MakeData(1, []int64{2, 1}, []DType{1, 2}),
		},
		{
			"len == 3",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			MakeData(1, []int64{3, 2, 1}, []DType{1, 2, 3}),
		},
		{
			"even",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			MakeData(1, []int64{6, 5, 4, 3, 2, 1}, []DType{1, 2, 3, 4, 5, 6}),
		},
		{
			"odd",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 2, 3, 4, 5, 6, 7}},
			MakeData(1, []int64{7, 6, 5, 4, 3, 2, 1}, []DType{1, 2, 3, 4, 5, 6, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.IndexReverse(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_DataReverse(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"len == 1",
			fields{1, []int64{1}, []DType{1}},
			MakeData(1, []int64{1}, []DType{1}),
		},
		{
			"len == 2",
			fields{1, []int64{1, 2}, []DType{1, 2}},
			MakeData(1, []int64{1, 2}, []DType{2, 1}),
		},
		{
			"len == 3",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			MakeData(1, []int64{1, 2, 3}, []DType{3, 2, 1}),
		},
		{
			"even",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 2, 3, 4, 5, 6}},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{6, 5, 4, 3, 2, 1}),
		},
		{
			"odd",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{1, 2, 3, 4, 5, 6, 7}},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{7, 6, 5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.DataReverse(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Diff(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		period int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"data length is less then period length",
			fields{1, []int64{1, 2}, []DType{1, 1}},
			args{3},
			MakeData(1, []int64{1, 2}, []DType{NaN, NaN}),
		},
		{
			"data length is equal to period length",
			fields{1, []int64{1, 2, 3}, []DType{1, 1, 2}},
			args{3},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
		},
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 1, 2, 3, 5, 8}},
			args{3},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, NaN, NaN, 2, 4, 6}),
		},
		{
			"diff 1",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []DType{1, 1, 2, 3, 5, 8}},
			args{1},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 0, 1, 1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Diff(tt.args.period); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Shift(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		periods int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"right shift",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{1},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, 1, 2}),
		},
		{
			"right shift overflow",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{4},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
		},
		{
			"right shift equal to data length",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{3},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
		},
		{
			"left shift",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{-1},
			MakeData(1, []int64{1, 2, 3}, []DType{2, 3, NaN}),
		},
		{
			"left shift -2",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{-2},
			MakeData(1, []int64{1, 2, 3}, []DType{3, NaN, NaN}),
		},
		{
			"left shift equal to data length",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{-3},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
		},
		{
			"left shift overflow",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{-4},
			MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Shift(tt.args.periods); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Resize(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		newLen int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"len + 0",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{3},
			MakeData(1, []int64{1, 2, 3}, []DType{1, 2, 3}),
		},
		{
			"len + 1",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{4},
			MakeData(1, []int64{1, 2, 3, math.MaxInt64}, []DType{1, 2, 3, NaN}),
		},
		{
			"len - 1",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{2},
			MakeData(1, []int64{1, 2}, []DType{1, 2}),
		},
		{
			"newLen == 0",
			fields{1, []int64{1, 2, 3}, []DType{1, 2, 3}},
			args{0},
			MakeData(1, []int64{}, []DType{}),
		},
		{
			"oldLen == 0, newLen == 3",
			fields{1, []int64{}, []DType{}},
			args{3},
			MakeData(1, []int64{math.MaxInt64, math.MaxInt64, math.MaxInt64}, []DType{NaN, NaN, NaN}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Resize(tt.args.newLen); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Resize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Equal(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	type args struct {
		r   Data
		eps DType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"NaNs vs NaNs",
			fields{1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, NaN, 4, NaN, 6, NaN, 8}},
			args{
				r:   MakeData(1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, NaN, 4, NaN, 6, NaN, 8}),
				eps: EpsFp32,
			},
			true,
		},
		{
			"zeros vs NaNs",
			fields{1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, 0, 4, 0, 6, 0, 8}},
			args{
				r:   MakeData(1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, NaN, 4, NaN, 6, NaN, 8}),
				eps: EpsFp32,
			},
			false,
		},
		{
			"NaNs vs zeros",
			fields{1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, NaN, 4, NaN, 6, NaN, 8}},
			args{
				r:   MakeData(1, []int64{2, 3, 4, 5, 6, 7, 8}, []DType{2, 0, 4, 0, 6, 0, 8}),
				eps: EpsFp32,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Equals(tt.args.r, tt.args.eps); got != tt.want {
				t.Errorf("Data.DataEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Lerp(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"simple 1 gap - 1",
			fields{1, []int64{0, 1, 2, 3}, []DType{0, 1, NaN, 3}},
			MakeData(
				1,
				[]int64{0, 1, 2, 3},
				[]DType{0, 1, 2, 3},
			),
		},
		{
			"simple 1 gap - 2",
			fields{3, []int64{0, 1, 2, 3}, []DType{2, 5, NaN, 11}},
			MakeData(
				3,
				[]int64{0, 1, 2, 3},
				[]DType{2, 5, 8, 11},
			),
		},
		{
			"simple 2 gaps row",
			fields{3, []int64{0, 1, 2, 3, 4}, []DType{2, 5, NaN, NaN, 11}},
			MakeData(
				3,
				[]int64{0, 1, 2, 3, 4},
				[]DType{2, 5, 7, 9, 11},
			),
		},
		{
			"simple 2 gaps at end",
			fields{3, []int64{0, 1, 2, 3}, []DType{2, 5, NaN, NaN}},
			MakeData(
				3,
				[]int64{0, 1, 2, 3},
				[]DType{2, 5, NaN, NaN},
			),
		},
		{
			"simple 2 gaps at begin",
			fields{3, []int64{0, 1, 2, 3}, []DType{NaN, NaN, 2, 5}},
			MakeData(
				3,
				[]int64{0, 1, 2, 3},
				[]DType{NaN, NaN, 2, 5},
			),
		},
		{
			"complex - 1",
			fields{3, []int64{-1, 0, 1, 2, 3, 4, 5}, []DType{NaN, 2, 5, NaN, NaN, 11, NaN}},
			MakeData(
				3,
				[]int64{-1, 0, 1, 2, 3, 4, 5},
				[]DType{NaN, 2, 5, 7, 9, 11, NaN},
			),
		},
		{
			"complex - 2",
			fields{
				3,
				[]int64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]DType{NaN, 2, 5, NaN, NaN, 11, NaN, 16, NaN, NaN, NaN},
			},
			MakeData(
				3,
				[]int64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]DType{NaN, 2, 5, 7, 9, 11, 13.5, 16, NaN, NaN, NaN},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Lerp(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Lerp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Cumsum(t *testing.T) {
	type fields struct {
		freq   int64
		index  []int64
		values []DType
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"one",
			fields{1, []int64{0}, []DType{5}},
			MakeData(
				1,
				[]int64{0},
				[]DType{5},
			),
		},
		{
			"simple",
			fields{1, []int64{0, 1, 2, 3}, []DType{5, 10, 15, 20}},
			MakeData(
				1,
				[]int64{0, 1, 2, 3},
				[]DType{5, 15, 30, 50},
			),
		},
		{
			"simple 1 gap - 1",
			fields{1, []int64{0, 1, 2, 3}, []DType{0, 1, NaN, 3}},
			MakeData(
				1,
				[]int64{0, 1, 2, 3},
				[]DType{0, 1, NaN, 4},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:   tt.fields.freq,
				index:  tt.fields.index,
				values: tt.fields.values,
			}
			if got := d.Cumsum(); !got.Equals(tt.want, EpsFp32) {
				t.Errorf("Data.Lerp() = %v, want %v", got, tt.want)
			}
		})
	}
}
