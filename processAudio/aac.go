package processAudio

import (
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/replace"
	"processAll/util"
	"strings"
)

func Audio2AAC(in GetFileInfo.BasicInfo) {
	// 执行转换
	fname := replace.ForFileName(in.PurgeName)
	out := strings.Join([]string{in.PurgePath, fname, ".aac"}, "")
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-ac", "1", out)
	err := util.ExecCommand(cmd)

	if err == nil {
		if err = os.RemoveAll(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误", err))
		} else {
			slog.Debug("删除成功", slog.String("源文件", in.FullPath))
		}
	}
}

func Audios2AAC(dir, pattern string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		Audio2AAC(in)
	}
}

func AllAudios2AAC(root, pattern string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		Audio2AAC(in)
	}
}
