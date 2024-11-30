package cslog

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
)

var (
	logger   *slog.Logger
	once     sync.Once
	logLevel = &slog.LevelVar{}
)

var (
	String = slog.String
	Int    = slog.Int
	Bool   = slog.Bool
	Any    = slog.Any
	Time   = slog.Time
)

// NewLogger create a text or json logger, base log/slog
// 如果没有指定日志目录， 默认: /var/log/程序名/
// 如果没有指定日志文件名，默认: 程序名.log
// 日志格式，默认Json格式； KeyValue文本格式: TERRA_LOG_KV=true，
// 设置日志级别：TERRA_LOG_LEVEL={debug|info|warn|error}
func NewLogger(json bool, logFile string) *slog.Logger {
	if logger != nil {
		return logger
	}
	curLogFile := logFileInit(logFile)
	once.Do(func() {
		var jsonStyle bool = true
		if os.Getenv("TERRA_LOG_KV") == "true" {
			jsonStyle = false
		} else {
			jsonStyle = json
		}
		log := &lumberjack.Logger{
			Filename:   curLogFile,
			MaxSize:    20, // MB
			MaxBackups: 30,
			MaxAge:     28,
			Compress:   true,
			LocalTime:  true,
		}
		opts := &slog.HandlerOptions{
			AddSource: false,
			Level:     logLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					if t, ok := a.Value.Any().(time.Time); ok {
						a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000"))
					}
				}
				return a
			},
		}
		if jsonStyle {
			logger = slog.New(slog.NewJSONHandler(log, opts))
		} else {
			logger = slog.New(slog.NewTextHandler(log, opts))
		}

	})
	updateLogLevel()
	slog.SetDefault(logger)
	return logger
}

func logFileInit(logFile string) string {
	logDir := "/var/log/"
	curLogFile := logFile
	execPath, err := os.Executable()
	if err != nil {
		curLogFile = logDir + "default.log"
	} else {
		execFileName := filepath.Base(execPath)
		logDir += execFileName + "/"
		curLogFile = logDir + execFileName + ".log"
	}

	if logFile != "" {
		dir := filepath.Dir(logFile)
		if dir != "." && dir != "/" {
			logDir = dir
			curLogFile = logFile
		}
	}
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		slog.Warn(fmt.Sprintf("create logDir %s failed", logDir))
		os.Exit(1)
	}
	return curLogFile
}

// 动态更新日志级别
func updateLogLevel() {
	// LevelDebug = -4, LevelInfo = 0, LevelWarn = 4, LevelError = 8
	level := os.Getenv("TERRA_LOG_LEVEL")
	if len(level) < 4 {
		return
	}
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		slog.Warn("update log level failed from OS ENV", Any("error:", err))
		return
	}
	logger.Info("update log level: " + level)
}
