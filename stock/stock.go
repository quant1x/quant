package stock

// 东方财富
// 字段分析
// 0  2           1=上证/2=深成
// 1  002547      股票代码
// 2  春兴精工     股票名称
// 3  7.77        当前价格
// 4  0.71       上涨价格
// 5  10.06       涨幅
// 6  2062612     206万手成交量
// 7  1549976928  成交额
// 8  13.31       振幅
// 9  7.77        盘中最高
// 10  6.83        盘中最低
// 11  6.93        开盘价
// 12  7.06        昨收盘价
// 13  0.00        5分钟涨速?
// 14  1.96        量比
// 15  25.86       换手率
// 16  148.87      动态市盈率
// 17  3.19        市净率
// 18  8765004174  总市值
// 19  6197961277  流通市值
// 20  81.12%      60日涨幅
// 21  36.8%       年初至今涨幅
// 22  0.00        涨速?
// 23  上市日期
// 24  最新有效日期
//
type StockInfo struct {
	// 1-上证sh, 2-深圳sz
	Market string `json:"market" array:"0"`
	// 完整代码
	FullCode string `json:"full_code"`
	//1:  "名字",
	Name string `json:"name" array:"2"`
	//2:  "代码",
	Code string `json:"code" array:"1"`
	// 上市日期
	Start string `json:"start" array:"23"`
	// 最新有效时间
	Stop string `json:"end" array:"24"`
	// 最新价格
	New string `json:"new" array:"3"`
	// 最高价
	High string `json:"high" array:"9"`
	// 最低价
	Low string `json:"low" array:"10"`
	// 开盘价
	Open string `json:"open" array:"11"`
	// 昨日收盘
	FrontClose string `json:"front_close" array:"12"`
}
