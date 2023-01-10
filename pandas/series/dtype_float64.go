//go:build !series_f32

package series

import (
	"github.com/quant1x/quant/pandas/series/math"
)

type DType = float64

const (
	EpsFp32 = 1e-7
	EpsFp64 = 1e-14
	Eps     = EpsFp64

	EnabledFloat32 = false
)

const maxFloat = math.MaxFloat64
