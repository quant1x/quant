package series

import (
	"testing"
)

func Test_fpEq(t *testing.T) {
	type args struct {
		v1  DType
		v2  DType
		eps DType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "eq zero vs zero",
			args: args{
				v1:  0,
				v2:  0,
				eps: EpsFp32,
			},
			want: true,
		},
		{
			name: "not eq zero vs 1",
			args: args{
				v1:  0,
				v2:  1,
				eps: EpsFp32,
			},
			want: false,
		},
		{
			name: "not eq 1 vs zero",
			args: args{
				v1:  1,
				v2:  0,
				eps: EpsFp32,
			},
			want: false,
		},
		{
			name: "eq near",
			args: args{
				v1:  1.000000000302,
				v2:  1.000000000501,
				eps: EpsFp32,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fpEq(tt.args.v1, tt.args.v2, tt.args.eps); got != tt.want {
				t.Errorf("fpEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fpZero(t *testing.T) {
	type args struct {
		v DType
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "1",
			args: args{
				v: 1,
			},
			want: 1,
		},
		{
			name: "1e-08",
			args: args{
				v: 1e-08,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fpZero(tt.args.v, EpsFp32); got != tt.want {
				t.Errorf("fpZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
