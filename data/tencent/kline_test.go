package tencent

import (
	"encoding/json"
	"fmt"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util"
	"github.com/quant1x/quant/stock"
	"testing"
	"time"
)

func TestCreateUrl(t *testing.T) {
	fmt.Println(createUrl("sh000001", stock.DEFAULT_DATALEN))
	fmt.Println(createUrl("sh000001", stock.DEFAULT_DATALEN))
	fmt.Println(createUrl("sh000001", stock.DEFAULT_DATALEN))
}

func TestHistory(t *testing.T) {
	kl := historyByDays("sh000001", 1)
	fmt.Printf("%+v\n", kl)

	kl = historyByDays("hk00700", 10000000)
	fmt.Printf("2. %+v\n", kl)
}

func TestTencentDataApi_HongKong(t *testing.T) {
	var is []int
	is = append(is, 1)
	var bs []int
	for _, item := range bs {
		fmt.Println("item = ", item)
	}
	is = append(is, bs...)
	defer logger.FlushLogger()
	api := TencentDataApi{}
	kl, err := api.CompleteKLine("hk00700")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(kl)
	fmt.Printf("%s\n", string(buf))
	t.Logf("%d", len(kl))
}

func TestTencentDataApi_CompleteKLine(t *testing.T) {
	var is []int
	is = append(is, 1)
	var bs []int
	for _, item := range bs {
		fmt.Println("item = ", item)
	}
	is = append(is, bs...)
	defer logger.FlushLogger()
	api := TencentDataApi{}
	kl, err := api.CompleteKLine("sh000001")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(kl)
	fmt.Printf("%s\n", string(buf))
	t.Logf("%d", len(kl))
}

func TestTencentDataApi_DailyFromDate(t *testing.T) {
	defer logger.FlushLogger()
	api := TencentDataApi{}
	startTime, _ := time.Parse(util.DateFormat, "2021-08-06")
	kl, dataLastDay, err := api.DailyFromDate("sh600600", startTime)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(kl)
	fmt.Printf("%+v, %s\n", dataLastDay, string(buf))
	t.Logf("%d", len(kl))
}

func TestCalculateRemainingDays(t *testing.T) {
	const (
		kYears = 21
		kDays  = 193
	)
	t1 := time.Date(2000, 7, 13, 9, 30, 0, 0, time.Local)
	//t1  = time.Date(2021, 7, 12, 9, 30, 0, 0, time.Local)
	t2 := time.Date(2021, 7, 13, 20, 0, 0, 0, time.Local)

	years, days := calculateRemainingDays(t2, t1)
	if years != kYears || days != kDays {
		t.Fatalf("计算错误, 年数[%d]=%d, 天数[%d]=%d", kYears, years, kDays, days)
	}
}
