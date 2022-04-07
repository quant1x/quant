package main

import (
	"flag"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/data/sina"
	"github.com/quant1x/quant/data/tencent"
	"github.com/quant1x/quant/models/Cache"
	"github.com/quant1x/quant/stock"
	"github.com/quant1x/quant/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 更新日线数据工具
func main() {
	defer logger.FlushLogger()
	var (
		path   string // 数据路径
		useCSV bool   // 是否使用CSV格式
	)
	flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.BoolVar(&useCSV, "csv", false, "use CSV format")
	flag.Parse()
	cache.Init(path, useCSV)

	fullCodes := data.GetCodeList()
	for _, code := range fullCodes {
		basicInfo, err := security.GetBasicInfo(code)
		if err == security.ErrCodeNotExist {
			// 证券代码不存在
			continue
		}
		if err != nil {
			// 其它错误 输出错误信息
			logger.Errorf("%s => %v", code, err)
			continue
		}
		if basicInfo.Delisting {
			logger.Errorf("%s => 已退市", code)
			continue
		}
		listTimestamp := int64(basicInfo.ListTimestamp)
		e := pullData(code, utils.UnixTime(listTimestamp))
		if e&stock.D_ENEED == 0 {
			sleep()
		}
	}
}

func sleep() {
	// 休眠2秒
	time.Sleep(time.Second * 2)
}

// 拉取数据
func pullData(fc string, listTime time.Time) int {
	fc, filename, ret := stock.KLinePath(fc)
	if ret != stock.D_OK {
		return ret
	}
	cache.CheckFilepath(filename)
	// 读取日线数据
	mapKLine := treemap.NewWithStringComparator()
	fileBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Debugf("code[%s]: K线数据文件不存在", fc)
	} else {
		// 读取日线数据
		cr := codec.NewReader(fileBuf)
		for {
			var kLine Cache.DayKLine
			err := kLine.ReadFrom(cr)
			if err != nil {
				break
			}
			mapKLine.Put(kLine.Date, kLine)
		}
	}
	// 取得当前缓存中最后一个日期
	canUpdateTime := utils.CanUpdateTime()
	// 上市日期
	nextTradingDay := listTime
	// 缓存的日期
	dataLastDay := listTime.AddDate(0, 0, -1)
	// 计算当前日期和最后一个日期相隔的天数
	dataLen := stock.DEFAULT_DATALEN
	dataLen = utils.KLineRequireDays(canUpdateTime, nextTradingDay)
	mkeys := mapKLine.Keys()
	mlen := len(mkeys)
	if mlen < 1 {
		logger.Debugf("code[%s]: 没有缓存数据", fc)
	} else {
		// 有缓存数据, 取得已存在数据的最后一天
		lastKey := mkeys[mlen-1]
		tmpKLine, ok := mapKLine.Get(lastKey)
		if ok {
			lastKLine, ok := tmpKLine.(Cache.DayKLine)
			if ok {
				logger.Debugf("code[%s]: nextTradingDay= %s", fc, lastKLine.Date)
				_nextTradingDay, err := utils.ParseTime(lastKLine.Date)
				if err == nil {
					// 往后顺延一天
					dataLastDay = _nextTradingDay
					_nextTradingDay = utils.NextUpdateTime(_nextTradingDay)
					dataLen = utils.KLineRequireDays(canUpdateTime, _nextTradingDay)
					nextTradingDay = _nextTradingDay
				} else {
					panic("解析日期失败...")
				}
			}
		}
	}
	cacheDate := dataLastDay.Format(util.DateFormat)
	updateDate := canUpdateTime.Format(util.TimeFormat)
	nextDate := nextTradingDay.Format(util.TimeFormat)
	if dataLen < 1 {
		logger.Infof("code[%s@%s]: 数据不需要更新, %s - %s", fc, cacheDate, nextDate, updateDate)
		return stock.D_ERROR | stock.D_ENEED
	}
	logger.Infof("code[%s@%s]: 自[%+v - %+v], 需要补充[%d]天日线数据", fc, cacheDate, nextDate, updateDate, dataLen)
	// 拉取日线数据
	apiTencent := new(tencent.TencentDataApi)
	apiSina := new(sina.SinaDataApi)
	//var dataApi stock.DataApi
	//dataApi = apiSina
	/*ha, _, err := dataApi.DailyFromDate(fc, nextTradingDay)
	if err != nil {
		logger.Debugf("code[%s]: %v", err)
		return stock.D_ERROR | stock.D_ENET
	}
	if len(ha) == 0 {
		logger.Errorf("code[%s]: not data", fc)
		return stock.D_ERROR | stock.D_ENEED
	}*/
	var ha []stock.DayKLine
	var ec int = -1
	if strings.HasPrefix(fc, "hk") {
		ha, ec = fetchData(apiTencent, fc, nextTradingDay)
	} else {
		ha, ec = fetchData(apiSina, fc, nextTradingDay)
		if ec != 0 || len(ha) == 0 {
			logger.Infof("切换tencent api...")
			ha, ec = fetchData(apiTencent, fc, nextTradingDay)
		}
	}
	if ec != 0 {
		return ec
	}

	//对于更细粒度的写入，先打开一个文件。
	//fw, err := os.Create(filename)
	fw, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, category.CACHE_FILE_MODE)
	if err != nil {
		logger.Errorf("code[%s]: 写日线文件失败, error[%+v]", fc, err)
		return stock.D_ERROR | stock.D_EDISK
	}
	count := len(ha)
	wrote := 0
	for j := 0; j < count; j++ {
		kl := ha[j]
		var csk Cache.DayKLine
		csk.Date = kl.Date
		tmpDay, err := utils.ParseTime(csk.Date)
		if err != nil {
			continue
		}
		// 跳过最后日期之前的数据
		if tmpDay.Before(nextTradingDay) {
			continue
		}
		nextTradingDay = tmpDay.AddDate(0, 0, 1)
		api.Copy(&csk, &kl)
		/*csk.Open = kl.Open
		csk.Close = kl.Close
		csk.High = kl.High
		csk.Low = kl.Low
		csk.Volume = kl.Volume*/

		_cos := codec.NewBuffer()
		weer := csk.WriteTo(_cos)
		if weer != nil {
			logger.Debugf("cache output error, %+v", weer)
		}
		fw.Write(_cos.ToBytes())
		wrote += 1
	}
	fw.Sync()
	fw.Close()
	if wrote > 0 {
		logger.Infof("code[%s]: 写日线文件, SUCCESS", fc)
	}

	return 0
}

func fetchData(dataApi stock.DataApi, code string, nextTradingDay time.Time) ([]stock.DayKLine, int) {
	ha, _, err := dataApi.DailyFromDate(code, nextTradingDay)
	if err != nil {
		logger.Debugf("code[%s]: %s, %v", code, dataApi.Name(), err)
		return nil, stock.D_ERROR | stock.D_ENET
	}
	if len(ha) == 0 {
		logger.Errorf("code[%s]: %s, not data", code, dataApi.Name())
		return nil, stock.D_ERROR | stock.D_ENEED
	}
	return ha, 0
}
