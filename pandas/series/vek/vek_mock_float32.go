//go:build series_f32 && !series_avx2

package vek

func AddScalar(y1 []float32, v float32) { panic("unreachable!") }
func SubScalar(y1 []float32, v float32) { panic("unreachable!") }
func MulScalar(y1 []float32, v float32) { panic("unreachable!") }
func DivScalar(y1 []float32, v float32) { panic("unreachable!") }

func Add(y1 []float32, y2 []float32)     { panic("unreachable!") }
func Sub(y1 []float32, y2 []float32)     { panic("unreachable!") }
func Mul(y1 []float32, y2 []float32)     { panic("unreachable!") }
func Div(y1 []float32, y2 []float32)     { panic("unreachable!") }
func Minimum(y1 []float32, y2 []float32) { panic("unreachable!") }
func Maximum(y1 []float32, y2 []float32) { panic("unreachable!") }
func Pow(y1 []float32, y2 []float32)     { panic("unreachable!") }

func Sqrt(y1 []float32)  { panic("unreachable!") }
func Abs(y1 []float32)   { panic("unreachable!") }
func Round(y1 []float32) { panic("unreachable!") }
func Ceil(y1 []float32)  { panic("unreachable!") }
func Floor(y1 []float32) { panic("unreachable!") }

func Min(y1 []float32) float32 { panic("unreachable!") }
func Max(y1 []float32) float32 { panic("unreachable!") }

func Dot(y1 []float32, y2 []float32) float32 { panic("unreachable!") }

func ArgMin(y1 []float32) int { panic("unreachable!") }
func ArgMax(y1 []float32) int { panic("unreachable!") }

func Repeat(dst []float32, v float32)       { panic("unreachable!") }
func ToInt64(dst []int64, y1 []float32)     { panic("unreachable!") }
func ToInt32(dst []int32, y1 []float32)     { panic("unreachable!") }
func ToFloat64(dst []float64, y1 []float32) { panic("unreachable!") }
func ToFloat32(dst []float32, y1 []float32) { panic("unreachable!") }

func Exp(y1 []float32)   { panic("unreachable!") }
func Cos(y1 []float32)   { panic("unreachable!") }
func Sin(y1 []float32)   { panic("unreachable!") }
func Log(y1 []float32)   { panic("unreachable!") }
func Log2(y1 []float32)  { panic("unreachable!") }
func Log10(y1 []float32) { panic("unreachable!") }
