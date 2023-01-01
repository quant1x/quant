package cache

import (
	"encoding/csv"
	"errors"
	"github.com/mymmsc/gox/logger"
	"github.com/quant1x/quant/models/Cache"
	"os"
	"strings"
	"syscall"
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
	if CACHE_TYPE == CACHE_CSV {
		filename += ".csv"
	}

	CheckFilepath(filename)
	if fr, err := os.Open(filename); err != nil {
		//ENOENT (2)
		if errors.Is(err, syscall.ENOENT) {
			logger.Debugf("code[%s]: K线数据文件不存在", fc)
			return nil
		} else {
			logger.Errorf("code[%s]: K线数据文件操作失败:%v", fc, err)
			return nil
		}
	} else {
		var kLine Cache.DayKLine
		// 读取日线数据
		mapKLine, err := kLine.ReadMapFromCsv(csv.NewReader(fr))
		if err != nil {
			logger.Errorf("code[%s]: K线数据文件读取失败:%v", fc, err)
			return nil
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
