package cache

import (
	"github.com/quant1x/quant/category"
	"os"
	"path/filepath"
)

const (
	INFO_PATH = "info"
	DAY_PATH  = "day"
)

var (
	// CACHE_ROOT_PATH cache路径
	CACHE_ROOT_PATH = category.DATA_ROOT_PATH
	// 是否启用CSV格式
	CACHE_DATA_CSV = false
)

// CheckFilepath
// 检查filename 文件路径, 如果不存在就创建
func CheckFilepath(filename string) error {
	dir := filepath.Dir(filename)
	return os.MkdirAll(dir, category.CACHE_DIR_MODE)
}

// Init
// 初始化缓存路径
func Init(path string, useCSV bool) error {
	err := os.MkdirAll(path, category.CACHE_DIR_MODE)
	if err != nil {
		return err
	}
	CACHE_ROOT_PATH = path
	CACHE_DATA_CSV = useCSV
	return nil
}

// GetInfoPath 证券信息路径
func GetInfoPath() string {
	return CACHE_ROOT_PATH + "/" + INFO_PATH
}

// GetDayPath 历史数据-日线缓存路径
func GetDayPath() string {
	return CACHE_ROOT_PATH + "/" + DAY_PATH
}
