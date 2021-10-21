package logger

import (
	"errors"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerInstance *zap.Logger

type Config struct {
	Path       string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

//InitZapLogger initialize zap logger
func InitZapLogger(cfg Config) error {

	//check if config is not null pointer
	if cfg.Path == "" || cfg.Filename == "" || cfg.MaxSize <= 0 {
		return errors.New("provide non-nil config")
	}

	writerSyncer := getLogWriter(cfg)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	loggerInstance = logger

	return nil
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(cfg Config) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Path + "/" + cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(args ...interface{}) {
	loggerInstance.Sugar().Info(args...)
}
func Infof(template string, args ...interface{}) {
	loggerInstance.Sugar().Infof(template, args...)
}
func Warn(args ...interface{}) {
	loggerInstance.Sugar().Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	loggerInstance.Sugar().Warnf(template, args...)
}
func Fatal(args ...interface{}) {
	loggerInstance.Sugar().Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	loggerInstance.Sugar().Fatalf(template, args...)
}
