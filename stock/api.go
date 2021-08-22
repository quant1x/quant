package stock

import "time"

type long = int64

// 股票数据常量

const (
	// FIVE_MINUTES 五分钟
	FIVE_MINUTES = 5

	// FIFTEEN_MINUTES 十五分钟
	FIFTEEN_MINUTES = 15

	// THIRTY_MINUTES 三十分钟
	THIRTY_MINUTES = 30

	// ONE_HOUR 一小时
	ONE_HOUR = 60

	// ONE_DAY 一天
	ONE_DAY = 4 * ONE_HOUR

	// ONE_WEEK 一周
	ONE_WEEK = 7 * ONE_DAY

	// DEFAULT_DATALEN 全部数据
	DEFAULT_DATALEN = 1000000

	urlHistory = "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData"
	// 资金流向
	//http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/MoneyFlow.ssi_ssfx_flzjtj?daima=sh600072

	PrefixMessage = "【CTP微信助手】"
	SuffixMessage = " \r\n--以上内容，不构成任何投资建议，据此进行相关操作，风险自担。"

	T_STOCK = 0x10000000 // 个股
	T_INDEX = 0x20000000 // 指数

	D_OK    = 0x00000000 // 数据正常
	D_ERROR = 0x40000000 // 数据错误
	D_ECODE = 0x00000001 // 代码错误
	D_ENET  = 0x00000002 // 网络异常
	D_EDATA = 0x00000004 // 数据错误
	D_EDISK = 0x00000008 // 写文件错误
	D_ENEED = 0x00000010 // 不需要更新

	DefaultValue = 0.0000
)

// StockApi stock interface
type StockApi interface {
	// 实时接口
	RealTime0(code string) (*RealTime, error)
	RealTime(code string) (*RealTime, error)
}

//DayKLine struct implement
type DayKLine struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// DataApi 数据接口
type DataApi interface {
	Name() string
	// CompleteKLine 获取全部日线数据
	CompleteKLine(code string) (kline []DayKLine, err error)
	// DailyFromDate 从 start 开始补全日线数据
	DailyFromDate(code string, start time.Time) ([]DayKLine, time.Time, error)
}