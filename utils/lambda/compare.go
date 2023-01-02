package lambda

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func BasicComparator(ele interface{}) (Compare, error) {
	if c, ok := ele.(Compare); ok {
		return c, nil
	}

	k := reflect.TypeOf(ele).Kind()

	support := []reflect.Kind{
		reflect.Int,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Int8,
		reflect.Int16,
		reflect.Uint32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
	}
	if contain(support, k) {
		return &BasicCompare{ele}, nil
	}

	return nil, errors.New("unknown type")
}

func contain(kinds []reflect.Kind, target reflect.Kind) bool {
	for _, k := range kinds {
		if k == target {
			return true
		}
	}
	return false
}

type Compare interface {
	CompareTo(a interface{}) int
}

type BasicCompare struct {
	v interface{}
}

func (p *BasicCompare) CompareTo(a interface{}) int {

	vt, at := reflect.TypeOf(p.v), reflect.TypeOf(a)
	if vt.Kind() != at.Kind() {
		panic(fmt.Sprintf("%s is not %s", vt.String(), at.String()))
	}

	switch p.v.(type) {
	case int:
		return p.v.(int) - a.(int)
	case uint8:
		return int(p.v.(uint8) - a.(uint8))
	case uint16:
		return int(p.v.(uint16) - a.(uint16))
	case uint32:
		return int(p.v.(uint32) - a.(uint32))
	case uint64:
		return int(p.v.(uint64) - a.(uint64))
	case int8:
		return int(p.v.(int8) - a.(int8))
	case int16:
		return int(p.v.(int16) - a.(int16))
	case int32:
		return int(p.v.(int32) - a.(int32))
	case int64:
		return int(p.v.(int64) - a.(int64))
	case float32:
		v := p.v.(float32) - a.(float32)
		if v > 0 {
			return int(math.Ceil(float64(v)))
		} else if v == 0 {
			return 0
		} else {
			return int(math.Floor(float64(v)))
		}

	case float64:
		v := p.v.(float64) - a.(float64)
		if v > 0 {
			return int(math.Ceil(v))
		} else if v == 0 {
			return 0
		} else {
			return int(math.Floor(v))
		}
	case string:
		return strings.Compare(p.v.(string), a.(string))
	default:
		panic("unknown type " + at.String())
	}
}
