package Database

import (
	"time"
)

type StockCode struct {
	Flag       string    `json:"flag" xorm:"default '01' comment('标志($): 00-禁止检测,01-正常检测') CHAR(2)"`
	Code       string    `json:"code" xorm:"not null comment('股票代码') index CHAR(32)"`
	FullCode   string    `json:"full_code" xorm:"not null comment('完整的股票代码') CHAR(32)"`
	Name       string    `json:"name" xorm:"not null comment('股票名称(?$)') CHAR(128)"`
	Operator   string    `json:"operator" xorm:"default 'system' comment('操作人(?$)') VARCHAR(50)"`
	Createtime time.Time `json:"createTime" xorm:"comment('创建时间($)') DATETIME"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
}
