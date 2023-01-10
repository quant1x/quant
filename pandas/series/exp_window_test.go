package series

import (
	"testing"
)

func TestExpWindow_Mean(t *testing.T) {
	type fields struct {
		data     Data
		atype    AlphaType
		param    DType
		adjust   bool
		ignoreNA bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"simple",
			fields{
				data: MakeData(
					1,
					[]int64{1, 2, 3, 4, 5},
					[]DType{0, 1, 2, NaN, 4},
				),
				atype:    AlphaCom,
				param:    0.5,
				adjust:   true,
				ignoreNA: false,
			},
			MakeData(
				1,
				[]int64{1, 2, 3, 4, 5},
				[]DType{0, 0.6923077, 1.575, 1.575, 3.198347},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := ExpWindow{
				data:     tt.fields.data,
				atype:    tt.fields.atype,
				param:    tt.fields.param,
				adjust:   tt.fields.adjust,
				ignoreNA: tt.fields.ignoreNA,
			}

			if got := w.Mean(); !got.Equals(tt.want, 10e-4) {
				t.Errorf("ExpWindow.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}
