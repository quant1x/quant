package index

import (
	"github.com/quant1x/quant/stock"
	"reflect"
)

// 单一对象反射获取数值
func get(obj interface{}, flag string) float64 {
	t := reflect.ValueOf(obj)
	if t.Kind() != reflect.Struct {
		return stock.DefaultValue
	}
	v := t.FieldByName(flag)
	k := v.Kind()
	switch k {
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return v.Float()
	default:
		return float64(v.Int())
	}
}

// quote
func slice_quote(slice interface{}, n int) (v reflect.Value, count int, err error) {
	v = reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		err = ErrArray
		return
	}
	if n < 1 {
		err = ErrArray
		return
	}
	count = v.Len()
	if count < n {
		err = ErrArray
		return
	}
	return
}

// 引用n周期前的flag浮点值
func Ref(slice interface{}, flag string, n int) float64 {
	v, count, err := slice_quote(slice, n)
	if err != nil {
		return stock.DefaultValue
	}
	hd := v.Index(count - 1 - n).Interface()

	return get(hd, flag)
}

// 引用n周期前的flag浮点值
func _ref(v reflect.Value, flag string, count int, n int) float64 {
	hd := v.Index(count - 1 - n).Interface()

	return get(hd, flag)
}

// Deprecated: 旧版本代码不利于复用
func Ref_v1(slice interface{}, flag string, n int) float64 {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return stock.DefaultValue
	}
	count := val.Len()
	if count < n+1 {
		return stock.DefaultValue
	}
	n = count - n - 1
	hd := val.Index(n).Interface()

	return get(hd, flag)
}

// 指标计算接口
type algorithmHandler = func(a, b float64) float64

// universal
// slice_universal
// iterator
// 切片通用遍历方法
func slice_universal(slice interface{}, flag string, n int, iterator algorithmHandler) float64 {
	v, count, err := slice_quote(slice, n)
	if err != nil {
		return stock.DefaultValue
	}
	var (
		ret    float64 = 0
		inited bool    = false
	)
	pos := count - n
	for i := 0; i < n; i++ {
		hd := v.Index(pos + i).Interface()
		cur := get(hd, flag)
		if !inited {
			ret = cur
			inited = true
			continue
		}
		ret = iterator(ret, cur)
	}
	return ret
}

// 计算n周期内的flag的总和
//
// Deprecated: 旧版本代码不利于复用
func SUM_v1(slice interface{}, flag string, n int) float64 {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return stock.DefaultValue
	}
	if n < 1 {
		return stock.DefaultValue
	}
	count := v.Len()
	if count < n {
		return stock.DefaultValue
	}
	var (
		val float64 = 0
	)
	for i := 0; i < n; i++ {
		hd := v.Index(count - 1 - i).Interface()
		val += get(hd, flag)
	}
	return val
}

// ExpMA EMA 当前值算法
// previous 前一值
// current 当前值
// EMAtoday = α * Pricetoday + ( 1 - α ) * EMAyesterday
// α = n + 1
func ExpMA(previous, current float64, n int) float64 {
	factor := float64(n) + 1
	return (previous*(factor-EmaWeight) + current*EmaWeight) / factor
}

// SUM 计算n周期内的flag的总和
func SUM(slice interface{}, flag string, n int) float64 {
	return slice_universal(slice, flag, n, func(a, b float64) float64 {
		return a + b
	})
}

// HHV 计算n周期内的flag的最大值
func HHV(slice interface{}, flag string, n int) float64 {
	return slice_universal(slice, flag, n, func(a, b float64) float64 {
		if a < b {
			return b
		}
		return a
	})
}

// LLV 计算n周期内的flag的最小值
func LLV(slice interface{}, flag string, n int) float64 {
	return slice_universal(slice, flag, n, func(a, b float64) float64 {
		if a > b {
			return b
		}
		return a
	})
}

// 指标计算接口
type BarHandler = func(a, b float64) bool

func slice_ssincen(slice interface{}, flag string, n int, iterator BarHandler) int {
	v, count, err := slice_quote(slice, n)
	if err != nil {
		return stock.DefaultValue
	}
	var (
		ret int = -1
		//inited bool = false
	)
	//pos := count - n
	for i := 0; i < n; i++ {
		// 对比成交量至少需要2天的数据
		if i < 1 {
			continue
		}
		// CompVal
		v1 := _ref(v, flag, count, i+1)
		v2 := _ref(v, flag, count, i+0)

		bRet := iterator(v1, v2)
		if bRet {
			ret = i
			break
		}
	}
	return ret
}

// BARSSINCEN N周期内第一次X不为0到现在的周期数,N为常量
func BARSSINCEN(slice interface{}, flag string, n int, iterator BarHandler) int {
	return slice_ssincen(slice, flag, n, iterator)
}
