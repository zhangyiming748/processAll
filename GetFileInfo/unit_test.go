package GetFileInfo

import (
	"fmt"
	"log/slog"
	"os"
	"processAll/chang"
	"processAll/mediaInfo"
	"testing"
)

func init() {
	go chang.RunNumGoroutineMonitor()
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	slog.SetDefault(logger)
}

func TestGetOneInfo(t *testing.T) {
	ret := GetFileInfo("/Users/zen/Pictures/譚詠麟 - 水中花 [6LJ2mJU4BpI].m4a")
	mi := ret.MediaInfo
	if v, ok := mi.(mediaInfo.AudioInfo); ok {
		t.Logf("mi is %+v\n", v)
	} else {
		fmt.Println("不ok")
	}
	t.Logf("%+v\n", ret)
}
func TestGetAllInfo(t *testing.T) {
	ret := GetAllFileInfo("/Users/zen/Pictures", "jpg;dmg;mp4")
	t.Logf("%+v\n", ret)
}

func TestGetAllFilesInfo(t *testing.T) {
	ret := GetAllFilesInfo("/mnt/e/BaiduNetdiskDownload/pikpak", "jpg;dmg;mp4;aac")
	t.Logf("%+v\n", ret)
}
func TestUnit(t *testing.T) {

}

// 四个协程28.463秒
// 十个协程23.26秒
func TestGetAllFilesInfoByChan(t *testing.T) {
	root := "/Users/zen/Downloads"
	ret := GetAllFilesInfoByChan(root, "mp4;dmg;jpg;png;avif")
	file, _ := os.OpenFile("summary.txt", 2|8|512, 0777)
	defer file.Close()
	for _, v := range ret {
		file.WriteString(fmt.Sprintf("%+v\n", v))
	}
	t.Log(ret)
}
func TestIsNotH265(t *testing.T) {
	root := "/Volumes/volume/未整理"
	ret := GetAllFilesInfoByChan(root, "webm;mkv;m4v;mp4;mov;avi;wmv;ts;rmvb;wma;avi;flv;rmvb")
	file, _ := os.OpenFile("report.txt", 2|8|512, 0777)
	for _, one := range ret {
		if mi, ok := one.MediaInfo.(mediaInfo.VideoInfo); ok {
			if mi.VideoFormat != "HEVC" {
				file.WriteString(fmt.Sprintf("不是H265的视频:%s\n", one.FullPath))
			} else if mi.VideoCodecID != "hvc1" {
				file.WriteString(fmt.Sprintf("H265的视频,没有正确的标签:%s\n", one.FullPath))
			}
		}
	}
}

// go test -v -run  TestGetVideoFrameFast ./
// go test -v -run TestGetAllFileInfos ./
// go test -v -run TestGetBitRate ./
func TestGetBitRate(t *testing.T) {

}
