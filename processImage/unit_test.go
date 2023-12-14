package processImage

import (
	"log/slog"
	"os"
	"processAll/GetFileInfo"
	"testing"
)

func init() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	slog.SetDefault(logger)
}
func TestOne(t *testing.T) {
	abs := "/Users/zen/Downloads/telegram/cache/Images/4965298626347773683_121.jpg"
	info := GetFileInfo.GetFileInfo(abs)
	ProcessImage(info, "3")
}
func TestAll(t *testing.T) {
	folder := "/Users/zen/Downloads/图片助手(ImageAssistant)_批量图片下载器/girlygirlpic.com"
	ProcessAllImages(folder, "jpg;png", "3")
}
