package tencent

import (
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

type Pair struct {
	Name  string
	Value string
}

type Query struct {
	Key   string
	Pairs [49]*Pair
}

var (
	invalidResponse = errors.New("invalid result")
	names           = map[int]string{
		0:  "交易所",
		1:  "名字",
		2:  "代码",
		3:  "当前价格",
		4:  "昨收",
		5:  "今开",
		6:  "成交量（手)",
		7:  "外盘",
		8:  "内盘",
		9:  "买一",
		10: "买一量（手）",
		11: "买二",
		12: "买二量（手）",
		13: "买三",
		14: "买三量（手）",
		15: "买四",
		16: "买四量（手）",
		17: "买五",
		18: "买五量（手）",
		19: "卖一",
		20: "卖一量",
		21: "卖二",
		22: "卖二量",
		23: "卖三",
		24: "卖三量",
		25: "卖四",
		26: "卖四量",
		27: "卖五",
		28: "卖五量",
		29: "最近逐笔成交",
		30: "时间",
		31: "涨跌",
		32: "涨跌%",
		33: "最高",
		34: "最低",
		35: "价格/成交量（手）/成交额",
		36: "成交量（手）",
		37: "成交额（万）",
		38: "换手率",
		39: "市盈率",
		40: "未知",
		41: "最高",
		42: "最低",
		43: "振幅",
		44: "流通市值",
		45: "总市值",
		46: "市净率",
		47: "涨停价",
		48: "跌停价",
		49: "量比",
		50: "未知",
		51: "均价",
		52: "动态市盈率",
		53: "静态市盈率",
		54: "-",
		55: "-",
		56: "-",
		57: "成交额",
	}
)

func queryAll() ([]*Query, error) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(runtime.NumCPU()))

	qf := func(key string) {
		if _, e := query(key); e != nil {
			log.Printf("query %s failed: %s", key, e)
		} else {
			log.Printf("query %s done", key)
		}
	}
	pool := make(chan struct{}, runtime.NumCPU())
	waiter := sync.WaitGroup{}
	const count = 1000
	qs := make([]*Query, 0, count)
	for i := 1; i < count; i++ {
		key := fmt.Sprintf("sz%06d", i)
		select {
		case pool <- struct{}{}:
			waiter.Add(1)
			go func(key string) {
				qf(key)
				<-pool
				waiter.Done()
			}(key)
		default:
			qf(key)
		}
	}
	waiter.Wait()
	return qs, nil
}

func query(key string) (*Query, error) {
	res, err := http.Get(url + key)
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
	if len(pairs) < len(names) {
		return nil, invalidResponse
	}
	q := &Query{Key: key}
	for i := range q.Pairs {
		q.Pairs[i] = &Pair{Name: names[i], Value: pairs[i]}
	}
	return q, nil
}
