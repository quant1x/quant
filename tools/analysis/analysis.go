package main

import (
	"encoding/json"
	"fmt"
	"github.com/quant1x/quant/utils/http"
	"strconv"
	"strings"
)

/*
* @doc
*
* @字段分析
*    2           1=上证/2=深成    0
*    002547      股票代码         1
*    春兴精工     股票名称         2
*    7.77        当前价格         3
*    0.71        上涨价格         4
*    10.06       涨幅            5
*    2062612     206万手成交量    6
*    1549976928  成交额          7
*    13.31       振幅            8
*    7.77        盘中最高         9
*    6.83        盘中最低         10
*    6.93        开盘价           11
*    7.06        昨收盘价         12
*    0.00        5分钟涨速?       13
*    1.96        量比            14
*    25.86       换手率          15
*    148.87      动态市盈率       16
*    3.19        市净率          17
*    8765004174  总市值
*    6197961277  流通市值
*    81.12%      60日涨幅
*    36.8%       年初至今涨幅
*    0.00        涨速?
*    上市日期
*    最新有效日期
*
 */
func main() {
	url := "http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=CT&token=4f1862fc3b5e77c150a2b985b12db0fd&sty=FCOIATC&cmd=C._A&st=(Code)&sr=1&p=1&ps=10000&_=1555752561456"
	data, err := http.HttpGet(url)
	if err != nil {
		return
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
		return
	}
	datas := ss

	fmt.Println(strings.Repeat("-", 102))
	fmt.Printf(strings.Repeat("|%-11s", 7)+"|\r\n", "名称", "代码", "价格", "涨幅", "成交额", "市盈率", "市净率")
	fmt.Println(strings.Repeat("-", 102))

	count := 0
	for _, data := range datas {
		arr := strings.Split(data, ",")
		if arr[4] != "-" && !strings.HasPrefix(arr[2], "*ST") && !strings.HasPrefix(arr[2], "ST") {
			tmpStrFloat := arr[16]
			tmpV, err := strconv.ParseFloat(tmpStrFloat, 64)
			if err != nil {
				continue
			}
			if tmpV < 5 {
				fmt.Printf("|"+strings.Repeat("%-14s", 7)+"\r\n",
					arr[2], arr[1], arr[3], arr[5], arr[7], arr[16], arr[17])

				count++
			}
		}
	}

	fmt.Println(strings.Repeat("-", 102))
	fmt.Printf("|%-92s|\r\n", fmt.Sprintf("今日可选股票数目: %d", count))
	fmt.Println(strings.Repeat("-", 102))

}
