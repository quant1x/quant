package series

import (
	"fmt"
	"testing"

	"github.com/quant1x/quant/pandas/series/math"
)

func BenchmarkPow3_2(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += math.Pow(DType(i), 1.5)
	}
	fmt.Println(s)
}

func BenchmarkSqrtMulasPow3_2(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += math.Sqrt(DType(i)) * DType(i)
	}
	fmt.Println(s)
}

func BenchmarkSumRef(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += 1.0
	}
	fmt.Println(s)
}

func TestMean(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMed(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "len = 0",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "len = 1",
			args: args{
				data: MakeData(1, []int64{1}, []DType{1}),
			},
			want: 1,
		},
		{
			name: "even len = 4",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{1, 2, 3, 4}),
			},
			want: 2.5,
		},
		{
			name: "odd len = 5",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			want: 3,
		},
		{
			name: "even len = 4, with negative values",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{-1, -3, 3, 4}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.data); !fpEq(got, tt.want, EpsFp32) && !(IsNA(got) || IsNA(tt.want)) {
				t.Errorf("Med() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgmin(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Argmin(tt.args.data); got != tt.want {
				t.Errorf("Argmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgmax(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Argmax(tt.args.data); got != tt.want {
				t.Errorf("Argmax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStd(t *testing.T) {
	type args struct {
		data Data
		mean DType
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Std(tt.args.data, tt.args.mean, 1); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkew(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "simple_with_nan",
			args: args{
				data: MakeValues([]DType{1, NaN, 1, 2}),
			},
			want: 1.7320508075688787,
		},
		{
			name: "simple_with",
			args: args{
				data: MakeValues([]DType{1, 1, 2}),
			},
			want: 1.7320508075688787,
		},
		{
			name: "simple_with3",
			args: args{
				data: MakeValues([]DType{1, 1, 1, 1, 2}),
			},
			want: 2.2360679774997902,
		},
		{
			name: "simple_2233",
			args: args{
				data: MakeValues([]DType{2, 2, 3, 3}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Skew(tt.args.data)
			if !fpEq(got, tt.want, 1e-4) {
				t.Errorf("Skew() = %v, want %v", got, tt.want)
			}
		})
	}
}
