package sina

import (
	"github.com/mymmsc/gox/api"
	"strings"
)

const (
	indexPatternString = "var hq_str_s_(\\w{7,8})=\"(.+)\"";
	stockPatterString = "var hq_str_(\\w{7,8})=\"(.+)\"";
)

func GetRealtime(code string) {

}

func GetHushenRealtime(code string, data []string) (*SinaHushenRealtime, error) {
	var rt SinaHushenRealtime

	err := api.Convert(data, &rt)
	if err != nil {
		return nil, err
	}
	rt.Code = code
	return &rt, nil
}

func GetHongkongRealtime(code string, data []string) (*SinaHonkongRealtime, error) {
	var rt SinaHonkongRealtime

	err := api.Convert(data, &rt)
	if err != nil {
		return nil, err
	}
	rt.Code = code
	rt.Date = strings.ReplaceAll(rt.Date, "/", "-")
	rt.Time = rt.Time + ":59"
	return &rt, nil
}