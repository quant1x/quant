package stock

// DataFrame 数据帧
type DataFrame struct {
	Length int       `json:"length"`
	Date   []string  `json:"date"`
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
	Volume []int64   `json:"volume"`
}
