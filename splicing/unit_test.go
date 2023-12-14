package splicing

import (
	"io"
	"log/slog"
	"os"
	"testing"
)

func init() {

	var opt slog.HandlerOptions

	opt = slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}

	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

func TestSpicingAAC(t *testing.T) {
	out := "扶她少女自慰日记9.aac"
	SpicingAAC(out)
}
