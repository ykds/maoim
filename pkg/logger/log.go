package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Level int8

const (
	INFO Level = iota + 1
	WARN
	ERROR
	PANIC
	FATAL
)

type Logger interface {
	Info(msg string)
	Infof(format string, args ...interface{})

	Warn(msg string)
	Warnf(format string, args ...interface{})

	Error(msg string)
	Errorf(format string, args ...interface{})

	Panic(msg string)
	Panicf(format string, args ...interface{})

	Fatal(msg string)
	Fatalf(format string, args ...interface{})

	Log(level Level, msg string)
	Logf(level Level, format string, args ...interface{})
}

type logger struct {
	c   *Config
	log *zap.Logger
}

func New(c *Config) Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	hook := lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		LocalTime:  true,
		Compress:   c.Compress,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	log := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(log)
	return &logger{
		c:   c,
		log: log,
	}
}

func (l logger) Info(msg string) {
	l.Log(INFO, msg)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Logf(INFO, format, args...)
}

func (l logger) Warn(msg string) {
	l.Log(WARN, msg)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.Logf(WARN, format, args...)
}

func (l logger) Error(msg string) {
	l.Log(ERROR, msg)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Logf(ERROR, format, args...)
}

func (l logger) Panic(msg string) {
	l.Log(PANIC, msg)
}

func (l logger) Panicf(format string, args ...interface{}) {
	l.Logf(PANIC, format, args...)
}

func (l logger) Fatal(msg string) {
	l.Log(FATAL, msg)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.Logf(FATAL, format, args...)
}

func (l logger) Log(level Level, msg string) {
	switch level {
	case INFO:
		l.log.Info(msg)
	case WARN:
		l.log.Warn(msg)
	case ERROR:
		l.log.Error(msg)
	case PANIC:
		l.log.Panic(msg)
	case FATAL:
		l.log.Fatal(msg)
	}
}

func (l logger) Logf(level Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Log(level, msg)
}
