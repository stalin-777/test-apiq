package logger

import (
	"errors"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerInstance *zap.Logger

//InitZapLogger initialize zap logger
func InitZapLogger(directory string, filename string, maxSize int, maxBackups int, maxAge int) error {

	//check if config is not null pointer
	if directory == "" || filename == "" || maxSize <= 0 {
		return errors.New("provide non-nil config")
	}

	writerSyncer := getLogWriter(directory, filename, maxSize, maxBackups, maxAge)
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

func getLogWriter(directory string, filename string, maxSize int, maxBackups int, maxAge int) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   directory + "/" + filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
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
