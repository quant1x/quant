package internal

import (
	"gitee.com/quant1x/data/category"
	"gitee.com/quant1x/data/category/date"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"time"
)

// RealTime 最新数据
type RealTime struct {
	FullCode string `json:"full_code"`
	//0:  "未知",
	UnknownCode string `json:"unknown_code" array:"0"`
	//1:  "名字",
	Name string `json:"name" array:"1"`
	//2:  "代码",
	Code string `json:"code" array:"2"`
	//3:  "当前价格",
	New string `json:"new" array:"3"`
	//4:  "昨收",
	Close string `json:"close" array:"4"`
	//5:  "今开",
	Open string `json:"open" array:"5"`
	//6:  "成交量（手)",
	Volume string `json:"volume" array:"6"`
	//7:  "外盘",
	OuterVol string `json:"outer_vol" array:"7"`
	//8:  "内盘",
	InnerVol string `json:"inner_vol" array:"8"`
	//9:  "买一",
	Buy1Price string `json:"buy1_price" array:"9"`
	//10: "买一量（手）",
	Buy1Vol string `json:"buy1_vol" array:"10"`
	//11: "买二",
	Buy2Price string `json:"buy2_price" array:"11"`
	//12: "买二量（手）",
	Buy2Vol string `json:"buy2_vol" array:"12"`
	//13: "买三",
	Buy3Price string `json:"buy3_price" array:"13"`
	//14: "买三量（手）",
	Buy3Vol string `json:"buy3_vol" array:"14"`
	//15: "买四",
	Buy4Price string `json:"buy4_price" array:"15"`
	//16: "买四量（手）",
	Buy4Vol string `json:"buy4_vol" array:"16"`
	//17: "买五",
	Buy5Price string `json:"buy5_price" array:"17"`
	//18: "买五量（手）",
	Buy5Vol string `json:"buy5_vol" array:"18"`
	//19: "卖一",
	Sell1Price string `json:"sell1_price" array:"19"`
	//20: "卖一量",
	Sell1Vol string `json:"sell1_vol" array:"20"`
	//21: "卖二",
	Sell2Price string `json:"sell2_price" array:"21"`
	//22: "卖二量",
	Sell2Vol string `json:"sell2_vol" array:"22"`
	//23: "卖三",
	Sell3Price string `json:"sell3_price" array:"23"`
	//24: "卖三量",
	Sell3Vol string `json:"sell3_vol" array:"24"`
	//25: "卖四",
	Sell4Price string `json:"sell4_price" array:"25"`
	//26: "卖四量",
	Sell4Vol string `json:"sell4_vol" array:"26"`
	//27: "卖五",
	Sell5Price string `json:"sell5_price" array:"27"`
	//28: "卖五量",
	Sell5Vol string `json:"sell5_vol" array:"28"`
	//29: "最近逐笔成交",
	Deals string `json:"deals" array:"29"`
	//30: "时间",
	Time string `json:"time" array:"30"`
	//31: "涨跌",
	RiseFall string `json:"rise_fall" array:"31"`
	//32: "涨跌%",
	RiseFallPercent string `json:"rise_fall_percent" array:"32"`
	//33: "最高",
	High string `json:"high" array:"33"`
	//34: "最低",
	Low string `json:"low" array:"34"`
	//35: "价格/成交量（手）/成交额",
	TransactionInformation string `json:"transaction_information" array:"35"`
	//36: "成交量（手）",
	Volume1 string `json:"volume1" array:"36"`
	//37: "成交额（万）",
	Amount string `json:"amount" array:"37"`
	//38: "换手率",
	TurnoverRate string `json:"turnover_rate" array:"38"`
	//39: "市盈率",
	PeRatio string `json:"pe_ratio" array:"39"`
	//40: "未知",
	Unknown string `json:"unknown" array:"40"`
	//41: "最高",
	High1 string `json:"high_1" array:"41"`
	//42: "最低",
	Low1 string `json:"low_1" array:"42"`
	//43: "振幅",
	Amplitude string `json:"amplitude" array:"43"`
	//44: "流通市值",
	FreeFloatMarketValue string `json:"free_float_market_value" array:"44"`
	//45: "总市值",
	TotalMarketValue string `json:"total_market_value" array:"45"`
	//46: "市净率",
	MarketRate string `json:"market_rate" array:"46"`
	//47: "涨停价",
	LimitUp string `json:"limit_up" array:"47"`
	//48: "跌停价",
	LimitDown string `json:"limit_down" array:"48"`
}

