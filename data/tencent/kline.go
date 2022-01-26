package tencent

import (
	"fmt"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/fastjson"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util"
	"github.com/quant1x/quant/data/qtimg"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/http"
	"github.com/quant1x/quant/stock"
	"github.com/quant1x/quant/utils"
	"strings"
	"time"
)

const (
	urlKLine = "https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%s,day,,,%d,qfq"
)

var ()

func createUrl(code string, datalen int) string {
	return fmt.Sprintf(urlKLine, code, datalen)
}

// 获取历史数据 - 通过天数
func historyByDays(fullCode string, datalen int) []TencentHistory {
	url := createUrl(fullCode, datalen)
	data, err := http.HttpGet(url)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return nil
	}

	obj, err := fastjson.ParseBytes(data)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return nil
	}
	errCode := obj.GetInt("code")
	if errCode != 0 {
		logger.Errorf("%d: %s\n", err, obj.GetString("msg"))
		return nil
	}
	_ = data
	biz := obj.GetObject("data")
	if biz == nil {
		logger.Errorf("数据非法\n")
		return nil
	}

	bizData := biz.Get(fullCode)
	if bizData == nil {
		logger.Errorf("数据非法\n")
		return nil
	}

	history := bizData.GetArray("day")
	if history == nil {
		history = bizData.GetArray("qfqday")
	}
	if history == nil {
		logger.Errorf("数据非法\n")
		return nil
	}
	var kl []TencentHistory
	for _, item := range history {
		dd, err := item.Array()

		if err != nil {
			continue
		}
		var hd []string
		for _, d0 := range dd {
			if d0.Type() != fastjson.TypeString {
				continue
			}
			sb, err := d0.StringBytes()
			if err != nil {
				logger.Fatalf("cannot obtain string: %s", err)
			}

			hd = append(hd, string(sb))
		}
		var kl0 TencentHistory
		err = api.Convert(hd, &kl0)
		if err == nil {
			kl = append(kl, kl0)
		}
	}
	//fmt.Printf("1. %+v\n", kl)
	return kl
}

// 获取某一年的数据
func historyByOneYear(code string, year int) []TencentHistory {
	yl := qtimg.GetDaily(code, year)
	var history []TencentHistory
	if len(yl) > 0 {
		for _, item := range yl {
			dl := TencentHistory{
				Date:   item.Date,
				Open:   item.Open,
				Close:  item.Close,
				High:   item.High,
				Low:    item.Low,
				Volume: int64(item.Volume),
			}
			history = append(history, dl)
		}
	}
	return history
}

// TencentDataApi 腾讯数据
type TencentDataApi struct {
}

func (this *TencentDataApi) Name() string {
	return "tencent"
}

// CompleteKLine 补全K线数据
func (this *TencentDataApi) CompleteKLine(code string) ([]stock.DayKLine, error) {
	// 从什么时间开始, 到什么时间结束
	// 0.1. 按照code所在的市场的休市时间, 到当前时间截止
	// 0.2. 补全全部的数据
	// 0.3. 腾讯的数据分两种, 一种是静态数据, 一种是动态接口
	// 1 获取个股的基本信息, 主要获取市场, 上市时间, 是否退市
	staticInfo, err := security.GetBasicInfo(code)
	if err != nil {
		return nil, err
	}

	//now := time.Now()
	//now = utils.DateZero(now)
	//now := utils.CanUpdateTime()
	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)

	// 计算需要补充多少年和多少天的数据
	startTime := listTime
	var kLines []stock.DayKLine
	var dataLastDay time.Time
	kLines, dataLastDay, err = this.DailyFromDate(code, startTime)
	if err != nil {
		return nil, err
	}
	listDay := listTime.Format(util.DateFormat)
	startDay := startTime.Format(util.DateFormat)
	endDay := dataLastDay.Format(util.DateFormat)
	logger.Infof("%s[%s]: %s -> %s", code, listDay, startDay, endDay)
	return kLines, nil
}

