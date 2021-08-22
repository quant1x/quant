package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDifferDays(t *testing.T) {
	layout := "2006-01-02 15:04:05"

	// just one second
	t1, _ := time.Parse(layout, "2007-01-02 23:59:59")
	t2, _ := time.Parse(layout, "2007-01-03 00:00:00")
	if DifferDays(t2, t1) != 1 {
		panic("one second but different day should return 1")
	}

	// just one day
	t1, _ = time.Parse(layout, "2007-01-02 23:59:59")
	t2, _ = time.Parse(layout, "2007-01-03 00:00:01")
	if DifferDays(t2, t1) != 1 {
		panic("just one day should return 1")
	}

	// just one day
	t1, _ = time.Parse(layout, "2007-01-02 23:59:59")
	t2, _ = time.Parse(layout, "2007-01-03 23:59:59")
	if DifferDays(t2, t1) != 1 {
		panic("just one day should return 1")
	}

	t1, _ = time.Parse(layout, "2017-09-01 10:00:00")
	t2, _ = time.Parse(layout, "2017-09-02 11:00:00")
	if DifferDays(t2, t1) != 1 {
		panic("just one day should return 1")
	}

	// more than one day
	t1, _ = time.Parse(layout, "2007-01-02 23:59:59")
	t2, _ = time.Parse(layout, "2007-01-04 00:00:00")
	if DifferDays(t2, t1) != 2 {
		panic("just one day should return 2")
	}
	// just 3 day
	t1, _ = time.Parse(layout, "2007-01-02 00:00:00")
	t2, _ = time.Parse(layout, "2007-01-05 00:00:00")
	if DifferDays(t2, t1) != 3 {
		panic("just 3 day should return 3")
	}

	// different month
	t1, _ = time.Parse(layout, "2007-01-02 00:00:00")
	t2, _ = time.Parse(layout, "2007-02-02 00:00:00")
	if DifferDays(t2, t1) != 31 {
		fmt.Println(DifferDays(t2, t1))
		panic("just one month:31 days should return 31")
	}

	// 29 days in 2mth
	t1, _ = time.Parse(layout, "2000-02-01 00:00:00")
	t2, _ = time.Parse(layout, "2000-03-01 00:00:00")
	if DifferDays(t2, t1) != 29 {
		fmt.Println(DifferDays(t2, t1))
		panic("just one month:29 days should return 29")
	}
	t1 = time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	t2 = time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	if DifferDays(t1, t2) != 1 {
		panic(fmt.Sprintf("just one day: should return 1 but got %v", DifferDays(t1, t2)))
	}

	t1 = time.Date(2018, 1, 10, 0, 0, 1, 100, time.UTC)
	t2 = time.Date(2018, 1, 9, 23, 59, 22, 100, time.UTC)
	if DifferDays(t1, t2) != 1 {
		panic(fmt.Sprintf("just one day: should return 1 but got %v", DifferDays(t1, t2)))
	}
}

func TestParseTime(t *testing.T) {
	var (
		tm time.Time
		err error
	)

	tm, err = ParseTime("900301")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
	tm, err = ParseTime("000301")
	fmt.Printf("time: %+v, error: %v\n", tm, err)

	tm, err = ParseTime("20000301")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
	tm, err = ParseTime("2000-03-02")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
	tm, err = ParseTime("20000303010203")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
	tm, err = ParseTime("2000-03-04 05:06:07")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
	tm, err = ParseTime("2000-03-05 06:07:08.123")
	fmt.Printf("time: %+v, error: %v\n", tm, err)
}

func TestKLineRequireDays(t *testing.T) {
	now, _  := ParseTime("2021-08-09 10:06:07")
	last, _ := ParseTime("2021-08-10 00:00:00")
	n := KLineRequireDays(now, last)
	fmt.Printf("require days = %d", n)
}