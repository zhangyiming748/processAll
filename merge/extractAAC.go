package merge

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/mediaInfo"
	"processAll/replace"
	"processAll/util"
	"strings"
)

func ExtractAAC(rootPath string) {
	roots := getall(rootPath)
	slog.Debug("根目录", slog.Any("roots", roots))
	for _, root := range roots {
		slog.Info("1", slog.String("1", root))
		secs := getall(root)
		for _, sec := range secs {
			slog.Info("2", slog.String("2", sec))
			entry := strings.Join([]string{sec, "entry.json"}, string(os.PathSeparator))
			name := getName(entry)
			name = replace.ForFileName(name)
			name = CutName(name)
			slog.Info("entry", slog.String("获取到的文件名", name))
			thirds := getall(sec)
			for _, third := range thirds {
				slog.Info("3", slog.String("3", third))
				video := strings.Join([]string{third, "video.m4s"}, string(os.PathSeparator))
				audio := strings.Join([]string{third, "audio.m4s"}, string(os.PathSeparator))
				fname := strings.Join([]string{name, "aac"}, ".")
				if isExist("/sdcard/Movies") {
					os.Mkdir("/sdcard/Movies/bili", 0777)
					fname = strings.Join([]string{"/sdcard/Movies/bili", fname}, string(os.PathSeparator))
				} else {
					fname = strings.Join([]string{rootPath, fname}, string(os.PathSeparator))
				}
				slog.Info("最终名称", slog.String("文件名", fname), slog.String("音频", audio))
				vInfo := GetFileInfo.GetFileInfo(video)
				mi, ok := vInfo.MediaInfo.(mediaInfo.VideoInfo)
				if ok {
					slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
				} else {
					slog.Warn("断言视频mediainfo结构体失败")
				}
				slog.Info("WARNING", slog.String("vTAG", mi.VideoCodecID))
				cmd := exec.Command("ffmpeg", "-i", audio, "-c:a", "copy", "-ac", "1", fname)
				if mi.VideoCodecID == "avc1" {
					cmd = exec.Command("ffmpeg", "-i", video, "-i", audio, "-c:v", "copy", "-c:a", "copy", "-ac", "1", fname)
				}
				err := util.ExecCommand(cmd)
				if err != nil {
					slog.Warn("哔哩哔哩合成出错", slog.Any("错误原文", err), slog.Any("命令原文", fmt.Sprint(cmd)))
					continue
				}
			}
		}
	}
}