var (
	stdApi *quotes.StdApi = nil
)

func prepare() *quotes.StdApi {
	if stdApi == nil {
		api_, err := quotes.NewStdApi()
		if err != nil {
			return nil
		}
		stdApi = api_
	}
	return stdApi
}

type QuoteSnapshot struct {
	Market         uint8   // 市场
	Code           string  // 代码
	Name           string  // 证券名称
	Active1        uint16  // 活跃度
	Price          float64 // 现价
	LastClose      float64 // 昨收
	Open           float64 // 开盘
	High           float64 // 最高
	Low            float64 // 最低
	ServerTime     string  // 时间
	ReversedBytes0 int     // 保留(时间 ServerTime)
	ReversedBytes1 int     // 保留
	Vol            int     // 总量
	CurVol         int     // 现量
	Amount         float64 // 总金额
	SVol           int     // 内盘
	BVol           int     // 外盘
	ReversedBytes2 int     // 保留
	ReversedBytes3 int     // 保留
	//BidLevels      []quotes.Level
	//AskLevels      []quotes.Level
	Bid1           float64
	Ask1           float64
	BidVol1        int
	AskVol1        int
	Bid2           float64
	Ask2           float64
	BidVol2        int
	AskVol2        int
	Bid3           float64
	Ask3           float64
	BidVol3        int
	AskVol3        int
	Bid4           float64
	Ask4           float64
	BidVol4        int
	AskVol4        int
	Bid5           float64
	Ask5           float64
	BidVol5        int
	AskVol5        int
	ReversedBytes4 uint16  // 保留
	ReversedBytes5 int     // 保留
	ReversedBytes6 int     // 保留
	ReversedBytes7 int     // 保留
	ReversedBytes8 int     // 保留
	Rate           float64 // 涨速
	Active2        uint16  // 活跃度
	TopNo          int     // 板块排名
	TopCode        string  // 领涨个股
	TopName        string  // 领涨个股名称
	TopRate        float64 // 领涨个股涨幅
	ZhanTing       int     // 涨停数
	Ling           int     // 平盘数
	Count          int     // 总数
}

// BatchSnapShot 批量获取即时行情数据快照
func BatchSnapShot(codes []string) []QuoteSnapshot {
	marketIds := []proto.Market{}
	symbols := []string{}
	for _, code := range codes {
		id, _, symbol := category.DetectMarket(code)
		if len(symbol) == 6 {
			marketIds = append(marketIds, id)
			symbols = append(symbols, symbol)
		}
	}
	tdxApi := prepare()
	list := []QuoteSnapshot{}
	hq, err := tdxApi.GetSecurityQuotes(marketIds, symbols)
	if err != nil {
		logger.Errorf("获取即时行情数据失败", err)
		return list
	}
	//fmt.Printf("%+v\n", hq)
	lastTradeday := time.Now().Format(category.INDEX_DATE)
	td := date.TradeRange("2023-01-01", lastTradeday)
	lastTradeday = td[len(td)-1]
	for _, v := range hq.List {
		snapshot := QuoteSnapshot{}
		_ = api.Copy(&snapshot, &v)
		if snapshot.LastClose == float64(0) {
			continue
		}
		if api.StartsWith(snapshot.Code, []string{"88"}) {
			snapshot.Code = "sh" + snapshot.Code
		} else {
			_, mname, mcode := category.DetectMarket(snapshot.Code)
			snapshot.Code = mname + mcode
		}
		list = append(list, snapshot)
	}
	return list
}
