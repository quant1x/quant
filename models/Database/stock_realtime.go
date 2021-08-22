package Database

import (
	"time"
)

type StockRealtime struct {
	Type            int       `json:"type" xorm:"default 2 comment('类型: 1-指数,2-股票') TINYINT(4)"`
	Date            time.Time `json:"date" xorm:"comment('日期') DATE"`
	Time            time.Time `json:"time" xorm:"comment('时间') TIME"`
	Code            string    `json:"code" xorm:"not null comment('证券代码') index CHAR(32)"`
	Name            string    `json:"name" xorm:"not null comment('证券名称') CHAR(32)"`
	Open            string    `json:"open" xorm:"default '0.000' comment('开盘价') VARCHAR(20)"`
	Close           string    `json:"close" xorm:"default '0.000' comment('收盘价') VARCHAR(20)"`
	Now             string    `json:"now" xorm:"default '0.000' comment('最新价') VARCHAR(20)"`
	High            string    `json:"high" xorm:"default '0.000' comment('最高价') VARCHAR(20)"`
	Low             string    `json:"low" xorm:"default '0.000' comment('最低价') VARCHAR(20)"`
	BuyPrice        string    `json:"buy_price" xorm:"default '0.000' comment('买入价') VARCHAR(20)"`
	SellPrice       string    `json:"sell_price" xorm:"default '0.000' comment('卖出价') VARCHAR(20)"`
	Volume          string    `json:"volume" xorm:"default '0' comment('成交量') VARCHAR(20)"`
	VolumePrice     string    `json:"volume_price" xorm:"default '0.000' comment('成交额') VARCHAR(20)"`
	Buy1Num         string    `json:"buy_1_num" xorm:"default '0' comment('委托买一量') VARCHAR(20)"`
	Buy1Price       string    `json:"buy_1_price" xorm:"default '0.000' comment('委托买一价') VARCHAR(20)"`
	Buy2Num         string    `json:"buy_2_num" xorm:"default '0' comment('委托买二量') VARCHAR(20)"`
	Buy2Price       string    `json:"buy_2_price" xorm:"default '0.000' comment('委托买二价') VARCHAR(20)"`
	Buy3Num         string    `json:"buy_3_num" xorm:"default '0' comment('委托买三量') VARCHAR(20)"`
	Buy3Price       string    `json:"buy_3_price" xorm:"default '0.000' comment('委托买三价') VARCHAR(20)"`
	Buy4Num         string    `json:"buy_4_num" xorm:"default '0' comment('委托买四量') VARCHAR(20)"`
	Buy4Price       string    `json:"buy_4_price" xorm:"default '0.000' comment('委托买四价') VARCHAR(20)"`
	Buy5Num         string    `json:"buy_5_num" xorm:"default '0' comment('委托买五量') VARCHAR(20)"`
	Buy5Price       string    `json:"buy_5_price" xorm:"default '0.000' comment('委托买五价') VARCHAR(20)"`
	Sell1Num        string    `json:"sell_1_num" xorm:"default '0' comment('委托卖一量') VARCHAR(20)"`
	Sell1Price      string    `json:"sell_1_price" xorm:"default '0.000' comment('委托卖一价') VARCHAR(20)"`
	Sell2Num        string    `json:"sell_2_num" xorm:"default '0' comment('委托卖二量') VARCHAR(20)"`
	Sell2Price      string    `json:"sell_2_price" xorm:"default '0.000' comment('委托卖二价') VARCHAR(20)"`
	Sell3Num        string    `json:"sell_3_num" xorm:"default '0' comment('委托卖三量') VARCHAR(20)"`
	Sell3Price      string    `json:"sell_3_price" xorm:"default '0.000' comment('委托卖三价') VARCHAR(20)"`
	Sell4Num        string    `json:"sell_4_num" xorm:"default '0' comment('委托卖四量') VARCHAR(20)"`
	Sell4Price      string    `json:"sell_4_price" xorm:"default '0.000' comment('委托卖四价') VARCHAR(20)"`
	Sell5Num        string    `json:"sell_5_num" xorm:"default '0' comment('委托卖五量') VARCHAR(20)"`
	Sell5Price      string    `json:"sell_5_price" xorm:"default '0.000' comment('委托买五价') VARCHAR(20)"`
	RiseFall        string    `json:"rise_fall" xorm:"default '0.000' comment('涨跌价') VARCHAR(20)"`
	RiseFallPercent string    `json:"rise_fall_percent" xorm:"default '0.000' comment('涨跌幅') VARCHAR(20)"`
}
