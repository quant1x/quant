//go:build !series_f32 && series_avx2

package vek

import "github.com/viterin/vek"

func AddScalar(y1 []float64, v float64) { vek.AddNumber_Inplace(y1, v) }
func SubScalar(y1 []float64, v float64) { vek.SubNumber_Inplace(y1, v) }
func MulScalar(y1 []float64, v float64) { vek.MulNumber_Inplace(y1, v) }
func DivScalar(y1 []float64, v float64) { vek.DivNumber_Inplace(y1, v) }

func Add(y1 []float64, y2 []float64)     { vek.Add_Inplace(y1, y2) }
func Sub(y1 []float64, y2 []float64)     { vek.Sub_Inplace(y1, y2) }
func Mul(y1 []float64, y2 []float64)     { vek.Mul_Inplace(y1, y2) }
func Div(y1 []float64, y2 []float64)     { vek.Div_Inplace(y1, y2) }
func Minimum(y1 []float64, y2 []float64) { vek.Minimum_Inplace(y1, y2) }
func Maximum(y1 []float64, y2 []float64) { vek.Maximum_Inplace(y1, y2) }
func Pow(y1 []float64, y2 []float64)     { vek.Pow_Inplace(y1, y2) }

func Sqrt(y1 []float64)  { vek.Sqrt_Inplace(y1) }
func Abs(y1 []float64)   { vek.Abs_Inplace(y1) }
func Round(y1 []float64) { vek.Round_Inplace(y1) }
func Ceil(y1 []float64)  { vek.Ceil_Inplace(y1) }
func Floor(y1 []float64) { vek.Floor_Inplace(y1) }

func Min(y1 []float64) float64 { return vek.Min(y1) }
func Max(y1 []float64) float64 { return vek.Max(y1) }

func Dot(y1 []float64, y2 []float64) float64 { return vek.Dot(y1, y2) }

func ArgMin(y1 []float64) int { return vek.ArgMin(y1) }
func ArgMax(y1 []float64) int { return vek.ArgMax(y1) }

func Repeat(dst []float64, v float64)       { vek.Repeat_Into(dst, v, len(dst)) }
func ToInt64(dst []int64, y1 []float64)     { vek.ToInt64_Into(dst, y1) }
func ToInt32(dst []int32, y1 []float64)     { vek.ToInt32_Into(dst, y1) }
func ToFloat64(dst []float64, y1 []float64) { panic("not implemented!") }
func ToFloat32(dst []float32, y1 []float64) { vek.ToFloat32_Into(dst, y1) }

// float32 exclusive.
func Exp(y1 []float64)   { panic("not implemented!") }
func Cos(y1 []float64)   { panic("not implemented!") }
func Sin(y1 []float64)   { panic("not implemented!") }
func Log(y1 []float64)   { panic("not implemented!") }
func Log2(y1 []float64)  { panic("not implemented!") }
func Log10(y1 []float64) { panic("not implemented!") }
