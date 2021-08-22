package Database

import (
	"time"
)

type StockUser struct {
	Flag       string    `json:"flag" xorm:"default '00' comment('标志($): 00-禁止,01-正常') CHAR(2)"`
	Memberid   string    `json:"memberId" xorm:"default '' comment('客户ID') index CHAR(128)"`
	Membername string    `json:"memberName" xorm:"default '' comment('客户姓名') index CHAR(128)"`
	Phone      string    `json:"phone" xorm:"not null comment('手机号码(?$)') index CHAR(32)"`
	Weixin     string    `json:"weixin" xorm:"default '' comment('微信id(?$)') CHAR(128)"`
	Email      string    `json:"email" xorm:"default '' comment('邮箱(?$)') CHAR(128)"`
	Createtime time.Time `json:"createTime" xorm:"comment('创建时间($)') DATETIME"`
	Senddate   time.Time `json:"sendDate" xorm:"comment('发送日期') DATE"`
	Operator   string    `json:"operator" xorm:"default 'system' comment('操作人(?$)') VARCHAR(50)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
}
