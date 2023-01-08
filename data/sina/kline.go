package sina

import (
	"encoding/json"
	"fmt"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/stock"
	"github.com/quant1x/quant/utils"
	"github.com/quant1x/quant/utils/http"
	"strconv"
	"time"
)

const (
	urlKLine    = "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=%s&scale=%d&datalen=%d"
	urlRealtime = "http://hq.sinajs.cn/list=%s"
)

func createUrl(code string, scale int, datalen int) string {
	return fmt.Sprintf(urlKLine, code, scale, datalen)
}

// GetHistory sina获取历史数据的唯一方法
func GetHistory(fullCode string, datalen int) ([]SinaHistory, error) {
	url := createUrl(fullCode, stock.ONE_DAY, datalen)

	data, err := http.HttpGet(url)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return nil, err
	}
	var kl []SinaHistory
	err = json.Unmarshal(data, &kl)
	if err != nil {
		logger.Errorf("data[%s], error=[%+v]\n", data, err)
		return nil, err
	}
	return kl, nil
}

// SinaDataApi 腾讯数据
type SinaDataApi struct {
}

func (this *SinaDataApi) Name() string {
	return "sina"
}

func (this *SinaDataApi) CompleteKLine(code string) ([]stock.DayKLine, error) {
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
	klines, _, err := this.DailyFromDate(code, startTime)
	return klines, err
}

func (this *SinaDataApi) DailyFromDate(code string, startTime time.Time) ([]stock.DayKLine, time.Time, error) {
	staticInfo, err := security.GetBasicInfo(code)
	if err != nil {
		return nil, time.Time{}, err
	}

	now := utils.CanUpdateTime()
	//now = utils.DateZero(now)
	listTime := time.Unix(int64(staticInfo.ListTimestamp), 0)

	// 计算需要补充多少年和多少天的数据
	if listTime.After(startTime) {
		startTime = listTime
	}
	days := utils.KLineRequireDays(now, startTime)
	var kLines []stock.DayKLine
	// 需要补充数据的最后一天
	nextTradingDay := utils.DateZero(startTime)
	// 测试时间比对
	//nextTradingDay = time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local)
	history, err := GetHistory(code, days)
	dls, lastDay := extract(nextTradingDay, history)
	//nextTradingDay = lastDay
	_ = lastDay
	kLines = append(kLines, dls...)
	listDay := listTime.Format(util.DateFormat)
	startDay := startTime.Format(util.DateFormat)
	endDay := now.Format(util.DateFormat)
	logger.Infof("%s[%s]: %s -> %s", code, listDay, startDay, endDay)
	return kLines, nextTradingDay, nil
}

// 转换行情数据为标准的K线数据
func extract(nextTradingDay time.Time, history []SinaHistory) ([]stock.DayKLine, time.Time) {
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
			dl.Open, _ = strconv.ParseFloat(item.Open, 64)
			dl.High, _ = strconv.ParseFloat(item.High, 64)
			dl.Low, _ = strconv.ParseFloat(item.Low, 64)
			dl.Close, _ = strconv.ParseFloat(item.Close, 64)
			dl.Volume, _ = strconv.ParseInt(item.Volume, 10, 64)

			kLines = append(kLines, dl)
		}
	}
	return kLines, nextTradingDay
}
