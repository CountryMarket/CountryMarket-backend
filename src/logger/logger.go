package logger

import (
	"fmt"
	"github.com/CountryMarket/CountryMarket-backend/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"os"
	"path"
	"time"
)

var LogrusLogger *logrus.Logger
var SlowLogger logger.Interface

type GormWriter struct {
	glog *logrus.Logger
}

//实现gorm/logger.Writer接口
func (m *GormWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.glog.Info(logstr)
}

func init() {
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/" + config.C.LogConf.LogPath + "/"
	}
	// 创建日志文件夹
	filenameFormat := time.Now().Format("2006-01-02 15:04:05")
	src := createLogFile(logFilePath, filenameFormat)

	LogrusLogger = logrus.New()
	LogrusLogger.SetLevel(logrus.DebugLevel)
	LogrusLogger.Out = src

	LogrusLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// gorm
	SlowLogger = logger.New(
		//设置Logger
		&GormWriter{glog: LogrusLogger},
		logger.Config{
			//慢SQL阈值
			SlowThreshold: time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Warn,
		},
	)
}

func createLogFile(logFilePath, filenameFormat string) *os.File {
	logFileName := config.C.LogConf.LogFileName + filenameFormat + ".log"
	fileName := path.Join(logFilePath, logFileName)

	// 检查是否能够成功创建日志文件
	checkFile := func(filename string) {
		// 以时间去命名日志文件
		// 先去判断文件名字是否合法
		if _, err := os.Stat(filename); err != nil {
			if _, err := os.Create(filename); err != nil {
				fmt.Println("create file failed")
				panic(err)
			}
		}
	}
	checkFile(fileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return src
}
