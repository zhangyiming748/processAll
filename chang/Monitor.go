package chang

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func RunNumGoroutineMonitor() {
	slog.Info(fmt.Sprintf("程序初始协程数量->%d\n", runtime.NumGoroutine()))
	file, err := os.OpenFile("go.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("协程数量->%d\n", runtime.NumGoroutine())
			file.WriteString(fmt.Sprintf("协程数量->%d\n", runtime.NumGoroutine()))
		}
	}
}
