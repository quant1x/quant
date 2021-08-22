package cache

import (
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/quant1x/quant/models/Cache"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/treemap"
	"io/ioutil"
	"strings"
)

// LoadKLine 加载日线文件
func LoadKLine(fc string) []Cache.DayKLine {
	fc = strings.TrimSpace(fc)
	if len(fc) != 7 && len(fc) != 8 {
		return nil
	}
	pos := len(fc) - 3
	fc = strings.ToLower(fc)
	// 组织存储路径
	filename := GetDayPath() + "/" + fc[0:pos] + "/" + fc
	CheckFilepath(filename)
	// 读取日线数据
	mapKLine := treemap.NewWithStringComparator()
	fileBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Debugf("code[%s]: K线数据文件不存在", fc)
		return nil
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
		var kls []Cache.DayKLine
		for _, v := range mapKLine.Values() {
			kl, ok := v.(Cache.DayKLine)
			if ok {
				kls = append(kls, kl)
			}
		}
		return kls
	}
}
