package logger

import "maoim/pkg/yaml"

var defaultLogger Logger

func InitLogger(configFile string) {
	c := Default()
	err := yaml.DecodeFile(configFile, c)
	if err != nil {
		panic(err)
	}
	defaultLogger = New(c)
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warn(msg string) {
	defaultLogger.Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Panic(msg string) {
	defaultLogger.Panic(msg)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Log(level Level, msg string) {
	defaultLogger.Log(level, msg)

}

func Logf(level Level, format string, args ...interface{}) {
	defaultLogger.Logf(level, format, args...)
}
