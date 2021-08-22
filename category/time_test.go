package category

import (
	"fmt"
	"github.com/mymmsc/gox/util"
	"testing"
	"time"
)

func Test_realTimeData(t *testing.T) {
	x := time.Date(2017, 02, 27, 17, 30, 20, 20, time.Local)
	fmt.Println(x.Format(util.Timestamp))
	for i := 0; i < 10; i++ {
		x = time.Now()
		fmt.Println(x.Format(util.DateFormat))
		fmt.Println(x.Format(util.TimeFormat))
		fmt.Println(x.Format(util.Timestamp))
		time.Sleep(time.Millisecond * 100)
	}

	tt, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.018273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
		time.Hour * 16,
		time.Hour * 24,
	}

	for _, d := range trunc {
		fmt.Printf("1: t.Truncate(%6s) = %s\n", d, tt.Truncate(d).Format("2006-01-02 15:04:05.999999999"))
		fmt.Printf("2: t.Truncate(%6s) = %s\n", d, tt.Truncate(d).Format("2006-01-02 15:04:05.000000000"))
	}
}
