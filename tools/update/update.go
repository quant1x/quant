package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"github.com/axgle/mahonia"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/treemap"
	"github.com/quant1x/quant/cache"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/data"
	"github.com/quant1x/quant/data/dfcf"
	"github.com/quant1x/quant/data/security"
	"github.com/quant1x/quant/data/sina"
	"github.com/quant1x/quant/data/tencent"
	"github.com/quant1x/quant/models/Cache"
	"github.com/quant1x/quant/stock"
	"github.com/quant1x/quant/utils"
	"github.com/robfig/cron/v3"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// 更新日线数据工具
func main() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	var (
		path       string // 数据路径
		logPath    string // 日志输出路径
		cronConfig string // 定时脚本
		cronTrue   bool   // 是否启用应用内定时器
	)
	flag.StringVar(&path, "path", category.DATA_ROOT_PATH, "stock history data path")
	flag.StringVar(&logPath, "log_path", category.LOG_ROOT_PATH+"/runtime.log", "log output dir")
	flag.StringVar(&cronConfig, "cron_config", "0 0 17 * * ?", "pull code data cron")
	flag.BoolVar(&cronTrue, "cron_true", false, "use crontab in application")
	flag.Parse()
	utils.InitLog(logPath, 500, 5, 5)
	logger.Info("data path: ", path, ",logPath:", logPath, ",cronConfig:", cronConfig)
	cache.Init(path)
	if !cronTrue {
		handleCodeData()
	} else {
		crontab := cron.New(cron.WithSeconds()) //精确到秒
		// 添加定时任务,
		crontab.AddFunc(cronConfig, handleCodeData)
		//启动定时器
		crontab.Start()
		select {
		case sig := <-c:
			{
				logger.Info("进程结束:", sig)
				os.Exit(1)
			}
		}
	}
}

func handleCodeData() {
	logger.Info("任务开始启动...")
	fullCodes := data.GetCodeList()
	updateSpe("https://www.hkex.com.hk/-/media/HKEX-Market/Mutual-Market/Stock-Connect/Eligible-Stocks/View-All-Eligible-Securities/SZSE_Securities_c.csv?la=zh-HK", "sz")
	updateSpe("https://www.hkex.com.hk/-/media/HKEX-Market/Mutual-Market/Stock-Connect/Eligible-Stocks/View-All-Eligible-Securities/SSE_Securities_c.csv?la=zh-HK", "sh")

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
	logger.Info("任务执行完毕.", time.Now())
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
	//构建目录
	cache.CheckFilepath(filename)
	// 读取本地日线数据
	var mapKLine *treemap.Map
	fcNotExist := false
	if fr, err := os.Open(filename); err != nil {
		//ENOENT (2)
		if errors.Is(err, syscall.ENOENT) {
			logger.Debugf("code[%s]: K线数据文件不存在", fc)
			fcNotExist = true
			mapKLine = treemap.NewWithStringComparator()
		} else {
			logger.Errorf("code[%s]: K线数据文件操作失败:%v", fc, err)
			return 0
		}
	} else {
		var kLine Cache.DayKLine
		mapKLine, err = kLine.ReadMapFromCsv(csv.NewReader(fr))
		if err != nil {
			logger.Errorf("code[%s]: K线数据文件读取失败:%v", fc, err)
			return 0
		}
		fcNotExist = mapKLine.Empty()
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
	apiDfcf := new(dfcf.EastmoneyApi)
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
		ha, ec = fetchData(apiDfcf, fc, nextTradingDay)
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
	count := len(ha)
	wrote := 0
	fw, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, category.CACHE_FILE_MODE)
	_writer := csv.NewWriter(fw)
	if fcNotExist && count > 0 {
		var cskHead Cache.DayKLine
		cskHead.Init(_writer)
	}
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

		csk.WriteCSV(_writer)
		wrote += 1
	}
	_writer.Flush()
	//fw.Sync()
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

func ChangeDecode(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 深股通特别证券/中华通特别证券名单（只可卖出）https://www.hkex.com.hk/-/media/HKEX-Market/Mutual-Market/Stock-Connect/Eligible-Stocks/View-All-Eligible-Securities/SZSE_Securities_c.csv?la=zh-HK
// 港股通特别证券/中华通特别证卷名单 (只可卖出) https://www.hkex.com.hk/-/media/HKEX-Market/Mutual-Market/Stock-Connect/Eligible-Stocks/View-All-Eligible-Securities/SSE_Securities_c.csv?la=zh-HK
func updateSpe(url string, market string) {
	//抓取HKEX csv文件
	speMap := fetchHKEX(url)
	//加载market到数组
	list, err := security.GetStaticBasic(market)
	if err != nil {
		logger.Errorf(market, "个股信息加载失败")
	} else {
		//遍历market数组 比对HKEX csv文件是否存在
		for k, item := range list {
			cCass, ok := speMap[item.Security.Code]
			if ok {
				list[k].CCass = cCass
				list[k].ASecurity = true
			} else {
				list[k].ASecurity = false
			}
		}
		u, _ := json.Marshal(list)
		//写入文件 json
		security.WriteBasicInfo(market, u)
	}

}

func fetchHKEX(url string) map[string]string {
	res, _ := http.Get(url)
	defer res.Body.Close()

	buffer, _ := ioutil.ReadAll(res.Body)
	src := string(buffer)

	reDecodeSrc := ChangeDecode(src, "utf16", "utf8")
	enterSrc := strings.Split(reDecodeSrc, "\n")

	var k, j, start int
	var startPoint bool
	startPoint = true
	for start = 0; start < len(enterSrc); start++ {
		rowSrc := strings.Split(enterSrc[start], "\t")
		if startPoint {
			//从序号为1的开始
			if strings.Compare(strings.TrimSpace(rowSrc[0]), "1") == 0 {
				startPoint = false
				break
			}
		}
	}
	j = 1

	szMap := make(map[string]string)

	for k = start; k < len(enterSrc); k++ {
		rowSrc := strings.Split(enterSrc[k], "\t")
		if len(rowSrc) > 1 {
			szMap[strings.TrimSpace(rowSrc[1])] = strings.TrimSpace(rowSrc[2])
		}
		j++
	}
	return szMap
}
