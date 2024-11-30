package main

import (
	"os"
	"sync"
	"time"

	"github.com/wlbyte/go-project/2-stdlib/log/cslog"
)

func main() {
	var wg sync.WaitGroup
	logger := cslog.NewLogger(true, "")
	wg.Add(1)
	go func() {
		for i := 0; i < 60; i++ {
			logger.Debug("log test", "user", os.Getenv("USER"), cslog.Any("debug", "test"))
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 60; i++ {
			logger.Info("log test", "user", os.Getenv("USER"), cslog.Any("test", 1234))
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
