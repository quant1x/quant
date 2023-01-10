//go:build series_f32

package series

import (
	"github.com/quant1x/quant/pandas/series/math"
)

type DType = float32

const (
	EpsFp32 = 1e-7
	EpsFp64 = 1e-14
	Eps     = EpsFp32

	EnabledFloat32 = true
)

const maxFloat = math.MaxFloat32
