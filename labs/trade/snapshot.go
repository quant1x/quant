package trade

import (
	"fmt"
	"gitee.com/quant1x/data/util"
	"github.com/mymmsc/gox/api"
	"reflect"
)

type QuoteSnapshot struct {
	Market          uint8   // 市场
	Code            string  `name:"证券代码"`  // 代码
	Name            string  `name:"证券名称"`  // 证券名称
	Active1         uint16  `name:"活跃度"`   // 活跃度
	LastClose       float64 `name:"昨收"`    // 昨收
	Open            float64 `name:"开盘价"`   // 开盘
	OpenZf          float64 `name:"开盘涨幅%"` // 开盘
	Price           float64 `name:"现价"`    // 现价
	ZhangDieFu      float64 `name:"涨跌幅%"`  // 涨跌幅
	OpenBuy         float64 `name:"溢价率%"`  // 集合竞价买入溢价
	High            float64 // 最高
	Low             float64 // 最低
	ServerTime      string  // 时间
	ReversedBytes0  int     // 保留(时间 ServerTime)
	ReversedBytes1  int     // 保留
	Vol             int     // 总量
	CurVol          int     // 现量
	Amount          float64 // 总金额
	SVol            int     // 内盘
	BVol            int     // 外盘
	IndexOpenAmount int     // 指数-集合竞价成交金额=开盘成交金额
	StockOpenAmount int     // 个股-集合竞价成交金额=开盘成交金额
	OpenVolume      int     `name:"开盘量"` // 集合竞价-开盘量, 单位是股
	Bid1            float64
	Ask1            float64
	BidVol1         int
	AskVol1         int
	Bid2            float64
	Ask2            float64
	BidVol2         int
	AskVol2         int
	Bid3            float64
	Ask3            float64
	BidVol3         int
	AskVol3         int
	Bid4            float64
	Ask4            float64
	BidVol4         int
	AskVol4         int
	Bid5            float64
	Ask5            float64
	BidVol5         int
	AskVol5         int
	ReversedBytes4  uint16  // 保留
	ReversedBytes5  int     // 保留
	ReversedBytes6  int     // 保留
	ReversedBytes7  int     // 保留
	ReversedBytes8  int     // 保留
	Rate            float64 // 涨速
	Active2         uint16  // 活跃度
	TopNo           int     // 板块排名
	TopCode         string  // 领涨个股
	TopName         string  // 领涨个股名称
	TopRate         float64 // 领涨个股涨幅
	ZhanTing        int     // 涨停数
	Ling            int     // 平盘数
	Count           int     // 总数
	LiuTongPan      float64 `name:"流通盘"`    // 流通盘
	FreeGuBen       float64 `name:"自由流通股本"` // 自由流通股本
	TurnZ           float64 `name:"开盘换手Z%"` // 开盘换手
}

func (this *QuoteSnapshot) Headers() []string {
	val := reflect.ValueOf(this)
	obj := reflect.ValueOf(this)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := util.InitTag(t, "name")
	var sRet []string
	if ma == nil {
		return sRet
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		field, ok := ma[i]
		if ok {
			sRet = append(sRet, field)
		}
	}
	return sRet
}

// Values 输出表格的行和列
func (this *QuoteSnapshot) Values() []string {
	val := reflect.ValueOf(this)
	obj := reflect.ValueOf(this)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := util.InitTag(t, "name")
	var sRet []string
	if ma == nil {
		return sRet
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
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
