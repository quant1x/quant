package linear

import (
	"fmt"
	"github.com/quant1x/quant/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearT1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LinearT1()
		})
	}
}

func TestLeastSquares(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		wantA float64
		wantB float64
	}{
		{
			name: "收盘价模拟",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{10.1, 10.2, 10.05, 10.01, 11},
			},
			wantA: 0.16100000000000136,
			wantB: 9.789000000000005,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := LeastSquares(tt.args.x, tt.args.y)
			assert.Equalf(t, tt.wantA, gotA, "LeastSquares(%v, %v)", tt.args.x, tt.args.y)
			assert.Equalf(t, tt.wantB, gotB, "LeastSquares(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}

func TestPredict(t *testing.T) {
	type args struct {
		y         float64
		slope     float64
		intercept float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "预测第6天收盘价",
			args: args{
				y:         6.0,
				slope:     0.16100000000000136,
				intercept: 9.789000000000005,
			},
			want: 10.755000000000013,
		},
		{
			name: "回测第5天收盘价",
			args: args{
				y:         5.0,
				slope:     0.16100000000000136,
				intercept: 9.789000000000005,
			},
			want: 10.594000000000012,
		},
		{
			name: "回测第4天收盘价",
			args: args{
				y:         4.0,
				slope:     0.16100000000000136,
				intercept: 9.789000000000005,
			},
			want: 10.43300000000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Predict(tt.args.y, tt.args.slope, tt.args.intercept), "Predict(%v, %v, %v)", tt.args.y, tt.args.slope, tt.args.intercept)
		})
	}
}

func TestPredictStock(t *testing.T) {
	df := cache.LoadDataFrame("sz002528")
	CLOSE := df.Close
	//CLOSE = []float64{1, 2, 3, 4, 5}
	data_len := len(CLOSE)
	fmt.Printf("raw   data length: %d \n", data_len)
	// 去掉最后1天的数据
	y := CLOSE[:data_len]
	y_length := len(y)
	fmt.Printf("train data length: %d, last data[%d]=%f \n", y_length, (y_length - 1), y[y_length-1])
	x := make([]float64, len(y))
	for i, v := range y {
		x[i] = float64(i)
		_ = v
	}

	k, b := LeastSquares(x, y)
	// 预测最后1天的下一个交易日的数据
	no := y_length
	fmt.Printf("no: %d, predicting...\n", no)
	p := Predict(float64(no), k, b)
	fmt.Printf("no: %d, predicted=%f\n", no, p)
}
