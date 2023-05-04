package trade

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
		snapshot.ZhangDieFu = ((snapshot.Price / snapshot.LastClose) - 1.00) * 100
		snapshot.LiuTongPan = cache.GetLiuTongPan(fullCode)
		snapshot.FreeGuBen = gbFee
		snapshot.TurnZ = 10000 * float64(snapshot.OpenVolume) / float64(gbFee)
		list = append(list, snapshot)
	}
	return list
}
