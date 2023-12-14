package processVideo

import (
	"fmt"
	"log/slog"
	"os"
	"processAll/GetFileInfo"
	"processAll/mediaInfo"
	"runtime"
)

func GetOutOfH265(root, pattern string) {
	report, err := os.OpenFile("report.md", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	h264, err := os.OpenFile("h264.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	h265, err := os.OpenFile("h265.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer report.Close()
	report.WriteString(fmt.Sprintf("|文件名|h264|未打标签的h265|\n"))
	report.WriteString(fmt.Sprintf("|:---:|:---:|:---:|\n"))

	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	length := len(infos)
	slog.Info("", slog.Int("总文件数", length))
	for i, in := range infos {
		fmt.Printf("正在处理第个%d /%d文件\n", i+1, length)
		mi, ok := in.MediaInfo.(mediaInfo.VideoInfo)
		if ok {
			slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
		} else {
			slog.Warn("断言视频mediainfo结构体失败")
		}
		if mi.VideoCodecID == "hvc1" {
			// 一定是h265视频
		} else if mi.VideoFormat == "HEVC" {
			// 一定是没有打标签的h265视频
			fmt.Printf("%s\t是h265视频但没有打标签\n", in.FullPath)
			report.WriteString(fmt.Sprintf("|%v||\u2713|\n", in.FullPath))
			h265.WriteString(fmt.Sprintf("%s\n", in.FullPath))
		} else {
			// 一定不是h265视频
			fmt.Printf("%s\t不是h265视频\n", in.FullPath)
			report.WriteString(fmt.Sprintf("|%v|\u2713||\n", in.FullPath))
			h264.WriteString(fmt.Sprintf("%s\n", in.FullPath))
		}
		report.Sync()
		runtime.GC()

	}
}
