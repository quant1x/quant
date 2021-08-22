package stock

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
