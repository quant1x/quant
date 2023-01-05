package data

import (
	"encoding/json"
	"fmt"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"github.com/quant1x/quant/category"
	"github.com/quant1x/quant/http"
	"github.com/quant1x/quant/stock"
	"strings"
)

func init() {
	logger.SetLogPath(category.LOG_ROOT_PATH)
}

func GetCodeList() []string {
	fullCodes := make([]string, 0)
	// 指数
	indexes := []string{"sh000001",
		"sh000905", "sz399001", "sz399006"}
	fullCodes = append(fullCodes, indexes...)

	// 更新代码
	// 上海
	// sh600000-sh600999
	{
		var (
			codeBegin = 600000
			codeEnd   = 600999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sh%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// sh601000-sh601999
	{
		var (
			codeBegin = 601000
			codeEnd   = 601999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sh%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// sh603000-sh603999
	{
		var (
			codeBegin = 603000
			codeEnd   = 603999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sh%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// sh688000-sh688999
	{
		var (
			codeBegin = 688000
			codeEnd   = 688999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sh%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// 深圳证券交易所
	// 深圳主板: sz000000-sz000999
	{
		var (
			codeBegin = 0
			codeEnd   = 999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sz000%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// 中小板: sz002000-sz002999
	{
		var (
			codeBegin = 2000
			codeEnd   = 2999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sz00%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	// 创业板: sz300000-sz300999
	{
		var (
			codeBegin = 300000
			codeEnd   = 300999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("sz%d", i)
			fullCodes = append(fullCodes, fc)
		}
	}
	fullCodes = fullCodes[0:0]
	// 港股: hk00001-hk09999
	{
		var (
			codeBegin = 1
			codeEnd   = 9999
		)
		for i := codeBegin; i <= codeEnd; i++ {
			fc := fmt.Sprintf("hk%05d", i)
			fullCodes = append(fullCodes, fc)
		}
	}

	return fullCodes
}

func getMarket0() []string {
	url := "http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=CT&token=4f1862fc3b5e77c150a2b985b12db0fd&sty=FCOIATC&cmd=C._A&st=(Code)&sr=1&p=1&ps=10000&_=1555752561456"
	url = "http://push2.eastmoney.com/api/qt/clist/get?fid=f12&po=0&pz=50000&pn=1&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5&fs=m:0+t:6+f:!2,m:0+t:13+f:!2,m:0+t:80+f:!2,m:1+t:2+f:!2,m:1+t:23+f:!2,m:0+t:7+f:!2,m:1+t:3+f:!2&fields=f12,f14"
	data, err := http.HttpGet(url)
	if err != nil {
		return nil
	}
	if data[0] == '(' {
		data = data[1:]
	}
	blen := len(data)
	if data[blen-1] == ')' {
		data = data[:blen-1]
	}

	var ss []string
	err = json.Unmarshal(data, &ss)
	if err != nil {
		return nil
	}
	return ss
}

func getCodeList_old() []stock.StockInfo {
	datas := getMarket0()
	var ss []stock.StockInfo
	for _, data := range datas {
		arr := strings.Split(data, ",")
		var si stock.StockInfo
		err := api.Convert(arr, &si)
		if err == nil {
			if strings.EqualFold(si.Market, "1") {
				si.FullCode = "sh"
			} else {
				si.FullCode = "sz"
			}
			si.FullCode = si.FullCode + si.Code
			ss = append(ss, si)
		}
	}
	return ss
}
