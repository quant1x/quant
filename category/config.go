package category

import (
	"github.com/mymmsc/gox/logger"
	"os"
)

const (
	// DATA_ROOT_PATH 数据根路径
	DATA_ROOT_PATH = "/opt/data/ctp"
	// KLINE_PATH 日线数据文件路径
	KLINE_PATH = DATA_ROOT_PATH + "/day"
	// CACHE_DIR_MODE 目录权限
	CACHE_DIR_MODE os.FileMode = 0755
	// CACHE_FILE_MODE 文件权限
	CACHE_FILE_MODE os.FileMode = 0644

	// DEBUG 调试开关
	DEBUG = false

	// LOG_ROOT_PATH 日志路径
	LOG_ROOT_PATH = "/opt/logs/ctp"
)

func init() {
	// 创建目录
	if err := os.MkdirAll(KLINE_PATH, CACHE_DIR_MODE); err != nil {
		panic(err)
	}
	logger.SetLogPath(LOG_ROOT_PATH)
}
