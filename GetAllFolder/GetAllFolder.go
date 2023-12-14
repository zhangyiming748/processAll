package GetAllFolder

import (
	"log/slog"
	"os"
	"strings"
)

var (
	all []string
)

/*
递归获取文件夹和全部子文件夹
不包括文件夹本身
*/
func List(dirname string) []string {
	fileInfos, _ := os.ReadDir(dirname)
	var folders []string
	for _, fi := range fileInfos {
		filename := strings.Join([]string{dirname, fi.Name()}, string(os.PathSeparator)) //拼写当前文件夹中所有的文件地址
		if fi.IsDir() {                                                                  //判断是否是文件夹 如果是继续调用把自己的地址作为参数继续调用
			if strings.Contains(filename, "/.") {
				slog.Debug("跳过隐藏文件夹", slog.Any("文件夹名", fi.Name()))
				continue
			}
			slog.Debug("获取到文件夹", slog.String("文件夹名", filename))
			all = append(all, filename)
			folders = append(folders, filename)
			List(filename) //递归调用
		}
	}
	return all
}
