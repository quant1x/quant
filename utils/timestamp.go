package utils

import (
	"github.com/mymmsc/gox/errors"
	"github.com/mymmsc/gox/util"
	"github.com/quant1x/quant/category"
	"strings"
	"time"
)

// UnixTime 毫秒数转time.Time
func UnixTime(timestamp int64) time.Time  {
	return time.Unix(timestamp, 0)
}

// ParseTime 解析时间差
func ParseTime(timestr string) (time.Time, error) {
	s := strings.TrimSpace(timestr)
	switch len(s) {
	case len(util.DateFormat):
		return time.ParseInLocation(util.DateFormat, s, time.Local)
	case len(util.DateFormat2):
		return time.ParseInLocation(util.DateFormat2, s, time.Local)
	case len(util.DateFormat3):
		return time.ParseInLocation(util.DateFormat3, s, time.Local)
	case len(util.TimeFormat2):
		return time.ParseInLocation(util.TimeFormat2, s, time.Local)
	case len(util.TimeFormat):
		return time.ParseInLocation(util.TimeFormat, s, time.Local)
	case len(util.Timestamp):
		return time.ParseInLocation(util.Timestamp, s, time.Local)
	default:
		return time.Time{}, errors.New("日期格式无法确定")
	}
}

// DifferDays 计算天数差
// 从 t1 回到 t2 需要多少天
func DifferDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// DateZero t 的0点0分0秒
func DateZero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// t 的16点整
// 参照  /category/time.go
func updateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), category.HistoryUpdateHour, category.HistoryUpdateMin, category.HistoryUpdateSec, 0, time.Local)
}

// 是否工作日
func isWorkday(t time.Time) bool {
	weekDay := t.Weekday()
	if weekDay == time.Sunday || weekDay == time.Saturday {
		return false
	}
	return true
}

// CanUpdate t 时间可以更新的时间, 调整到对应时分秒
func CanUpdate(t time.Time) bool {
	now := time.Now()
	if !isWorkday(t) {
		return false
	}
	ut := updateTime(t)
	return now.After(ut)
}

// CanUpdateTime 是否可以更新
func CanUpdateTime() time.Time {
	t := time.Now()
	for {
		if CanUpdate(t) {
			break
		} else {
			t = updateTime(t)
			t = t.AddDate(0, 0, -1)
		}
	}
	return t
}

// NextUpdateTime 下一个可以更新的日期
func NextUpdateTime(t time.Time) time.Time {
	for {
		// 日期顺延一天
		t = t.AddDate(0, 0, 1)
		// 如果不能更新
		if isWorkday(t) {
			break
		}
	}
	return t
}

func KLineRequireDays(currentDate, lastDay time.Time) int {
	if currentDate.Before(lastDay) {
		return 0
	}
	offset := 0
	//t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	//t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	hours := int(currentDate.Sub(lastDay).Hours())
	if CanUpdate(currentDate) && (hours % 24) != 0 {
		offset = 1
	}
	n := int(hours / 24)
	return n + offset
}