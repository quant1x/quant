package linear

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas/stat"
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
	df := stock.KLine("002528")
	fmt.Println(df)
	length := df.Nrow() - 1
	df1 := df.Subset(length-3, length)
	fmt.Println(df1)
	CLOSE := df1.Col("low").DTypes()
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

	fmt.Println("------------------------------------------------------------")
	p1 := stat.PolyFit(y, x, 2)
	fmt.Println("p1 =", p1)
	fmt.Println("------------------------------------------------------------")

	k, b := LeastSquares(x, y)
	// 预测最后1天的下一个交易日的数据
	no := y_length
	fmt.Printf("no: %d, predicting...\n", no)
	p := Predict(float64(no), k, b)
	fmt.Printf("no: %d, predicted=%f\n", no, p)
}

func TestPolyFit(t *testing.T) {
	x := []float64{0.0, 0.1, 0.2, 0.3, 0.5, 0.8, 1.0}
	y := []float64{1.0, 0.41, 0.50, 0.61, 0.91, 2.02, 2.46}
	A := stat.PolyFit(x, y, 2)
	fmt.Println("A =", A)

	//A2 := []float64{3.131561350718812, -1.2400367769976413, 0.7355767301905694}
	z1 := stat.PolyVal(A, x)
	fmt.Println("z1 =", z1)

	W := 5
	A2 := stat.PolyFit(y, stat.Range[float64](W), 1)
	fmt.Println("A2 =", A2)
	x2 := stat.Repeat[float64](float64(W), W)
	z2 := stat.PolyVal(A2, x2)
	fmt.Println("z2 =", z2)
}
