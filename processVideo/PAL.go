package processVideo

import (
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetFileInfo"
	"processAll/alert"
	"processAll/replace"
	"processAll/util"
	"strings"
)

func FixAll4x3s(root, pattern, threads string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		Fix4x3(in, threads)
	}
}

func Fix4x3s(src, pattern, threads string) {
	infos := GetFileInfo.GetAllFileInfo(src, pattern)
	for _, in := range infos {
		slog.Debug("横屏视频", slog.Any("视频信息", in))
		Fix4x3(in, threads)
	}
}

func Fix4x3(in GetFileInfo.BasicInfo, threads string) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("错误", slog.String("文件信息", in.FullPath))
			alert.Customize("failed", alert.Ava)
		}
	}()

	dst := in.PurgePath                                 //原始目录
	dst = strings.Join([]string{dst, "resolution"}, "") //二级目录
	fname := in.PurgeName                               //仅文件名
	fname = replace.ForFileName(fname)
	mp4 := strings.Join([]string{fname, "mp4"}, ".")
	os.Mkdir(dst, 0777)
	slog.Debug("新建文件夹", slog.String("全名", dst))
	out := strings.Join([]string{dst, mp4}, string(os.PathSeparator))
	slog.Debug("io", slog.String("源文件:", in.FullPath), slog.String("输出文件:", out))
	var cmd *exec.Cmd = exec.Command("ffmpeg", "-i", in.FullPath, "-aspect", "4:3", "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-threads", threads, out)
	err := util.ExecCommand(cmd)
	if err == nil {
		if err = os.Remove(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.String("源文件", in.FullPath), slog.Any("错误文本", err))
		} else {
			slog.Debug("删除成功", slog.String("源文件", in.FullPath))
		}
	}
}
