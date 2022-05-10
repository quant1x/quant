package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(logPath string, maxSize int, maxBackups int, maxAge int) {
	lf := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxSize,    // 日志文件大小，单位是 MB
		MaxBackups: maxBackups, // 最大过期日志保留个数
		MaxAge:     maxAge,     // 保留过期文件最大时间，单位 天
		Compress:   true,       // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}
	log.SetOutput(lf) // logrus 设置日志的输出方式
}
