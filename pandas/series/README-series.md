# Series

![test](https://github.com/WinPooh32/series/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/WinPooh32/series)](https://goreportcard.com/report/github.com/WinPooh32/series)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/series.svg)](https://pkg.go.dev/github.com/WinPooh32/series)

Liteweight library of series data processing functions written in pure ideomatic Go (golang). Inspired by python library [Pandas](https://github.com/pandas-dev/pandas).

Mostly designed for ordered **time** series data.

[Optionally accelerated](https://github.com/WinPooh32/series/issues/2) by AVX2 instructions.

## Functions

- Column/Scalar math operators:
  - Add +
  - Sub -
  - Mul *
  - Div /
  - Mod %

- Column math functions:
  - Cos, Sin, Tan
  - Pow, Log, Log2, Exp
  - Sign
  - Min
  - Max
  - ...
  - Apply custom function

- Rolling aggregation by functions:
  - Sum
  - Mean
  - Median (only for sorted values)
  - Min
  - Max
  - Skew
  - Variance
  - Std (standard deviation)
  - Apply custom function

- Exponential rolling aggregation:
  - (not) adjusted Mean

- Resampling:
  - Upsampling empty values filling:
    - Interpolate by linear method;
    - Pad known values;
    - Keep empty values.
  - Downsampling aggregation functions:
    - Sum
    - Mean
    - Max
    - Min
    - First
    - Last
    - Apply custom function
- Series manipulations:
  - Slice, Clone
  - Sort, Reverse, Compare (indices or values)
  - Diff, Shift
  - Fill N/A values: interpolate, pad existing values or replace by constant.

## Drawing plots

`series.Data` implements `gonum/plot/plotter.XYer` interface. You can check these plotters:

- [gonum/plot](https://github.com/gonum/plot)
- [go-hep/hep](https://github.com/go-hep/hep/tree/main/hplot)
- [pplcc/plotext](https://github.com/pplcc/plotext)

## Data types

For enabling **float32** use build tag `series_f32`, otherwise **float64** will be used as data type.

Build command example:

``` shell
go build -tags series_f32
```

## Examples

[financial technical indicators](https://github.com/WinPooh32/fta/blob/master/fta.go)

[drawing plots](https://github.com/WinPooh32/fta/blob/master/examples/render-plots/main.go)
