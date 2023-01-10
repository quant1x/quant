//go:build series_f32 && series_avx2

package vek

import "github.com/viterin/vek/vek32"

func AddScalar(y1 []float32, v float32) { vek32.AddNumber_Inplace(y1, v) }
func SubScalar(y1 []float32, v float32) { vek32.SubNumber_Inplace(y1, v) }
func MulScalar(y1 []float32, v float32) { vek32.MulNumber_Inplace(y1, v) }
func DivScalar(y1 []float32, v float32) { vek32.DivNumber_Inplace(y1, v) }

func Add(y1 []float32, y2 []float32)     { vek32.Add_Inplace(y1, y2) }
func Sub(y1 []float32, y2 []float32)     { vek32.Sub_Inplace(y1, y2) }
func Mul(y1 []float32, y2 []float32)     { vek32.Mul_Inplace(y1, y2) }
func Div(y1 []float32, y2 []float32)     { vek32.Div_Inplace(y1, y2) }
func Minimum(y1 []float32, y2 []float32) { vek32.Minimum_Inplace(y1, y2) }
func Maximum(y1 []float32, y2 []float32) { vek32.Maximum_Inplace(y1, y2) }
func Pow(y1 []float32, y2 []float32)     { vek32.Pow_Inplace(y1, y2) }

func Sqrt(y1 []float32) { vek32.Sqrt_Inplace(y1) }
func Abs(y1 []float32)  { vek32.Abs_Inplace(y1) }

func Round(y1 []float32) { vek32.Round_Inplace(y1) }
func Ceil(y1 []float32)  { vek32.Ceil_Inplace(y1) }
func Floor(y1 []float32) { vek32.Floor_Inplace(y1) }

func Min(y1 []float32) float32 { return vek32.Min(y1) }
func Max(y1 []float32) float32 { return vek32.Max(y1) }

func Dot(y1 []float32, y2 []float32) float32 { return vek32.Dot(y1, y2) }

func ArgMin(y1 []float32) int { return vek32.ArgMin(y1) }
func ArgMax(y1 []float32) int { return vek32.ArgMax(y1) }

func Repeat(dst []float32, v float32)   { vek32.Repeat_Into(dst, v, len(dst)) }
func ToInt64(dst []int64, y1 []float32) { vek32.ToInt64_Into(dst, y1) }
func ToInt32(dst []int32, y1 []float32) { vek32.ToInt32_Into(dst, y1) }

func ToFloat64(dst []float64, y1 []float32) { vek32.ToFloat64_Into(dst, y1) }
func ToFloat32(dst []float32, y1 []float32) { panic("not implemented!") }

// float32 exclusive.
func Exp(y1 []float32)   { vek32.Exp_Inplace(y1) }
func Cos(y1 []float32)   { vek32.Cos_Inplace(y1) }
func Sin(y1 []float32)   { vek32.Sin_Inplace(y1) }
func Log(y1 []float32)   { vek32.Log_Inplace(y1) }
func Log2(y1 []float32)  { vek32.Log2_Inplace(y1) }
func Log10(y1 []float32) { vek32.Log10_Inplace(y1) }