func (this *TencentDataApi) DailyFromDate(code string, startTime time.Time) ([]stock.DayKLine, time.Time, error) {
	staticInfo, err := security.GetBasicInfo(code)
	if err != nil {
		return nil, time.Time{}, err
	}

	//now := time.Now()
	//now = utils.DateZero(now)
	now := utils.CanUpdateTime()
	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)

	// 计算需要补充多少年和多少天的数据
	if listTime.After(startTime) {
		startTime = listTime
	}
	years, days := calculateRemainingDays(now, startTime)
	var kLines []stock.DayKLine
	// 需要补充数据的最后一天
	nextTradingDay := utils.DateZero(startTime)
	if !strings.HasPrefix(code, "hk") {
		for i := 0; i < years; i++ {
			// https://data.gtimg.cn/flashdata/hushen/daily/10/sz399006.js
			klinesOfOneYear := historyByOneYear(code, startTime.Year()+i)
			dls, lastDay := extract(nextTradingDay, klinesOfOneYear)
			nextTradingDay = lastDay
			kLines = append(kLines, dls...)
		}
	}
	// 测试时间比对
	//nextTradingDay = time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local)
	days = utils.KLineRequireDays(now, startTime)
	history := historyByDays(code, days)
	dls, lastDay := extract(nextTradingDay, history)
	nextTradingDay = lastDay
	_ = lastDay
	kLines = append(kLines, dls...)
	listDay := listTime.Format(util.DateFormat)
	startDay := startTime.Format(util.DateFormat)
	endDay := nextTradingDay.Format(util.DateFormat)
	logger.Infof("%s[%s]: %s -> %s", code, listDay, startDay, endDay)
	return kLines, nextTradingDay, nil
}

func (this *TencentDataApi) DailyByDays(code string, days int) ([]stock.DayKLine, error) {
	staticInfo, err := security.GetBasicInfo(code)
	if err != nil {
		return nil, err
	}

	//now := time.Now()
	//now = utils.DateZero(now)
	//startTime := now.AddDate(0, 0, -1*days)
	startTime := utils.CanUpdateTime()
	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)

	// 计算需要补充多少年和多少天的数据
	if listTime.After(startTime) {
		startTime = listTime
	}
	var kLines []stock.DayKLine
	// 需要补充数据的最后一天
	nextTradingDay := utils.DateZero(startTime)
	history := historyByDays(code, days)
	dls, lastDay := extract(nextTradingDay, history)
	nextTradingDay = lastDay
	kLines = append(kLines, dls...)
	listDay := listTime.Format(util.DateFormat)
	startDay := startTime.Format(util.DateFormat)
	endDay := nextTradingDay.Format(util.DateFormat)
	logger.Infof("%s[%s]: %s -> %s", code, listDay, startDay, endDay)
	return kLines, nil
}

// 转换行情数据为标准的K线数据
func extract(nextTradingDay time.Time, history []TencentHistory) ([]stock.DayKLine, time.Time) {
	var kLines []stock.DayKLine
	if len(history) > 0 {
		for _, item := range history {
			_lastDay, _ := utils.ParseTime(item.Date)
			_lastDay = utils.DateZero(_lastDay)
			if _lastDay.Before(nextTradingDay) {
				continue
			}
			nextTradingDay = _lastDay.AddDate(0, 0, 1)
			var dl stock.DayKLine
			dl.Date = item.Date
			dl.Open = item.Open
			dl.Close = item.Close
			dl.High = item.High
			dl.Low = item.Low
			dl.Volume = item.Volume
			kLines = append(kLines, dl)
		}
	}
	return kLines, nextTradingDay
}

// 需要补充多少年的数据, 当年生需要多少天
// 腾讯自选股的数据, 需要分两段进行, 跨年的历史数据用年数据的js
func calculateRemainingDays(t1, t2 time.Time) (int, int) {
	// 校验时间前后
	if t1.After(t2) {
		t1, t2 = t2, t1
	}

	now := time.Now()
	currentYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	// 需要补多少年的数据
	var years int = 0
	// 计算多少年
	if currentYear.After(t1) {
		years = currentYear.Year() - t1.Year()
	}
	// 需要补当年多少天的数据
	var days int = 0
	if years == 0 {
		days = utils.KLineRequireDays(t2, t1)
	} else {
		days = utils.KLineRequireDays(t2, currentYear)
	}

	return years, days
}
