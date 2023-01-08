package main

import (
	"github.com/mymmsc/gox/api"
	"reflect"
)

const (
	// MaximumResultDays 结果最大天数
	MaximumResultDays int = 3
)

var (
	mapTag map[reflect.Type]map[int]string = nil
)

func init() {
	//fmt.Println("reflect_array init...")
	mapTag = make(map[reflect.Type]map[int]string)
	//fmt.Println("reflect_array init...OK")
}

func initTag(t reflect.Type, tagName string) map[int]string {
	ma, mok := mapTag[t]
	if mok {
		return ma
	}
	ma = nil
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		tag := field.Tag
		if len(tag) > 0 {
			tv, ok := tag.Lookup(tagName)
			if ok {
				if ma == nil {
					ma = make(map[int]string)
					mapTag[t] = ma
				}
				ma[i] = tv
			}
		}
	}
	return ma
}

/*
func Convert(data []string, v interface{}) error {
	val := reflect.ValueOf(v)
	//t := reflect.TypeOf(v)
	//fieldNum := val.NumField()
	//_ = fieldNum
	obj := reflect.ValueOf(v)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := initTag(t, "name")
	if ma == nil {
		return errors.New("can not Convert")
	}
	dl := len(data)
	for i := 0; i < dl; i++ {
		field, ok := ma[i]
		if ok {
			dv := data[i]
			ov := obj.FieldByName(field.Name)
			if ov.CanSet() {
				var value interface{}
				switch ov.Interface().(type) {
				case string:
					value = string(dv)
				case int8:
					t := api.ParseInt(dv)
					value = int8(t)
				case int16:
					t := api.ParseInt(dv)
					value = int16(t)
				case int32:
					t := api.ParseInt(dv)
					value = int32(t)
				case int64:
					t := api.ParseInt(dv)
					value = int64(t)
				case uint8:
					t := api.ParseUint(dv)
					value = uint8(t)
				case uint16:
					t := api.ParseUint(dv)
					value = uint16(t)
				case uint32:
					t := api.ParseUint(dv)
					value = uint32(t)
				case uint64:
					t := api.ParseUint(dv)
					value = t
				case float32:
					t := api.ParseFloat(dv)
					value = float32(t)
				case float64:
					t := api.ParseFloat(dv)
					value = t
				case bool:
					t, _ := strconv.ParseBool(dv)
					value = t
				default:
					value = dv
				}
				ov.Set(reflect.ValueOf(value))
			}
		}
	}
	return nil
}
*/
// 策略结果
type ResultInfo struct {
	Code         string  `name:"证券代码" json:"code" csv:"code" array:"0"`
	Name         string  `name:"证券名称" json:"name" csv:"name" array:"1"`
	Date         string  `name:"信号日期" json:"date" csv:"date" array:"2"`
	Buy          float64 `name:"委托价格" json:"buy" csv:"buy" array:"3"`
	Sell         float64 `name:"目标价格" json:"sell" csv:"sell" array:"4"`
	StrategyCode int     `name:"策略编码" json:"strategy_code" csv:"strategy_code" array:"5"`
	StrategyName string  `name:"策略名称" json:"strategy_name" csv:"strategy_name" array:"6"`
}

func (this *ResultInfo) Headers() []string {
	val := reflect.ValueOf(this)
	//t := reflect.TypeOf(v)
	//fieldNum := val.NumField()
	//_ = fieldNum
	obj := reflect.ValueOf(this)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := initTag(t, "name")
	var sRet []string
	if ma == nil {
		return sRet
	}
	dl := len(ma)
	for i := 0; i < dl; i++ {
		field, ok := ma[i]
		if ok {
			sRet = append(sRet, field)
		}
	}
	return sRet
}

func (this *ResultInfo) Values() []string {
	val := reflect.ValueOf(this)
	//t := reflect.TypeOf(v)
	//fieldNum := val.NumField()
	//_ = fieldNum
	obj := reflect.ValueOf(this)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := initTag(t, "name")
	var sRet []string
	if ma == nil {
		return sRet
	}
	dl := len(ma)
	for i := 0; i < dl; i++ {
		_, ok := ma[i]
		if ok {
			ov := obj.Field(i).Interface()
			value := api.ToString(ov)
			sRet = append(sRet, value)
		}
	}
	return sRet
}
