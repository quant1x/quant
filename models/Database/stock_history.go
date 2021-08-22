package Database

import (
	"time"
)

type StockHistory struct {
	Day        time.Time `json:"day" xorm:"comment('日期') index DATE"`
	Code       string    `json:"code" xorm:"not null comment('股票代码') index VARCHAR(20)"`
	Open       string    `json:"open" xorm:"not null comment('开盘价') VARCHAR(20)"`
	High       string    `json:"high" xorm:"not null comment('最高价') VARCHAR(20)"`
	Low        string    `json:"low" xorm:"not null comment('最低价') VARCHAR(20)"`
	Close      string    `json:"close" xorm:"not null comment('收盘价') VARCHAR(20)"`
	Volume     string    `json:"volume" xorm:"not null comment('成交量') VARCHAR(20)"`
	Ma5        string    `json:"MA5" xorm:"not null comment('MA5价') VARCHAR(20)"`
	Ma5Volume  string    `json:"MA5_volume" xorm:"not null comment('MA5量') VARCHAR(20)"`
	Ma10       string    `json:"MA10" xorm:"not null comment('MA10价') VARCHAR(20)"`
	Ma10Volume string    `json:"MA10_volume" xorm:"not null comment('MA10量') VARCHAR(20)"`
	Ma30       string    `json:"MA30" xorm:"not null comment('MA30价') VARCHAR(20)"`
	Ma30Volume string    `json:"MA30_volume" xorm:"not null comment('MA30量') VARCHAR(20)"`
}
