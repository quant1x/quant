package http

import (
	"testing"
)

func TestHttpGet0(t *testing.T) {
	HttpGet0("http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sh000001&scale=60&datalen=1000000")
}

func TestHttpGet(t *testing.T) {
	HttpGet("http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sh000001&scale=60&datalen=1000000")
}
