package formula

import "testing"

func TestREF(t *testing.T) {
	type args struct {
		slice []float64
		n     int
	}
	var tests = []struct {
		name string
		args args
		want float64
	}{
		{
			name: "t01",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     0,
			},
			want: 4.00,
		},
		{
			name: "t02",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     1,
			},
			want: 3.00,
		},
		{
			name: "t03",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     3,
			},
			want: 1.00,
		},
		{
			name: "t04",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     4,
			},
			want: 0.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := REF(tt.args.slice, tt.args.n); got != tt.want {
				t.Errorf("REF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSMA(t *testing.T) {
	type args struct {
		slice []float64
		n     int
		m     int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "t01",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     0,
				m:     1,
			},
			want: 0.00,
		},
		{
			name: "t02",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     1,
				m:     1,
			},
			want: 0.00,
		},
		{
			name: "t03",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     3,
				m:     1,
			},
			want: 2.4444444444444446,
		},
		{
			name: "t04",
			args: args{
				slice: []float64{1.00, 2.00, 3.00, 4.00},
				n:     4,
				m:     1,
			},
			want: 2.125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SMA(tt.args.slice, tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("SMA() = %v, want %v", got, tt.want)
			}
		})
	}
}
