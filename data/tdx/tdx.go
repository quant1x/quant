package tdx

import (
	tdx "gitee.com/quant1x/gotdx"
	"gitee.com/quant1x/gotdx/proto"
	"strings"
)

type TdxApi struct {
	tdx.TcpClient
}

var (
	_client *tdx.TcpClient = nil
)

func NewApi() *tdx.TcpClient {
	host := "119.147.212.81"
	port := 7709
	_client := tdx.NewClient(&tdx.Opt{Host: host, Port: port})
	if _client != nil {
		_client.Connect()
	}
	return _client
}

func prepare() *tdx.TcpClient {
	if _client == nil {
		_client = NewApi()
	}
	return _client
}

func startsWith(str string, prefixs []string) bool {
	if len(str) == 0 || len(prefixs) == 0 {
		return false
	}
	for _, prefix := range prefixs {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// 判断股票ID对应的证券市场匹配规则
//
// ['50', '51', '60', '90', '110'] 为 sh
// ['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
// ['5', '6', '9'] 开头的为 sh， 其余为 sz
func getStockMarket(symbol string) string {
	//:param string: False 返回市场ID，否则市场缩写名称
	//:param symbol: 股票ID, 若以 'sz', 'sh' 开头直接返回对应类型，否则使用内置规则判断
	//:return 'sh' or 'sz'

	market := "sh"
	if startsWith(symbol, []string{"sh", "sz", "SH", "SZ"}) {
		market = strings.ToLower(symbol[0:2])
	} else if startsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if startsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"4", "8"}) {
		market = "bj"
	}
	return market
}

func getStockMarketId(symbol string) uint8 {
	market := getStockMarket(symbol)
	marketId := tdx.MARKET_SH
	if market == "sh" {
		marketId = tdx.MARKET_SH
	} else if market == "sz" {
		marketId = tdx.MARKET_SZ
	} else if market == "bj" {
		marketId = tdx.MARKET_BJ
	}
	//# logger.debug(f"market => {market}")

	return uint8(marketId)
}

// GetKLine 获取日K线
func GetKLine(code string, start uint16, count uint16) *proto.SecurityBarsReply {
	client := prepare()
	marketId := getStockMarketId(code)
	data, _ := client.GetSecurityBars(tdx.KLINE_TYPE_RI_K, marketId, code, start, count)
	return data
}
