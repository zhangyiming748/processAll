package processImage

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

/*
转换所有子文件夹下的图片为AVIF
*/
func ProcessAllImages(root, pattern, threads string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		ProcessImage(in, threads)
	}
}

/*
转换一个文件夹下的图片为AVIF
*/
func ProcessImages(dir, pattern, threads string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		ProcessImage(in, threads)
	}
}

/*
转换一张图片为AVIF
*/
func ProcessImage(in GetFileInfo.BasicInfo, threads string) {
	mi, ok := in.MediaInfo.(mediaInfo.ImageInfo)
	if ok {
		slog.Debug("断言图片mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
	} else {
		slog.Warn("断言图片mediainfo结构体失败")
	}

	cleanName := replace.ForFileName(in.PurgeName)
	out := strings.Join([]string{in.PurgePath, cleanName, ".avif"}, "")

	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libaom-av1", "-still-picture", "1", "-threads", threads, out)
	slog.Debug("ffmpeg", slog.Any("生成的命令", fmt.Sprint(cmd)))
	err := util.ExecCommand(cmd)

	if err == nil {
		if err = os.RemoveAll(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
		} else {
			slog.Debug("删除成功", slog.Any("源文件", in.FullPath))
		}
	}
}
