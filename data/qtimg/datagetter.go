package qtimg

import (
	"fmt"
	"github.com/quant1x/quant/http"
	"github.com/mymmsc/gox/encoding"
	"strconv"
	"strings"
)

//某只某年股票日行情
//http://data.gtimg.cn/flashdata/hushen/daily/15/sz000750.js
//某只某年股票周行情
//http://data.gtimg.cn/flashdata/hushen/weekly/sz000868.js
//实时行情
//http://qt.gtimg.cn/q=sz000858
//资金流向
//http://qt.gtimg.cn/q=ff_sz000858
//盘口分析
//http://qt.gtimg.cn/q=s_pksz000858
//简要信息
//http://qt.gtimg.cn/q=s_sz000858

const (
	// UrlRealTime 实时数据
	UrlRealTime = "http://qt.gtimg.cn/q=%s"
	// UrlFundFlow 资金流向
	UrlFundFlow = "http://qt.gtimg.cn/q=ff_%s"
	// UrlPk 盘口
	UrlPk = "http://qt.gtimg.cn/q=s_pk%s"
	// UrlInfo 摘要
	UrlInfo = "http://qt.gtimg.cn/q=s_%s"
	// UrlDaily 日线
	UrlDaily = "http://data.gtimg.cn/flashdata/hushen/daily/%02d/%s.js"
	// UrlWeekly 周线
	UrlWeekly = "http://data.gtimg.cn/flashdata/hushen/weekly/%s.js"
)

func checkErr(err error) {}

