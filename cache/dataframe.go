package cache

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

func (this DataFrame) Offset(n int) (date []string, vOpen []float64, vClose []float64, vHigh []float64, vLow []float64, vVolume []int64) {
	end := n
	date = this.Date[:end]
	vOpen = this.Open[:end]
	vClose = this.Close[:end]
	vHigh = this.High[:end]
	vLow = this.Close[:end]
	vVolume = this.Volume[:end]
	return
}
