package Database

import (
	"time"
)

type StockMessages struct {
	Flag       string    `json:"flag" xorm:"default '00' comment('标志($): 00-未发送,01-正常发送,97-丢弃的策略消息') CHAR(2)"`
	Phone      string    `json:"phone" xorm:"not null comment('客户ID') index CHAR(32)"`
	Code       string    `json:"code" xorm:"not null comment('股票代码') index CHAR(32)"`
	Policy     string    `json:"policy" xorm:"default '' comment('策略') index VARCHAR(512)"`
	Price      string    `json:"price" xorm:"default '' comment('交易价格') CHAR(20)"`
	Remark     string    `json:"remark" xorm:"comment('策略命中备注') TEXT"`
	Createtime time.Time `json:"createTime" xorm:"comment('创建时间($)') DATETIME"`
	Senddate   time.Time `json:"sendDate" xorm:"comment('发送日期') DATE"`
	Operator   string    `json:"operator" xorm:"default 'system' comment('操作人(?$)') VARCHAR(50)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
}
