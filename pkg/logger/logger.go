package logger

import (
	"os"

	"github.com/chrugez/go-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap{
	logLevel :=  config.Log_level //"debug"
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel{
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
		case "warn":
		level = zapcore.WarnLevel
		case "error":
		level = zapcore.ErrorLevel
		default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.File_log_name, //"/var/log/myapp/foo.log"
		MaxSize: config.Max_age, //megabytes
		MaxBackups: config.Max_backups,
		MaxAge: config.Max_age, //days
		Compress: config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder{
	encodeConfig := zap.NewProductionEncoderConfig()

	//1744907387.6061563 -> 2025-04-17T23:29:47.605+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -> Time
	encodeConfig.TimeKey = "time"
	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.log.go:22"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder 
	return zapcore.NewJSONEncoder(encodeConfig)
}