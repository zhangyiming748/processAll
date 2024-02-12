package processVideo

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/alert"
	"processAll/mediaInfo"
	"processAll/replace"
	"processAll/util"
	"strings"
)

func ResizeAllVideos(root, pattern, threads string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		ResizeVideo(in, threads)
	}
}
func ResizeVideos(dir, pattern, threads string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		ResizeVideo(in, threads)
	}
}
func ResizeVideo(in GetFileInfo.BasicInfo, threads string) {

	mi, ok := in.MediaInfo.(mediaInfo.VideoInfo)
	if ok {
		slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
	} else {
		slog.Warn("断言视频mediainfo结构体失败")
	}
	if mi.VideoWidth <= 1920 || mi.VideoHeight <= 1920 {
		slog.Debug("跳过", slog.String("正常尺寸的视频", in.FullPath))
		return
	}
	if mi.VideoWidth > mi.VideoHeight {
		slog.Debug("横屏视频", slog.Any("视频信息", in))
		Resize(in, threads, "1920x1080")
	} else if mi.VideoWidth < mi.VideoHeight {
		slog.Debug("竖屏视频", slog.Any("视频信息", in))
		Resize(in, threads, "1080x1920")
	} else {
		slog.Debug("正方形视频", slog.Any("视频信息", in))
		Resize(in, threads, "1920x1920")
	}
}
func Resize(in GetFileInfo.BasicInfo, threads string, p string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("错误", slog.String("文件信息", in.FullPath))
			alert.Customize("failed", alert.Ava)
		}
	}()
	mi, ok := in.MediaInfo.(mediaInfo.VideoInfo)
	if ok {
		slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
	} else {
		slog.Warn("断言视频mediainfo结构体失败")
	}
	dst := in.PurgePath // 文件所在路径 包含最后一个路径分隔符
	if strings.Contains(in.PurgePath, "resize") {
		return
	}
	dst = strings.Join([]string{dst, "resize"}, "") //二级目录
	fname := replace.ForFileName(in.PurgeName)      //仅文件名
	fname = strings.Join([]string{fname, "mp4"}, ".")
	os.Mkdir(dst, 0777)
	slog.Debug("新建文件夹", slog.String("全名", dst))
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("源文件:", in.FullPath), slog.String("输出文件:", out))
	var cmd *exec.Cmd
	switch p {
	case "1920x1080":
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1080", "-c:v", "copy", "-tag:v", "hvc1", "-ac", "1", "-threads", threads, out)
		if mi.VideoCodecID != "hvc1" {
			cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1080", "-c:v", "libx265", "-tag:v", "hvc1", "-ac", "1", "-threads", threads, out)
		}
	case "1080x1920":
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1920", "-c:v", "copy", "-tag:v", "hvc1", "hvc1", "-ac", "1", "-threads", threads, out)
		if mi.VideoCodecID != "hvc1" {
			cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=-1:1920", "-c:v", "libx265", "-tag:v", "hvc1", "-ac", "1", "-threads", threads, out)
		}
	case "1920x1920":
		cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=1920:1920", "-c:v", "copy", "-tag:v", "hvc1", "-ac", "1", "-threads", threads, out)
		if mi.VideoCodecID != "hvc1" {
			cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-strict", "-2", "-vf", "scale=1920:1920", "-c:v", "libx254", "-tag:v", "hvc1", "hvc1", "-ac", "1", "-threads", threads, out)
		}
	default:
		slog.Warn("不正常的视频源", slog.Any("视频信息", in.FullPath))
	}
	slog.Debug("ffmpeg", slog.String("生成的命令", fmt.Sprintf("生成的命令是:%s", cmd)))
	if err := util.ExecCommand(cmd); err != nil {
		slog.Warn("resize发生错误", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", fmt.Sprint(err)), slog.String("源文件", in.FullPath))
		return
	}

	if err := os.Remove(in.FullPath); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误文本", err))
	} else {
		slog.Warn("删除成功", slog.String("源文件", in.FullPath))
	}
}
