// logger/log.go

package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	// 设置终端输出为文本格式 json为JSONFormatter
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// // 创建日志文件
	// logFileHook := &lumberjack.Logger{
	// 	Filename:   "logfile.log",
	// 	MaxSize:    10, // MB
	// 	MaxBackups: 3,
	// 	MaxAge:     1, // Days
	// 	LocalTime:  true,
	// }
	//
	// // 设置日志文件输出为JSON格式
	// logFile := logrus.New()
	// logFile.Out = logFileHook
	// logFile.Formatter = &logrus.JSONFormatter{}
	// logFile.Level = logrus.DebugLevel

	// logFile = &logrus.Logger{
	// 	Out:       logFileHook,
	// 	Formatter: &logrus.JSONFormatter{},
	// 	Level:     logrus.DebugLevel,
	// }

	Log.SetLevel(logrus.DebugLevel)

	// 添加终端输出
	Log.SetOutput(os.Stdout)

	// 使用 MultiWriter 将日志输出到终端和文件
	// Log.SetOutput(io.MultiWriter(os.Stdout, logFile.Out))

	// log.Hooks.Add(NewLogFileHook(logFileHook))

	// log.AddHook(NewLogFileHook(logFileHook))

}

func GetLogger() *logrus.Logger {
	return Log
}

func NewLogFileHook(logFileHook *lumberjack.Logger) logrus.Hook {
	lfsHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  logFileHook,
			logrus.ErrorLevel: logFileHook,
		},
		&logrus.JSONFormatter{}, // 日志文件使用json
	)

	return lfsHook
}
