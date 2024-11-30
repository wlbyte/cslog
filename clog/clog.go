package clog

import (
	"log"
	"os"
)

const (
	LogLevelError = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

type Logger struct {
	logLevel int // level: 0(Error),1(Warn),2(Info),3(Debug)
	Error    func(format string, args ...any)
	Warn     func(format string, args ...any)
	Info     func(format string, args ...any)
	Debug    func(format string, args ...any)
}

func EmptyLogf(format string, args ...any) {}

// NewLogger create a logger, base log
// logLevel: 0(Error),1(Warn),2(Info),3(Debug)
// prefix: 日志前缀
func NewLogger(logLevel int, prefix string) *Logger {
	logger := &Logger{logLevel, EmptyLogf, EmptyLogf, EmptyLogf, EmptyLogf}

	logf := func(levelStr string) func(format string, args ...any) {
		return log.New(os.Stdout, prefix+" ["+levelStr+"] ", log.LstdFlags|log.Lshortfile).Printf
	}
	if logger.logLevel >= LogLevelDebug {
		logger.Debug = logf("DEBUG")
	}
	if logger.logLevel >= LogLevelInfo {
		logger.Info = logf("INFO")
	}
	if logger.logLevel >= LogLevelWarn {
		logger.Warn = logf("WARN")
	}
	if logger.logLevel >= LogLevelError {
		logger.Error = logf("ERROR")
	}
	return logger
}

// getLogLevelFromEnv get log level from os by name log_level
func GetLogLevelFromEnv() (level int) {
	levelStr := os.Getenv("log_level")
	switch levelStr {
	case "debug", "Debug":
		level = 3
	case "warn", "warning", "Warn":
		level = 1
	case "error", "Error":
		level = 0
	default:
		level = 2
	}
	return
}
