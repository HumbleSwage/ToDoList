package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		LogrusObj.Out = src
		return
	}
	// 日志对象的实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	filename := path.Join(logFilePath, logFileName)
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		if _, err := os.Create(filename); err != nil {
			return nil, err
		}
	}
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
