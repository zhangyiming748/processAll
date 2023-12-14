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

/*
执行文件夹和子文件夹中音频增大电平
*/
func LouderAllAudios(root, pattern string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		LouderAudio(in)
	}
}

/*
执行一个文件夹中音频增大电平
*/
func LouderAudios(dir, pattern string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		LouderAudio(in)
	}
}

/*
单个音频增大电平的完整函数
*/
func LouderAudio(in GetFileInfo.BasicInfo) {
	src := in.PurgePath                              //原文件目录 带有最后一个 /
	dst := strings.Join([]string{src, "Louder"}, "") //目标文件目录
	os.Mkdir(dst, 0777)
	fname := replace.ForFileName(in.PurgeName)
	fname = strings.Join([]string{fname, "aac"}, ".")
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
	//跳过已经增大电平的文件夹
	if strings.Contains(in.FullPath, "Louder") {
		return
	}
	Louder(in.FullPath, out)

}

/*
仅使用输入输出和增大电平
*/
func Louder(in, out string) {
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", "volume=3.0", out)
	util.ExecCommand(cmd)
	if err := os.RemoveAll(in); err != nil {
		slog.Warn("删除失败", slog.String("源文件", in), slog.Any("错误内容", err))
	} else {
		slog.Debug("删除成功", slog.String("源文件", in))
	}
}
