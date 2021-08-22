package category

const (
	// USEC_PER_SEC number of microseconds per second
	USEC_PER_SEC int64 = 1000000
	// MsecPerSec number of milliseconds per second
	MsecPerSec int64 = 1000
	// SecondOfDay 一天的秒数
	SecondOfDay int64 = 24 * 60 * 60
	// MillisecondsOfDay 一天的毫秒数
	MillisecondsOfDay int64 = SecondOfDay * MsecPerSec
	// RealTimenterval 实时数据间隔时间, 单位毫秒
	RealTimenterval int64 = 5 * 1000

	// NullState 正常状态, 字符串"01"
	NullState   = "00"
	NormalState = "01"

	// 历史数据获取的时间, 时, 分, 秒
	HistoryUpdateHour = 17
	HistoryUpdateMin = 0
	HistoryUpdateSec = 0
)
