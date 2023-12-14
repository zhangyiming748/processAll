package GBK2UTF8

import (
	"github.com/zhangyiming748/mahonia"
	"log/slog"
	"os"
	"path"
	"processAll/GetFileInfo"
	"strings"
	"unicode/utf8"
)

func AllGBKs2UTF8(root, pattern string) {
	infos := GetFileInfo.GetAllFilesInfo(root, pattern)
	for _, in := range infos {
		GBK2UTF8(in)
	}
}
func GBKs2UTF8(dir, pattern string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for _, in := range infos {
		GBK2UTF8(in)
	}
}

func GBK2UTF8(info GetFileInfo.BasicInfo) {
	fp := info.FullPath
	prefix := path.Dir(fp)
	newFp := strings.Join([]string{prefix, "utf8", info.FullName}, string(os.PathSeparator))
	base := path.Dir(newFp)
	os.MkdirAll(base, 0777)
	slog.Debug("执行前的文件基本信息", slog.String("输入", fp), slog.String("输出", newFp), slog.String("前缀", prefix), slog.String("新前缀", base))
	if isUTF8(fp) {
		writeUTF8(newFp, readUTF8(fp))
		slog.Debug("skip", slog.String("编码已经是UTF8,直接复制", info.FullName))
	} else {
		u8 := readGB18030(fp)
		nums := writeUTF8(newFp, u8)
		slog.Debug("文件写入", slog.String("文件名", newFp), slog.Int("字符数", nums))
	}
	if err := os.Remove(fp); err != nil {
		slog.Error("删除源文件出错", slog.Any("错误文本", err), slog.String("文件名", fp))
	} else {
		slog.Debug("删除源文件", slog.String("文件名", fp))
	}
}

func isUTF8(src string) bool {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("查询是否为UTF8时产生错误", slog.Any("错误文本", err))
		}
	}()
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return utf8.Valid(file)
}

func readGB18030(src string) string {
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		slog.Error("编码不存在", slog.Any("错误文本", err))
	}
	return decoder.ConvertString(string(file))
}

func readUTF8(src string) string {
	file, err := os.ReadFile(src)
	if err != nil {
		slog.Error("读取utf8产生错误", slog.Any("错误文本", err))
	}
	return string(file)
}

func writeUTF8(dst, s string) int {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("转写utf8产生错误", slog.Any("错误文本", err))
		}
	}()
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		slog.Error("打开目标文件产生错误", slog.Any("错误文本", err))
	}

	writeString, err := f.WriteString(s)
	if err != nil {
		slog.Error("写文件产生错误", slog.Any("错误文本", err))
	}
	return writeString
}
