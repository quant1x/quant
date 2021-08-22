package sina

type SinaHistory struct {
	// day        时间
	Date string `json:"day"`
	// open       开盘价
	Open string `json:"open"`
	// high       最高价
	High string `json:"high"`
	// low        最低价
	Low string `json:"low"`
	// close      收盘价
	Close string `json:"close"`
	// volume     成交量, 单位股, 除以100为手
	Volume string `json:"volume"`
	// MA5        五日平均价
	MA5 float64 `json:"ma_price5"`
	// MA5Volume  五日平均交易量
	MA5Volume int64 `json:"ma_volume5"`
	// MA10       十日平均价
	MA10 float64 `json:"ma_price10"`
	// MA10Volume 十日平均交易量
	MA10Volume int64 `json:"ma_volume10"`
	// MA30       三十日平均价
	MA30 float64 `json:"ma_price30"`
	// MA30Volume 三十日平均交易量
	MA30Volume int64 `json:"ma_volume30"`
}

// SinaHushenRealtime A股行情实时数据
type SinaHushenRealtime struct {
	Code           string  `json:"code"`
	Name           string  `name:"股票名称" array:"0"`
	Open           float64 `name:"今日开盘" array:"1"`
	YesterdayClose float64 `name:"昨日收盘" array:"2"`
	New            float64 `name:"当前价" array:"3"`
	High           float64 `name:"最高价" array:"4"`
	Low            float64 `name:"最低价" array:"5"`
	BuyPrice       float64 `name:"竞买价,买一" array:"6"`
	SellPrice      float64 `name:"竞卖价,卖一" array:"7"`
	Volume         int64   `name:"成交量" array:"8"`
	VolumePrice    float64 `name:"成交总金额" array:"9"`
	Buy1Num        int64   `name:"买一申请数" array:"10"`
	Buy1Price      float64 `name:"买一报价" array:"11"`
	Buy2Num        int64   `name:"买二申请数" array:"12"`
	Buy2Price      float64 `name:"买二报价" array:"13"`
	Buy3Num        int64   `name:"买三申请数" array:"14"`
	Buy3Price      float64 `name:"买三报价" array:"15"`
	Buy4Num        int64   `name:"买四申请数" array:"16"`
	Buy4Price      float64 `name:"买四报价" array:"17"`
	Buy5Num        int64   `name:"买五申请数" array:"18"`
	Buy5Price      float64 `name:"买五报价" array:"19"`
	Sell1Num       int64   `name:"卖一申请数" array:"20"`
	Sell1Price     float64 `name:"卖一报价" array:"21"`
	Sell2Num       int64   `name:"卖二申请数" array:"22"`
	Sell2Price     float64 `name:"卖二报价" array:"23"`
	Sell3Num       int64   `name:"卖三申请数" array:"24"`
	Sell3Price     float64 `name:"卖三报价" array:"25"`
	Sell4Num       int64   `name:"卖四申请数" array:"26"`
	Sell4Price     float64 `name:"卖四报价" array:"27"`
	Sell5Num       int64   `name:"卖五申请数" array:"28"`
	Sell5Price     float64 `name:"卖五报价" array:"29"`
	Date           string  `name:"日期" array:"30"`
	Time           string  `name:"时间" array:"31"`
}

//0: CHENMING PAPER,
//1: 晨鸣纸业,
//2: 4.630, Open
//3: 4.620, front close
//4: 4.660, High
//5: 4.600, Low
//6: 4.650, New
//7: 0.030, 涨跌
//8: 0.649, 涨跌率
//9: 4.63000, 买一
//10: 4.65000, 卖一
//11: 12848135, 成交额
//12: 2766750, 成交量
//13: 0.000,
//14: 0.000,
//15: 9.878, 52周最该
//16: 3.033, 52周最低
//17: 2021/08/12,
//18: 11:57

// SinaHonkongRealtime A股行情实时数据
type SinaHonkongRealtime struct {
	Code            string
	EnglishName     string  `name:"英文名称" array:"0"`
	Name            string  `name:"股票名称" array:"1"`
	Open            float64 `name:"今日开盘" array:"2"`
	YesterdayClose  float64 `name:"昨日收盘" array:"3"`
	High            float64 `name:"最高价" array:"4"`
	Low             float64 `name:"最低价" array:"5"`
	New             float64 `name:"当前价" array:"6"`
	RiseFall        float64 `name:"涨跌" array:"7"`
	RiseFallPercent float64 `name:"涨跌率" array:"8"`
	BuyPrice        float64 `name:"买一价格" array:"9"`
	SellPrice       float64 `name:"卖一价格" array:"10"`
	VolumePrice     float64 `name:"成交额" array:"11"`
	Volume          int64   `name:"成交量" array:"12"`
	Unknown13       float64 `name:"未知13" array:"13"`
	Unknown14       float64 `name:"未知13" array:"14"`
	High52          float64 `name:"最高价" array:"15"`
	Low52           float64 `name:"最低价" array:"16"`
	Date            string  `name:"日期" array:"17"`
	Time            string  `name:"时间" array:"18"`
}
