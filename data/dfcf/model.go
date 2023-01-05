package dfcf

type DfcfHistory struct {
	// date        时间
	Date string `json:"date" array:"0"`
	// open       开盘价
	Open float64 `json:"open" array:"1"`
	// high       最高价
	High float64 `json:"high" array:"3"`
	// low        最低价
	Low float64 `json:"low" array:"4"`
	// close      收盘价
	Close float64 `json:"close" array:"2"`
	// volume     成交量, 单位股, 除以100为手
	Volume int64 `json:"volume" array:"5"`
}
