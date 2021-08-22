package tencent

import (
	"github.com/quant1x/quant/stock"
	"github.com/mymmsc/gox/api"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
)

type TencentStock struct {
}

func (ts *TencentStock) RealTime(code string) (*stock.RealTime, error) {
	res, err := http.Get(url + code)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return nil, err
	}
	bs := string(b)
	begin := strings.Index(bs, "\"")
	end := strings.LastIndex(bs, "~")
	if begin == -1 || end == -1 {
		return nil, invalidResponse
	}
	result := strings.Trim(bs[begin+1:end], "\r\n ")
	if result == "" {
		return nil, invalidResponse
	}
	pairs := strings.Split(result, "~")
	rt := &stock.RealTime{
		FullCode:code,
	}
	err = api.Convert(pairs, rt)
	return rt, err
}

func (ts *TencentStock) RealTime0(code string) (*stock.RealTime, error) {
	res, err := http.Get(url + code)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return nil, err
	}
	bs := string(b)
	begin := strings.Index(bs, "\"")
	end := strings.LastIndex(bs, "~")
	if begin == -1 || end == -1 {
		return nil, invalidResponse
	}
	result := strings.Trim(bs[begin+1:end], "\r\n ")
	if result == "" {
		return nil, invalidResponse
	}
	pairs := strings.Split(result, "~")
	/*if len(pairs) != len(names) {
		return nil, invalidResponse
	}*/
	rt := &stock.RealTime{
		FullCode:code,
	}
	//0:  "未知",
	rt.UnknownCode = pairs[0]
	//1:  "名字",
	rt.Name = pairs[1]
	//2:  "代码",
	rt.Code = pairs[2]
	//3:  "当前价格",
	rt.New = pairs[3]
	//4:  "昨收",
	rt.Close = pairs[4]
	//5:  "今开",
	rt.Open = pairs[5]
	//6:  "成交量（手)",
	rt.Volume  = pairs[6]
	//7:  "外盘",
	rt.OuterVol = pairs[7]
	//8:  "内盘",
	rt.InnerVol  = pairs[8]
	//9:  "买一",
	rt.Buy1Price  = pairs[9]
	//10: "买一量（手）",
	rt.Buy1Vol  = pairs[10]
	//11: "买二",
	rt.Buy2Price = pairs[11]
	//12: "买二量（手）",
	rt.Buy2Vol = pairs[12]
	//13: "买三",
	rt.Buy3Price = pairs[13]
	//14: "买三量（手）",
	rt.Buy3Vol  = pairs[14]
	//15: "买四",
	rt.Buy4Price = pairs[15]
	//16: "买四量（手）",
	rt.Buy4Vol = pairs[16]
	//17: "买五",
	rt.Buy5Price = pairs[17]
	//18: "买五量（手）",
	rt.Buy5Vol = pairs[18]
	//19: "卖一",
	rt.Sell1Price = pairs[19]
	//20: "卖一量",
	rt.Sell1Vol = pairs[20]
	//21: "卖二",
	rt.Sell2Price = pairs[21]
	//22: "卖二量",
	rt.Sell2Vol = pairs[22]
	//23: "卖三",
	rt.Sell3Price = pairs[23]
	//24: "卖三量",
	rt.Sell3Vol = pairs[24]
	//25: "卖四",
	rt.Sell4Price = pairs[25]
	//26: "卖四量",
	rt.Sell4Vol = pairs[26]
	//27: "卖五",
	rt.Sell5Price = pairs[27]
	//28: "卖五量",
	rt.Sell5Vol = pairs[28]
	//29: "最近逐笔成交",
	rt.Deals = pairs[29]
	//30: "时间",
	rt.Time = pairs[30]
	//31: "涨跌",
	rt.RiseFall = pairs[31]
	//32: "涨跌%",
	rt.RiseFallPercent = pairs[32]
	//33: "最高",
	rt.High = pairs[33]
	//34: "最低",
	rt.Low = pairs[34]
	//35: "价格/成交量（手）/成交额",
	rt.TransactionInformation = pairs[35]
	//36: "成交量（手）",
	rt.Volume1 = pairs[36]
	//37: "成交额（万）",
	rt.Amount = pairs[37]
	//38: "换手率",
	rt.TurnoverRate = pairs[38]
	//39: "市盈率",
	rt.PeRatio = pairs[39]
	//40: "未知",
	rt.Unknown = pairs[40]
	//41: "最高",
	rt.High1 = pairs[41]
	//42: "最低",
	rt.Low1 = pairs[42]
	//43: "振幅",
	rt.Amplitude = pairs[43]
	//44: "流通市值",
	rt.FreeFloatMarketValue = pairs[44]
	//45: "总市值",
	rt.TotalMarketValue = pairs[45]
	//46: "市净率",
	rt.MarketRate = pairs[46]
	//47: "涨停价",
	rt.LimitUp = pairs[47]
	//48: "跌停价",
	rt.LimitDown = pairs[48]
	return rt, nil
}