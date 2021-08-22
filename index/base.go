package index

import (
	"github.com/mymmsc/gox/errors"
	"github.com/quant1x/quant/models/Cache"
	"github.com/quant1x/quant/stock"
	"reflect"
)

const (
	Open   = "Open"
	Close  = "Close"
	High   = "High"
	Low    = "Low"
	Volume = "Volume"
)

var (
	ErrCode = errors.New("股票代码不存在")
	ErrData = errors.New("数据错误")
)

// 指标计算接口
type IndexHandler = func(n int, x int) (error)

// 公式接口
type FormulaHandler = func(hds []Cache.DayKLine, n int, x int) (error)

type Formula interface {
	Len() int
	Load(code string) (error)
	Data() interface{}
}

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

// 引用n周期前的flag浮点值
func Ref(slice interface{}, flag string, n int) float64 {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return stock.DefaultValue
	}
	count := val.Len()
	if count < n + 1 {
		return stock.DefaultValue
	}
	n = count - n - 1
	hd := val.Index(n).Interface()

	return get(hd, flag)
}

// 引用n周期前的flag整型值
func RefInt(slice interface{}, flag string, n int) int64 {
	return int64(Ref(slice, flag, n))
}

// 计算n周期内的flag的总和
func SUM(slice interface{}, flag string, n int) (float64) {
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

/*

// 金叉, a 上穿 b
// a和b对调就是死叉
func (self StockHistory) Cross(a, b string) bool {
	f1 := self.ref(a, 1)
	f2 := self.ref(b, 1)

	if f1 >= f2 {
		return false
	} else {
		f1 = self.ref(a, 0)
		f2 = self.ref(b, 0)

		if f1 <= f2 {
			return false
		} else {
			return true
		}
	}
}
*/

type CompVal struct {
	Data  interface{} // 数据切片
	Flag  string      // 字段名
	Cycle int         // 周期
}

// 金叉判断, 当期 a > b 并且 前一期 a < b
// a和b对调就是死叉
func Cross(a, b CompVal) bool {
	sa := reflect.ValueOf(a.Data)
	sb := reflect.ValueOf(b.Data)
	if sa.Kind() != reflect.Slice || sb.Kind() != reflect.Slice {
		return false
	}

	ca := sa.Len()
	cb := sb.Len()
	if ca < a.Cycle || cb < b.Cycle {
		return false
	}
	// 首先判断前一周期是否 a < b
	fa1 := Ref(a.Data, a.Flag, a.Cycle + 1)
	fb1 := Ref(b.Data, b.Flag, b.Cycle + 1)
	if fa1 >= fb1 {
		return false
	} else {
		fa0 := Ref(a.Data, a.Flag, a.Cycle + 0)
		fb0 := Ref(b.Data, b.Flag, b.Cycle + 0)

		if fa0 <= fb0 {
			return false
		} else {
			return true
		}
	}
}

// EMA 当前值算法
// previous 前一值
// current 当前值
// EMAtoday = α * Pricetoday + ( 1 - α ) * EMAyesterday
// α = n + 1
func ExpMA(previous, current float64, n int) float64 {
	factor := float64(n) + 1
	return (previous * (factor - EmaWeight) + current * EmaWeight) / factor
}
