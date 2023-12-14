package processVideo

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/replace"
	"processAll/util"
	"strconv"
	"strings"
)

/*
执行文件夹和子文件夹中音视频加速
*/
func SpeedUpAllVideos(root, pattern string, speed string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		SpeedupVideo(in, speed)
	}
}

/*
执行一个文件夹中音视频加速
*/
func SpeedUpVideos(dir, pattern string, speed string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		SpeedupVideo(in, speed)
	}
}

/*
加速单个音视频的完整函数
*/
func SpeedupVideo(in GetFileInfo.BasicInfo, speed string) {
	dst := strings.Join([]string{in.PurgePath, "speed"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	fname := replace.ForFileName(in.PurgeName)
	fname = strings.Join([]string{fname, "mp4"}, ".")
	slog.Debug("补全后的 fname", slog.String("fname", fname))
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
	//跳过已经加速的文件夹
	if strings.Contains(in.FullPath, "speed") {
		return
	}
	speedUp(in.FullPath, out, speed)
}

/*
仅使用输入输出和加速参数执行命令
*/
func speedUp(in, out string, speed string) {
	//ffmpeg -i 5_6253787118179453662.mp4 -y -vf "setpts=0.8*PTS" -filter:a "atempo=1.25" -c:v libx265 -c:a aac -ac 1 -tag:v hvc1 6253787118179453662.mp4
	ff := audio2video(speed)
	pts := strings.Join([]string{"setpts=", ff, "*PTS"}, "")
	atempo := strings.Join([]string{"atempo", ff}, "=")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", atempo, "-vf", pts, "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", out)
	util.ExecCommand(cmd)
	if err := os.RemoveAll(in); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in), slog.Any("错误内容", err))
	} else {
		slog.Debug("删除成功", slog.String("源文件", in))
	}
}

/*
获取一个等效ffmpeg音视频参数
*/
func audio2video(speed string) string {
	audio, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		slog.Warn("解析加速参数错误,退出程序", slog.String("错误原文", fmt.Sprint(err)))
		os.Exit(1)
	}
	video := 1 / audio
	slog.Debug("转换后的原始参数", slog.Float64("Video", video))
	final := fmt.Sprintf("%.2f", video)
	slog.Debug("保留两位小数的原始参数", slog.String("final", final))
	return final
}
