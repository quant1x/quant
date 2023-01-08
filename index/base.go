package index

import (
	"github.com/mymmsc/gox/errors"
	"github.com/quant1x/quant/models/Cache"
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
	ErrCode  = errors.New("股票代码不存在")
	ErrData  = errors.New("数据错误")
	ErrArray = errors.New("数组未对齐")
)

// 指标计算接口
type IndexHandler = func(n int, x int) error

// 公式接口
type FormulaHandler = func(hds []Cache.DayKLine, n int, x int) error

type Formula interface {
	Len() int
	Data() interface{}
	Load(code string) error
}

// 引用n周期前的flag整型值
func RefInt(slice interface{}, flag string, n int) int64 {
	return int64(Ref(slice, flag, n))
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

// Cross 金叉判断
// 当期 a > b 并且 前一期 a < b
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
	fa1 := Ref(a.Data, a.Flag, a.Cycle+1)
	fb1 := Ref(b.Data, b.Flag, b.Cycle+1)
	if fa1 >= fb1 {
		return false
	} else {
		fa0 := Ref(a.Data, a.Flag, a.Cycle+0)
		fb0 := Ref(b.Data, b.Flag, b.Cycle+0)

		if fa0 <= fb0 {
			return false
		} else {
			return true
		}
	}
}
