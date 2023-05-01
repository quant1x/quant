package internal

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/category"
	"gitee.com/quant1x/data/category/trading"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"time"
)

var (
	stdApi *quotes.StdApi = nil
)

func prepare() *quotes.StdApi {
	if stdApi == nil {
		api_, err := quotes.NewStdApi()
		if err != nil {
			return nil
		}
		stdApi = api_
	}
	return stdApi
}

type QuoteSnapshot struct {
	Market          uint8   // 市场
	Code            string  // 代码
	Name            string  // 证券名称
	Active1         uint16  // 活跃度
	Price           float64 // 现价
	LastClose       float64 // 昨收
	ZhangDieFu      float64 // 涨跌幅
	Open            float64 // 开盘
	High            float64 // 最高
	Low             float64 // 最低
	ServerTime      string  // 时间
	ReversedBytes0  int     // 保留(时间 ServerTime)
	ReversedBytes1  int     // 保留
	Vol             int     // 总量
	CurVol          int     // 现量
	Amount          float64 // 总金额
	SVol            int     // 内盘
	BVol            int     // 外盘
	IndexOpenAmount int     // 指数-集合竞价成交金额=开盘成交金额
	StockOpenAmount int     // 个股-集合竞价成交金额=开盘成交金额
	OpenVolume      int     // 集合竞价-开盘量, 单位是股
	Bid1            float64
	Ask1            float64
	BidVol1         int
	AskVol1         int
	Bid2            float64
	Ask2            float64
	BidVol2         int
	AskVol2         int
	Bid3            float64
	Ask3            float64
	BidVol3         int
	AskVol3         int
	Bid4            float64
	Ask4            float64
	BidVol4         int
	AskVol4         int
	Bid5            float64
	Ask5            float64
	BidVol5         int
	AskVol5         int
	ReversedBytes4  uint16  // 保留
	ReversedBytes5  int     // 保留
	ReversedBytes6  int     // 保留
	ReversedBytes7  int     // 保留
	ReversedBytes8  int     // 保留
	Rate            float64 // 涨速
	Active2         uint16  // 活跃度
	TopNo           int     // 板块排名
	TopCode         string  // 领涨个股
	TopName         string  // 领涨个股名称
	TopRate         float64 // 领涨个股涨幅
	ZhanTing        int     // 涨停数
	Ling            int     // 平盘数
	Count           int     // 总数
	LiuTongPan      float64 // 流通盘
	FreeGuBen       float64 // 自由流通股本
	TurnZ           float64 // 开盘换手
}

// BatchSnapShot 批量获取即时行情数据快照
func BatchSnapShot(codes []string) []QuoteSnapshot {
	marketIds := []proto.MarketType{}
	symbols := []string{}
	for _, code := range codes {
		id, _, symbol := category.DetectMarket(code)
		if len(symbol) == 6 {
			marketIds = append(marketIds, id)
			symbols = append(symbols, symbol)
		}
	}
	tdxApi := prepare()
	list := []QuoteSnapshot{}
	hq, err := tdxApi.GetSecurityQuotes(marketIds, symbols)
	if err != nil {
		logger.Errorf("获取即时行情数据失败", err)
		return list
	}
	//fmt.Printf("%+v\n", hq)
	lastTradeday := time.Now().Format(category.INDEX_DATE)
	td := trading.TradeRange("2023-01-01", lastTradeday)
	lastTradeday = td[len(td)-1]
	for _, v := range hq.List {
		snapshot := QuoteSnapshot{}
		_ = api.Copy(&snapshot, &v)
		if snapshot.LastClose == float64(0) {
			continue
		}

		fullCode := category.GetMarketName(v.Market) + v.Code
		gbFee := cache.GetFreeGuBen(fullCode)
		snapshot.Code = fullCode
		snapshot.TurnZ = 100 * float64(snapshot.OpenVolume) / float64(gbFee)
		snapshot.ZhangDieFu = ((snapshot.Price / snapshot.LastClose) - 1.00) * 100
		list = append(list, snapshot)
	}
	return list
}
