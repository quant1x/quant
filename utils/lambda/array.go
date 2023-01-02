package lambda

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// make Array from source(TIn[] type)
// source support array or slice type
func LambdaArray(source interface{}) Array {
	t := reflect.TypeOf(source)
	arr := _array{source, t, t.Elem(), reflect.ValueOf(source)}
	if !arr.IsSlice() {
		err := fmt.Errorf("source type is %s, not array ", arr.arrayType.Kind())
		panic(err)
	}
	return &arr
}

type Array interface {

	// return true when obj is slice
	IsSlice() bool

	// array join into string
	// eg
	// JoinOptions.express func(u user) string {return u.name }
	// JoinOptions.Symbol default `,`
	Join(options JoinOptions) string

	// array filter
	// eg: arr.Filter(func(ele int) bool{ return ele>10})
	Filter(express interface{}) Array

	// sort by quick
	// eg
	Sort(express interface{}) Array

	// sort by quick multithreading
	SortMT(express interface{}) Array

	// map to new array
	// express func(el T) T{ return T }
	Map(express interface{}) Array

	// append element
	Append(elements ...interface{}) Array

	// maximum of array
	// express eg: express func(ele TIn) TOut{ return TOut },TOut must be number Type or Compare
	Max(express interface{}) interface{}

	// minimum of array
	// express eg: express func(ele TIn) TOut{ return TOut },TOut must be number Type or Compare
	Min(express interface{}) interface{}

	// Determines whether the Array contains any elements
	Any(express interface{}) bool

	// Determines whether the condition is satisfied for all elements in the Array
	All(express interface{}) bool

	// Returns a number indicating how many elements in the specified Array satisfy the condition
	Count(express interface{}) int

	// Returns the first element of an Array that satisfies the condition
	First(express interface{}) (interface{}, error)

	// Returns the last element of an Array that satisfies the condition
	Last(express interface{}) (interface{}, error)

	// Returns the zero based index of the first occurrence in an Array
	Index(i int) (interface{}, error)

	// skip and Returns the elements
	Take(skip, count int) Array

	// sum of the values returned by the expression
	Sum(express interface{}) interface{}

	// average of the values returned by the expression
	Average(express interface{}) float64

	// Determines whether the array contains the specified element
	// number type use default comparator
	// other type can implements Compare
	Contains(express interface{}) bool

	// array or slice pointer
	// Array.Pointer().([]T or [n]T)
	Pointer() interface{}
}

func innerLambdaArray(value reflect.Value) Array {
	t := value.Type()
	arr := _array{value, t, t.Elem(), value}
	return &arr
}

