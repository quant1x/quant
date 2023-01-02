package lambda

import (
	"fmt"
	"reflect"
)

type _int struct {
	v interface{}
	t reflect.Type
}

func (i _int) Value() interface{} {
	return i.v
}

func (i *_int) SetZero() {
	switch i.t.Kind() {
	case reflect.Int:
		i.v = 0
	case reflect.Uint8:
		i.v = uint8(0)
	case reflect.Uint16:
		i.v = uint16(0)
	case reflect.Uint32:
		i.v = uint32(0)
	case reflect.Uint64:
		i.v = uint64(0)
	case reflect.Int8:
		i.v = int8(0)
	case reflect.Int16:
		i.v = int16(0)
	case reflect.Int32:
		i.v = int32(0)
	case reflect.Int64:
		i.v = int64(0)
	}
}

type _float struct {
	v interface{}
	t reflect.Type
}

func (f _float) Value() interface{} {
	return f.v
}

func (f *_float) SetZero() {
	switch f.t.Kind() {
	case reflect.Float32:
		f.v = float32(0)
	case reflect.Float64:
		f.v = float64(0)
	}
}

func Adder(t reflect.Type) Add {

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
	}
	vk := t.Kind()
	if contain(support, vk) {
		var add Add
		if vk == reflect.Float64 || vk == reflect.Float32 {
			add = &_float{t: t}
		} else {
			add = &_int{t: t}
		}
		add.SetZero()
		return add
	}
	panic("not support type " + t.String())
}

type Add interface {
	Add(value reflect.Value)
	Value() interface{}
	SetZero()
}

func (i *_int) Add(value reflect.Value) {
	if value.Kind() != i.t.Kind() {
		panic(fmt.Sprintf("not support %s add %s", i.t.String(), value.Type().String()))
	}
	switch value.Type().Kind() {
	case reflect.Int:
		sum := i.v.(int)
		sum += value.Interface().(int)
		i.v = sum
	case reflect.Uint8:
		sum := i.v.(uint8)
		sum += value.Interface().(uint8)
		i.v = sum
	case reflect.Uint16:
		sum := i.v.(uint16)
		sum += value.Interface().(uint16)
		i.v = sum
	case reflect.Uint32:
		sum := i.v.(uint32)
		sum += value.Interface().(uint32)
		i.v = sum
	case reflect.Uint64:
		sum := i.v.(uint64)
		sum += value.Interface().(uint64)
		i.v = sum
	case reflect.Int8:
		sum := i.v.(int8)
		sum += value.Interface().(int8)
		i.v = sum
	case reflect.Int16:
		sum := i.v.(int16)
		sum += value.Interface().(int16)
		i.v = sum
	case reflect.Int32:
		sum := i.v.(int32)
		sum += value.Interface().(int32)
		i.v = sum
	case reflect.Int64:
		sum := i.v.(int64)
		sum += value.Interface().(int64)
		i.v = sum
	}
}

func (f *_float) Add(value reflect.Value) {
	if value.Kind() != f.t.Kind() {
		panic(fmt.Sprintf("not support %s add %s", f.t.String(), value.Type().String()))
	}

	switch value.Type().Kind() {
	case reflect.Float32:
		sum := f.v.(float32)
		sum += value.Interface().(float32)
		f.v = sum
	case reflect.Float64:
		sum := f.v.(float64)
		sum += value.Interface().(float64)
		f.v = sum
	}
}
