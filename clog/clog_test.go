package clog

import (
	"fmt"
	"testing"
)

var logLevel int = 2

func TestLogT(t *testing.T) {
	fmt.Println("Error -----------------------")
	logger := NewLogger(LogLevelError, "sdwan")
	logger.Debug("%s", "debug log test")
	logger.Info("%s", "info log test")
	logger.Warn("%s", "warn log test")
	logger.Error("%s", "error log test")

	fmt.Println("Warn -----------------------")
	logger = NewLogger(LogLevelWarn, "sdwan")
	logger.Debug("%s", "debug log test")
	logger.Info("%s", "info log test")
	logger.Warn("%s", "warn log test")
	logger.Error("%s", "error log test")

	fmt.Println("Info -----------------------")
	logger = NewLogger(LogLevelInfo, "sdwan")
	logger.Debug("%s", "debug log test")
	logger.Info("%s", "info log test")
	logger.Warn("%s", "warn log test")
	logger.Error("%s", "error log test")

	fmt.Println("Debug -----------------------")
	logger = NewLogger(LogLevelDebug, "sdwan")
	logger.Debug("%s", "debug log test")
	logger.Info("%s", "info log test")
	logger.Warn("%s", "warn log test")
	logger.Error("%s", "error log test")
}

func BenchmarkLog(b *testing.B) {
	b.ReportAllocs()
	logger := NewLogger(LogLevelError, "sdwan")
	for i := 0; i < b.N; i++ {
		logger.Debug("%s", "debug log test")
		logger.Info("%s", "info log test")
		logger.Warn("%s", "warn log test")
		logger.Error("%s", "error log test")
	}
}

// 执行空函数和执行if判断性能对比
func emptyFunction() {}

func BenchmarkEmptyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptyFunction()
	}
}

func BenchmarkIfCondition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if true {
			// Do nothing
		}
	}
}