type _array struct {

	// source array
	source interface{}
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

func (p *_array) Contains(express interface{}) bool {
	sz := p.Len()
	if express == nil {
		panic("express is null")
	}
	if t := reflect.TypeOf(express); t.Kind() == reflect.Func {
		expType := reflect.TypeOf(express)
		checkExpress(expType, []reflect.Type{p.elementType}, []reflect.Type{reflect.TypeOf(true)})
		fn := reflect.ValueOf(express)
		for i := 0; i < sz; i++ {
			ret := fn.Call([]reflect.Value{p.value.Index(i)})
			if ret[0].Interface().(bool) {
				return true
			}
		}
	} else if tor, err := BasicComparator(express); err == nil {
		for i := 0; i < sz; i++ {
			if tor.CompareTo(p.value.Index(i).Interface()) == 0 {
				return true
			}
		}
	} else if eq, ok := express.(Equal); ok {
		for i := 0; i < sz; i++ {
			if eq.Equals(p.value.Index(i).Interface()) {
				return true
			}
		}
	} else {
		panic("unknown type " + t.String())
	}
	return false
}

func (p *_array) Average(express interface{}) float64 {
	length := p.Len()
	if length == 0 {
		return float64(0)
	}
	sum := p.Sum(express)

	switch sum.(type) {
	case int:
		return float64(sum.(int)) / float64(length)
	case uint8:
		return float64(sum.(uint8)) / float64(length)
	case uint16:
		return float64(sum.(uint16)) / float64(length)
	case uint32:
		return float64(sum.(uint32)) / float64(length)
	case uint64:
		return float64(sum.(uint64)) / float64(length)
	case int8:
		return float64(sum.(int8)) / float64(length)
	case int16:
		return float64(sum.(int16)) / float64(length)
	case int32:
		return float64(sum.(int32)) / float64(length)
	case int64:
		return float64(sum.(int64)) / float64(length)
	case float32:
		return float64(sum.(float32)) / float64(length)
	case float64:
		return sum.(float64) / float64(length)
	default:
		panic("unknown type " + reflect.TypeOf(sum).String())
	}
}

func (p *_array) Append(elements ...interface{}) Array {
	ret := LambdaArray(elements).Map(func(ele interface{}) reflect.Value {
		if t := reflect.TypeOf(ele); t.Kind() != p.elementType.Kind() {
			panic(fmt.Sprintf("element type[%s] is not %s.", t.String(), p.elementType.String()))
		}
		return reflect.ValueOf(ele)
	}).Pointer().([]reflect.Value)
	p.value = reflect.Append(p.value, ret...)
	return p
}

func (p *_array) Any(express interface{}) bool {
	if express == nil {
		return p.Len() > 0
	}
	checkExpress(
		reflect.TypeOf(express),
		[]reflect.Type{p.elementType},
		[]reflect.Type{reflect.TypeOf(true)})

	length := p.Len()
	fn := reflect.ValueOf(express)
	for i := 0; i < length; i++ {
		if fn.Call([]reflect.Value{p.value.Index(i)})[0].Interface().(bool) {
			return true
		}
	}
	return false
}

func (p *_array) All(express interface{}) bool {
	if express == nil {
		return p.Len() > 0
	}
	checkExpress(
		reflect.TypeOf(express),
		[]reflect.Type{p.elementType},
		[]reflect.Type{reflect.TypeOf(true)})
	length := p.Len()
	fn := reflect.ValueOf(express)
	for i := 0; i < length; i++ {
		if !fn.Call([]reflect.Value{p.value.Index(i)})[0].Interface().(bool) {
			return false
		}
	}
	if p.Len() > 0 {
		return true
	} else {
		return false
	}
}

func (p *_array) Count(express interface{}) int {
	if express == nil {
		return p.Len()
	}
	checkExpress(
		reflect.TypeOf(express),
		[]reflect.Type{p.elementType},
		[]reflect.Type{reflect.TypeOf(true)})
	fn := reflect.ValueOf(express)
	count := 0
	p.EachV(func(v reflect.Value, _ int) {
		if fn.Call([]reflect.Value{v})[0].Interface().(bool) {
			count++
		}
	})
	return count
}

func (p *_array) Find(express interface{}, start, step int) (interface{}, error) {
	length := p.Len()
	if length == 0 {
		return nil, errors.New("empty array")
	}

	if express == nil {
		return p.value.Index(0).Interface(), nil
	}
	checkExpress(
		reflect.TypeOf(express),
		[]reflect.Type{p.elementType},
		[]reflect.Type{reflect.TypeOf(true)})
	fn := reflect.ValueOf(express)
	for i := start; i < length && i >= 0; i += step {
		if ele := p.value.Index(i); fn.Call([]reflect.Value{ele})[0].Interface().(bool) {
			return ele.Interface(), nil
		}
	}
	return nil, errors.New("not found")
}

func (p *_array) First(express interface{}) (interface{}, error) {
	return p.Find(express, 0, 1)
}

func (p *_array) Last(express interface{}) (interface{}, error) {
	return p.Find(express, p.Len()-1, -1)
}

func (p *_array) Index(i int) (interface{}, error) {
	if i < p.Len() {
		return p.value.Index(i), nil
	}
	return nil, errors.New(fmt.Sprintf("%d out of range", i))
}

func (p *_array) Take(skip, count int) Array {
	length := p.value.Len()

	ret := reflect.MakeSlice(p.arrayType, 0, 0)
	for i := skip; i < length; i++ {
		if count > 0 {
			ret = reflect.Append(ret, p.value.Index(i))
			count--
		}
		if count == 0 {
			break
		}
	}
	return innerLambdaArray(ret)
}

func (p *_array) Sum(express interface{}) interface{} {

	var add Add
	if express == nil {
		add = Adder(p.elementType)
	} else {
		checkExpress(
			reflect.TypeOf(express),
			[]reflect.Type{p.elementType},
			nil)
		add = Adder(reflect.TypeOf(express).Out(0))
	}

	length := p.Len()
	if length == 0 {
		return add.Value()
	}

	fn := reflect.ValueOf(express)

	fv := func(i int) reflect.Value {
		v := p.value.Index(i)
		if express == nil {
			return v
		}
		return fn.Call([]reflect.Value{v})[0]
	}

	for i := 0; i < length; i++ {
		add.Add(fv(i))
	}
	return add.Value()
}

func (p *_array) Pointer() interface{} {
	return p.value.Interface()
}

func (p *_array) IsSlice() bool {
	return p.arrayType.Kind() == reflect.Slice || p.arrayType.Kind() == reflect.Array
}

func (p *_array) Len() int {
	return p.value.Len()
}

// check the function express
// exp the express function type
// in express function parameter types
// out express function return types
func checkExpress(exp reflect.Type, in []reflect.Type, out []reflect.Type) {
	if exp.Kind() != reflect.Func {
		panic("express is not a func express")
	}
	// check in
	numIn := exp.NumIn()
	lenIn := len(in)
	if numIn != lenIn {
		panic(fmt.Errorf("lambda express parameter count must be %d", lenIn))
	}
	for i := 0; i < lenIn; i++ {
		if in[i].Kind() != exp.In(i).Kind() {
			panic(fmt.Errorf("lambda express the %d'th parameter Type must be %s,not %s,func=%s",
				i, exp.In(i).String(), in[i].String(), exp.String()))
		}
	}
	if out == nil {
		return
	}
	// check output
	numOut := exp.NumOut()
	lenOut := len(out)
	if numOut != lenOut {
		panic(fmt.Errorf("lambda express return Types count must be %d", lenOut))
	}
	for i := 0; i < lenOut; i++ {
		if out[i].Kind() != exp.Out(i).Kind() {
			panic(fmt.Errorf("lambda express the %d'th return Type must be %s", i, exp.Out(i).String()))
		}
	}
}

// check the function express
func checkExpressRARTO(express interface{}, in []reflect.Type) reflect.Type {
	t := reflect.TypeOf(express)
	if t.NumOut() == 0 {
		panic("lambda express must has only one return-value.")
	}
	ot := t.Out(0)
	checkExpress(t, in, []reflect.Type{ot})
	return ot
}

func (p *_array) Map(express interface{}) Array {
	in := []reflect.Type{p.elementType}
	ot := checkExpressRARTO(express, in)

	var result reflect.Value
	length := p.Len()
	// slice or array
	isSlice := p.arrayType.Kind() == reflect.Slice
	var element reflect.Value
	if isSlice {
		result = reflect.MakeSlice(reflect.SliceOf(ot), p.Len(), p.Len())
		element = result
	} else {
		result = reflect.New(reflect.ArrayOf(length, ot))
		element = result.Elem()
	}

	funcValue := reflect.ValueOf(express)
	params := []reflect.Value{reflect.ValueOf(0)}
	for i := 0; i < length; i++ {
		params[0] = p.value.Index(i)
		trans := funcValue.Call(params)
		v := element.Index(i)
		v.Set(trans[0])
	}

	return innerLambdaArray(result)
}

type JoinOptions struct {
	Symbol  string
	express interface{}
}

func (p *_array) Join(option JoinOptions) string {
	if option.express != nil {
		return p.Map(option.express).Join(JoinOptions{Symbol: option.Symbol})
	}
	if p.elementType.Kind() != reflect.String {
		panic("the array is not string array")
	}
	if option.Symbol == "" {
		option.Symbol = ","
	}
	length := p.Len()
	var build strings.Builder
	for i := 0; i < length; i++ {
		s := p.value.Index(i).Interface().(string)
		build.WriteString(s)
		if i < length-1 {
			build.WriteString(option.Symbol)
		}
	}
	return build.String()
}

func (p *_array) Filter(express interface{}) Array {
	in := []reflect.Type{p.elementType}
	ft := reflect.TypeOf(express)
	ot := reflect.TypeOf(true)
	checkExpress(ft, in, []reflect.Type{ot})

	ret := reflect.MakeSlice(reflect.SliceOf(p.elementType), 0, 0)
	funcValue := reflect.ValueOf(express)
	params := []reflect.Value{reflect.ValueOf(0)}
	length := p.Len()
	for i := 0; i < length; i++ {
		params[0] = p.value.Index(i)
		trans := funcValue.Call(params)
		if trans[0].Interface().(bool) {
			ret = reflect.Append(ret, params[0])
		}
	}
	return innerLambdaArray(ret)
}

func (p *_array) SortByBubble(express interface{}) Array {
	in := []reflect.Type{p.elementType, p.elementType}
	ft := reflect.TypeOf(express)
	ot := reflect.TypeOf(true)
	checkExpress(ft, in, []reflect.Type{ot})

	length := p.Len()
	v := reflect.ValueOf(0)
	funcValue := reflect.ValueOf(express)
	params := []reflect.Value{v, v}
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			params[0] = p.value.Index(j)
			params[1] = p.value.Index(j + 1)
			trans := funcValue.Call(params)
			if !trans[0].Interface().(bool) {
				temp := params[0].Interface()
				p.value.Index(j).Set(params[1])
				p.value.Index(j + 1).Set(reflect.ValueOf(temp))
			}
		}
	}

	return p
}

