package qtimg

type RealTimeData struct {
	Market       int     // [0] 代表交易所: 200-美股(us),100-港股(hk),51-深圳(sz),1-上海(sh)
	Name         string  // [1]股票名称
	Gid          string  // [2]股票编号
	NowPri       float64 // [3]当前价格
	YestClosePri float64 // [4]昨日收盘价
	OpeningPri   float64 // [5]今日开盘价
	TraNumber    int64   // [6]成交量
	Outter       int64   // [7]外盘
	Inner        int64   // [8]内盘
	BuyOne       int64   // [9]买一报价
	BuyOnePri    float64 // [10]买一
	BuyTwo       int64   // [11]买二
	BuyTwoPri    float64 // [12]买二报价
	BuyThree     int64   // [13]买三
	BuyThreePri  float64 // [14]买三报价
	BuyFour      int64   // [15]买四
	BuyFourPri   float64 // [16]买四报价
	BuyFive      int64   // [17]买五
	BuyFivePri   float64 // [18]买五报价
	SellOne      int64   // [19]卖一
	SellOnePri   float64 // [20]卖一报价
	SellTwo      int64   // [21]卖二
	SellTwoPri   float64 // [22]卖二报价
	SellThree    int64   // [23]卖三
	SellThreePri float64 // [24]卖三报价
	SellFour     int64   // [25]卖四
	SellFourPri  float64 // [26]卖四报价
	SellFive     int64   // [27]卖五
	SellFivePri  float64 // [28]卖五报价
	// [29]最近逐笔成交
	Time      string  // [30]时间
	Change    float64 // [31]涨跌
	ChangePer float64 // [32]涨跌%
	YodayMax  float64 // [33]今日最高价
	YodayMin  float64 // [34]今日最低价
	// [35]价格/成交量（手）/成交额
	TradeCount int64   // [36]成交量
	TradeAmont int64   // [37]成交额
	ChangeRate float64 // [38]换手率
	PERatio    float64 // [39]市盈率
	// [40]
	// [41]最高
	// [42]最低
	MaxMinChange float64 // [43]振幅
	MarketAmont  float64 // [44]流通市值
	TotalAmont   float64 // [45]总市值
	PBRatio      float64 // [46]市净率
	HighPri      float64 // [47]涨停价
	LowPri       float64 // [48]跌停价
}

type FundFlow struct {
	Gid    string  //[0] 代码
	BigIn  float64 //[1] 主力流入
	BigOut float64 //[2] 主力流出
	//[3] 主力净流入
	//[4] 主力净流入/资金流入流出总和
	SmallIn  float64 //[5] 散户流入
	SmallOut float64 //[6] 散户流出
	//[7] 散户净流入
	//[8] 散户净流入/资金流入流出总和
	//[9] 资金流入流出总和1+2+5+6
	//[10] 未知
	//[11] 未知
	Name string //[12] 名字
	Date string //[13] 日期
}

type PKData struct {
	BuyBig    float64 //0: 买盘大单
	BuySmall  float64 //1: 买盘小单
	SellBig   float64 //2: 卖盘大单
	SellSmall float64 //3: 卖盘小单
}

type StockInfo struct {
	//0: 未知
	Name       string  //1: 名字
	Gid        string  //2: 代码
	Price      float64 //3: 当前价格
	Change     float64 //4: 涨跌
	ChangePer  float64 //5: 涨跌%
	TradeCount float64 //6: 成交量（手）
	TradeAmont float64 //7: 成交额（万）
	//8:
	TotalAmont float64 //9: 总市值
}

type HistoryData struct {
	Date   string
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
}
