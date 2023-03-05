package main

import (
	"fmt"
	"gitee.com/quant1x/data/stock"
	"gitee.com/quant1x/pandas/stat"
	"github.com/mymmsc/gox/api"
	"github.com/quant1x/quant/indicator"
	"github.com/quant1x/quant/labs/linear"
	"github.com/quant1x/quant/labs/sample"
	"reflect"
	"sort"
)

const (
	// MaximumResultDays 结果最大天数
	MaximumResultDays int = 3
	// CACHE_STRATEGY_PATH 策略文件存储路径
	CACHE_STRATEGY_PATH = "strategy"
)

var (
	mapTag map[reflect.Type]map[int]string = nil
)

func init() {
	mapTag = make(map[reflect.Type]map[int]string)
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

// ResultInfo 策略结果
type ResultInfo struct {
	Code         string  `name:"证券代码" json:"code" csv:"code" array:"0"`
	Name         string  `name:"证券名称" json:"name" csv:"name" array:"1"`
	Date         string  `name:"信号日期" json:"date" csv:"date" array:"2"`
	Buy          float64 `name:"委托价格" json:"buy" csv:"buy" array:"3"`
	Sell         float64 `name:"目标价格" json:"sell" csv:"sell" array:"4"`
	StrategyCode int     `name:"策略编码" json:"strategy_code" csv:"strategy_code" array:"5"`
	StrategyName string  `name:"策略名称" json:"strategy_name" csv:"strategy_name" array:"6"`
	Tendency     string  `name:"短线趋势" json:"tendency" csv:"tendency" array:"7"`
	Open         float64 `name:"预计开盘" json:"open" csv:"open" array:"8"`
	CLOSE        float64 `name:"预计收盘" json:"close" csv:"close" array:"9"`
	High         float64 `name:"预计最高" json:"high" csv:"high" array:"10"`
	Low          float64 `name:"预计最低" json:"low" csv:"low" array:"11"`
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

// Values 输出表格的行和列
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
			var str string
			switch v := ov.(type) {
			case float32:
				str = fmt.Sprintf("%.02f", v)
			case float64:
				str = fmt.Sprintf("%.02f", v)
			default:
				str = api.ToString(ov)
			}
			sRet = append(sRet, str)
		}
	}
	return sRet
}

// Predict 预测趋势
func (this *ResultInfo) Predict() {
	N := 3
	df := stock.KLine(this.Code)
	limit := stat.RangeFinite(-N)
	OPEN := df.Col("open").Select(limit)
	CLOSE := df.Col("close").Select(limit)
	HIGH := df.Col("high").Select(limit)
	LOW := df.Col("low").Select(limit)
	lastClose := stat.AnyToFloat64(CLOSE.IndexOf(-1))
	po := linear.CurveRegression(OPEN).IndexOf(-1).(stat.DType)
	pc := linear.CurveRegression(CLOSE).IndexOf(-1).(stat.DType)
	ph := linear.CurveRegression(HIGH).IndexOf(-1).(stat.DType)
	pl := linear.CurveRegression(LOW).IndexOf(-1).(stat.DType)
	if po > lastClose {
		this.Tendency = "高开"
	} else if po == lastClose {
		this.Tendency = "平开"
	} else {
		this.Tendency = "低开"
	}
	if pl > ph {
		this.Tendency += ",冲高回落"
	} else if pl > pc {
		this.Tendency += ",探底回升"
	} else if pc < pl {
		this.Tendency += ",趋势向下"
	} else {
		this.Tendency += ",短线向好"
	}

	fs := []float64{float64(po), float64(pc), float64(ph), float64(pl)}
	sort.Float64s(fs)
	this.Open = fs[1]
	this.CLOSE = fs[2]
	this.High = fs[3]
	this.Low = fs[0]

	_ = lastClose
}

// Cross 预测趋势
func (this *ResultInfo) Cross() bool {
	//fmt.Println(this.Code)
	N := 1
	df := stock.KLine(this.Code)
	df = linear.CrossTrend(df)
	if df.Nrow() < 2 {
		return false
	}
	cross := df.Col("cross").IndexOf(-N).(bool)
	cross1 := df.Col("cross").IndexOf(-N - 1).(bool)
	if cross && !cross1 {
		return true
	}
	return false
}

// DetectVolume 检测量能变化
func (this *ResultInfo) DetectVolume() bool {
	N := 10
	df := stock.KLine(this.Code)
	dates := df.Col("date").Select(stat.RangeFinite(-N)).Values().([]string)
	df = stock.Tick(this.Code, dates)
	if df.Nrow() < 2 {
		return false
	}
	bv := df.Col("bv").IndexOf(-1).(float64)
	sv := df.Col("sv").IndexOf(-1).(float64)
	bs := df.Col("bs").IndexOf(-1).(float64)
	if bv > sv && bs < 0 {
		return true
	}
	return false
}

// Sample 处理结果的置信区间
func (this *ResultInfo) Sample() bool {
	//fmt.Println(this.Code)
	N := 89
	df := stock.KLine(this.Code)
	if df.Err != nil {
		return false
	}
	if df.Nrow() < N {
		return false
	}
	ret := indicator.F89K(df, N)
	if ret.Nrow() < 1 {
		return false
	}
	df = sample.ConfidenceInterval(ret)
	ci := df.Col("cib").IndexOf(-1).(bool)
	if ci {
		return true
	}
	return false
}