func (p *_array) CopyValue() reflect.Value {
	var arr reflect.Value
	if p.IsSlice() {
		arr = reflect.MakeSlice(reflect.SliceOf(p.elementType), p.Len(), p.Len())
	} else {
		arr = reflect.New(reflect.ArrayOf(p.Len(), p.elementType))
	}
	reflect.Copy(arr, p.value)
	return arr
}

func (p *_array) Sort(express interface{}) Array {
	in := []reflect.Type{p.elementType, p.elementType}
	ft := reflect.TypeOf(express)
	ot := reflect.TypeOf(true)
	checkExpress(ft, in, []reflect.Type{ot})

	length := p.Len()
	v := reflect.ValueOf(0)
	funcValue := reflect.ValueOf(express)
	params := []reflect.Value{v, v}

	p.value = p.CopyValue()

	compare := func(a reflect.Value, b int) bool {
		params[0], params[1] = a, p.value.Index(b)
		return funcValue.Call(params)[0].Interface().(bool)
	}

	var inner func(int, int)
	// quick sort
	inner = func(l, r int) {
		if l < r {
			i, j, x := l, r, p.value.Index(l)
			x = reflect.ValueOf(x.Interface())
			for i < j {
				for i < j && compare(x, j) {
					j--
				}
				if i < j {
					p.value.Index(i).Set(p.value.Index(j))
					i++
				}
				for i < j && !compare(x, i) {
					i++
				}
				if i < j {
					p.value.Index(j).Set(p.value.Index(i))
					j--
				}
			}
			p.value.Index(i).Set(x)
			inner(l, i-1)
			inner(i+1, r)
		}
	}
	inner(0, length-1)
	return p
}