func GetRealtime(code string) *RealTimeData {
	url := fmt.Sprintf(UrlRealTime, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := encoding.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	fmt.Println(strings.Join(dataArray, "\n"))
	data := new(RealTimeData)
	data.Name = strings.Replace(dataArray[1], " ", "", -1)
	data.Gid = dataArray[2]
	data.NowPri, _ = strconv.ParseFloat(dataArray[3], 64)
	data.YestClosePri, _ = strconv.ParseFloat(dataArray[4], 64)
	data.OpeningPri, _ = strconv.ParseFloat(dataArray[5], 64)
	data.TraNumber, _ = strconv.ParseInt(dataArray[6], 10, 64)
	data.Outter, _ = strconv.ParseInt(dataArray[7], 10, 64)
	data.Inner, _ = strconv.ParseInt(dataArray[8], 10, 64)
	data.BuyOnePri, _ = strconv.ParseFloat(dataArray[9], 64)
	data.BuyOne, _ = strconv.ParseInt(dataArray[10], 10, 64)
	data.BuyTwoPri, _ = strconv.ParseFloat(dataArray[11], 64)
	data.BuyTwo, _ = strconv.ParseInt(dataArray[12], 10, 64)
	data.BuyThreePri, _ = strconv.ParseFloat(dataArray[13], 64)
	data.BuyThree, _ = strconv.ParseInt(dataArray[14], 10, 64)
	data.BuyFourPri, _ = strconv.ParseFloat(dataArray[15], 64)
	data.BuyFour, _ = strconv.ParseInt(dataArray[16], 10, 64)
	data.BuyFivePri, _ = strconv.ParseFloat(dataArray[17], 64)
	data.BuyFive, _ = strconv.ParseInt(dataArray[18], 10, 64)
	data.SellOnePri, _ = strconv.ParseFloat(dataArray[19], 64)
	data.SellOne, _ = strconv.ParseInt(dataArray[20], 10, 64)
	data.SellTwoPri, _ = strconv.ParseFloat(dataArray[21], 64)
	data.SellTwo, _ = strconv.ParseInt(dataArray[22], 10, 64)
	data.SellThreePri, _ = strconv.ParseFloat(dataArray[23], 64)
	data.SellThree, _ = strconv.ParseInt(dataArray[24], 10, 64)
	data.SellFourPri, _ = strconv.ParseFloat(dataArray[25], 64)
	data.SellFour, _ = strconv.ParseInt(dataArray[26], 10, 64)
	data.SellFivePri, _ = strconv.ParseFloat(dataArray[27], 64)
	data.SellFive, _ = strconv.ParseInt(dataArray[28], 10, 64)
	data.Time = dataArray[30]
	data.Change, _ = strconv.ParseFloat(dataArray[31], 64)
	data.ChangePer, _ = strconv.ParseFloat(dataArray[32], 64)
	data.YodayMax, _ = strconv.ParseFloat(dataArray[33], 64)
	data.YodayMin, _ = strconv.ParseFloat(dataArray[34], 64)
	data.TradeCount, _ = strconv.ParseInt(dataArray[36], 10, 64)
	data.TradeAmont, _ = strconv.ParseInt(dataArray[37], 10, 64)
	data.ChangeRate, _ = strconv.ParseFloat(dataArray[38], 64)
	data.PERatio, _ = strconv.ParseFloat(dataArray[39], 64)
	data.MaxMinChange, _ = strconv.ParseFloat(dataArray[43], 64)
	data.MarketAmont, _ = strconv.ParseFloat(dataArray[44], 64)
	data.TotalAmont, _ = strconv.ParseFloat(dataArray[45], 64)
	data.PBRatio, _ = strconv.ParseFloat(dataArray[46], 64)
	data.HighPri, _ = strconv.ParseFloat(dataArray[47], 64)
	data.LowPri, _ = strconv.ParseFloat(dataArray[48], 64)
	return data
}

func GetPK(code string) *PKData {
	url := fmt.Sprintf(UrlPk, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := encoding.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(PKData)
	data.BuyBig, _ = strconv.ParseFloat(strings.Split(dataArray[0], "\"")[1], 64)
	data.BuySmall, _ = strconv.ParseFloat(dataArray[1], 64)
	data.SellBig, _ = strconv.ParseFloat(dataArray[2], 64)
	data.SellSmall, _ = strconv.ParseFloat(strings.Split(dataArray[3], "\"")[0], 64)
	return data

}

// GetFundFlow 资金流向
func GetFundFlow(code string) *FundFlow {
	url := fmt.Sprintf(UrlFundFlow, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := encoding.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(FundFlow)
	data.Gid = code
	data.BigIn, _ = strconv.ParseFloat(dataArray[1], 64)
	data.BigOut, _ = strconv.ParseFloat(dataArray[2], 64)
	data.SmallIn, _ = strconv.ParseFloat(dataArray[5], 64)
	data.SmallOut, _ = strconv.ParseFloat(dataArray[6], 64)
	data.Name = dataArray[12]
	data.Date = dataArray[13]
	return data
}

func GetInfo(code string) *StockInfo {
	url := fmt.Sprintf(UrlInfo, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := encoding.NewDecoder("gbk").ConvertString(string(body))
	dataArray := strings.Split(res, "~")
	data := new(StockInfo)
	data.Name = dataArray[1]
	data.Gid = dataArray[2]
	data.Price, _ = strconv.ParseFloat(dataArray[3], 64)
	data.Change, _ = strconv.ParseFloat(dataArray[4], 64)
	data.ChangePer, _ = strconv.ParseFloat(dataArray[5], 64)
	data.TradeCount, _ = strconv.ParseFloat(dataArray[6], 64)
	data.TradeAmont, _ = strconv.ParseFloat(dataArray[7], 64)
	data.TotalAmont, _ = strconv.ParseFloat(strings.Split(dataArray[9], "\"")[0], 64)
	return data
}

// 获取某一年的日线数据
func GetDaily(code string, year int) []*HistoryData {
	url := fmt.Sprintf(UrlDaily, year%100, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := string(body)
	dataArray := strings.Split(res, "\\n\\")
	list := []*HistoryData{}
	for index, str := range dataArray {
		if index == 0 || index == len(dataArray)-1 {
			// 跳过第一行和最后一行
		} else {
			data := strings.Split(str, " ")
			entity := new(HistoryData)
			date := strings.Replace(data[0], "\n", "", -1)
			entity.Date = fmt.Sprintf("%04d-%2s-%2s", year, date[2:4], date[4:6])
			entity.Open, _ = strconv.ParseFloat(data[1], 64)
			entity.Close, _ = strconv.ParseFloat(data[2], 64)
			entity.High, _ = strconv.ParseFloat(data[3], 64)
			entity.Low, _ = strconv.ParseFloat(data[4], 64)
			entity.Volume, _ = strconv.ParseFloat(data[5], 64)
			list = append(list, entity)
		}
	}
	return list
}

func GetWeekly(code string) []*HistoryData {
	url := fmt.Sprintf(UrlWeekly, code)
	body, err := http.HttpGet(url)
	if body == nil || err != nil {
		return nil
	}
	res := string(body)
	dataArray := strings.Split(res, "\\n\\")
	list := []*HistoryData{}
	for index, str := range dataArray {
		if index == 0 || index == len(dataArray)-1 {

		} else {
			data := strings.Split(str, " ")
			entity := new(HistoryData)
			entity.Date = strings.Replace(data[0], "\n", "", -1)
			entity.Open, _ = strconv.ParseFloat(data[1], 64)
			entity.Close, _ = strconv.ParseFloat(data[2], 64)
			entity.High, _ = strconv.ParseFloat(data[3], 64)
			entity.Low, _ = strconv.ParseFloat(data[4], 64)
			entity.Volume, _ = strconv.ParseFloat(data[5], 64)
			list = append(list, entity)
		}
	}
	return list
}
