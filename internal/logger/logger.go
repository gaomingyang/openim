package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func init() {
	//read config

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
			StacktraceKey: "stacktrace",
		},
	}

	//cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//	enc.AppendString(t.Format("Jan 2 2006 3:04:05 pm MST"))
	//}

	lumberLog := lumberjack.Logger{
		//Filename:   viper.GetString("log.outPut"), //不管用
		Filename:   "./log/zap.log",
		MaxSize:    1, //单位M
		MaxBackups: 10,
		MaxAge:     28,
	}

	writeSyncer := zapcore.AddSync(&lumberLog)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	)
	Logger = zap.New(core, zap.AddCaller())
}

// Sync 确保在程序退出之前，所有的日志都被刷新到输出目标。
// 在 zap 中，日志条目可能会被缓冲在内存中，以提高写入性能。如果程序突然退出（例如，通过 os.Exit，panic，或系统崩溃），这些缓冲的日志条目可能没有机会被写入到目标输出中，从而导致日志丢失。调用 Sync 方法可以确保缓冲区中的所有日志条目都被写入到目标输出中。
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