func (p *_array) SortMT(express interface{}) Array {
	in := []reflect.Type{p.elementType, p.elementType}
	ft := reflect.TypeOf(express)
	ot := reflect.TypeOf(true)
	checkExpress(ft, in, []reflect.Type{ot})

	funcValue := reflect.ValueOf(express)

	compare := func(a, b reflect.Value) bool {
		return funcValue.Call([]reflect.Value{a, b})[0].Interface().(bool)
	}
	var quick func(arr reflect.Value, ch chan reflect.Value)
	quick = func(arr reflect.Value, ch chan reflect.Value) {
		if arr.Len() == 1 {
			ch <- arr.Index(0)
			close(ch)
			return
		}
		if arr.Len() == 0 {
			close(ch)
			return
		}

		left := reflect.MakeSlice(reflect.SliceOf(p.elementType), 0, 0)
		right := reflect.MakeSlice(reflect.SliceOf(p.elementType), 0, 0)
		length := arr.Len()
		x := arr.Index(0)
		for i := 1; i < length; i++ {
			curr := arr.Index(i)
			if compare(x, curr) {
				left = reflect.Append(left, curr)
			} else {
				right = reflect.Append(right, curr)
			}
		}
		lch := make(chan reflect.Value, left.Len())
		rch := make(chan reflect.Value, right.Len())
		go quick(left, lch)
		go quick(right, rch)
		for v := range lch {
			ch <- v
		}
		ch <- x
		for v := range rch {
			ch <- v
		}
		close(ch)
	}
	ch := make(chan reflect.Value)
	go quick(p.value, ch)
	values := reflect.MakeSlice(reflect.SliceOf(p.elementType), 0, 0)
	for v := range ch {
		values = reflect.Append(values, v)
	}
	return innerLambdaArray(values)
}

func (p *_array) maxOrMin(express interface{}, isMax bool) interface{} {
	if express != nil {
		in := []reflect.Type{p.elementType}
		ft := reflect.TypeOf(express)
		ot := reflect.TypeOf(0)
		checkExpress(ft, in, []reflect.Type{ot})
	}
	var m reflect.Value
	var mc interface{}

	funcValue := reflect.ValueOf(express)
	f := func(value reflect.Value) interface{} {
		if express == nil {
			return value.Interface()
		}
		return funcValue.Call([]reflect.Value{value})[0].Interface()
	}

	p.EachV(func(v reflect.Value, index int) {
		vc := f(v)
		if index == 0 {
			m = v
			mc = vc
		} else {
			tor, err := BasicComparator(vc)
			if err != nil {
				panic(err)
			}
			if isMax {
				if tor.CompareTo(mc) > 0 {
					m = v
					mc = vc
				}
			} else {
				if tor.CompareTo(mc) < 0 {
					m = v
					mc = vc
				}
			}
		}
	})
	return m.Interface()
}

func (p *_array) Max(express interface{}) interface{} {
	return p.maxOrMin(express, true)
}

func (p *_array) Min(express interface{}) interface{} {
	return p.maxOrMin(express, false)
}

func (p *_array) EachV(fn func(v reflect.Value, i int)) {
	if fn == nil {
		return
	}
	length := p.Len()

	for i := 0; i < length; i++ {
		fn(p.value.Index(i), i)
	}
}
