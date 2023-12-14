package processVideo

import (
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/replace"
	"processAll/util"
	"strings"
)

func RotateAllVideos(dir, pattern, direction, threads string) {
	infos := GetFileInfo.GetAllFilesInfo(dir, pattern)
	for _, in := range infos {
		RotateVideo(in, direction, threads)
	}
}
func RotateVideos(dir, pattern, direction, threads string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		RotateVideo(in, direction, threads)
	}
}
func RotateVideo(in GetFileInfo.BasicInfo, direction, threads string) {
	if strings.Contains(in.PurgePath, "rotate") {
		return
	}
	dst := strings.Join([]string{in.PurgePath, "rotate"}, "")
	os.Mkdir(dst, os.ModePerm)
	fname := in.PurgeName
	fname = replace.ForFileName(fname)
	fname = strings.Join([]string{fname, "mp4"}, ".")
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	var cmd *exec.Cmd
	var transport string
	switch direction {
	case "ToRight":
		transport = "transpose=1"
	case "ToLeft":
		transport = "transpose=2"
	default:
		return
	}
	cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-vf", transport, "-c:v", "libx265", "-c:a", "aac", "-tag:v", "hvc1", "-threads", threads, out)
	err := util.ExecCommand(cmd)
	if err == nil {
		if err = os.RemoveAll(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
		} else {
			slog.Debug("删除成功", slog.Any("源文件", in.FullPath))
		}
	}
}
