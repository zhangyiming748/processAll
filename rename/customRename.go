package rename

import (
	"fmt"
	"os"
	"processAll/GetAllFolder"
	"processAll/GetFileInfo"
	"processAll/replace"
	"runtime"
	"strings"
)

/*
get all name in folders and replace clean
*/
func cleanName(root, pattern string) {
	folders := GetAllFolder.List(root)
	folders = append(folders, root)
	for _, folder := range folders {
		files := GetFileInfo.GetAllFilesInfo(folder, pattern)
		for _, file := range files {
			fmt.Printf("%+v\n", file)
			clean(file)
		}
		runtime.GC()
	}
}
func clean(info GetFileInfo.BasicInfo) {
	file, err := os.OpenFile("report.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	oldPurgeName := info.PurgeName
	newPurgeName := replace.ForFileName(oldPurgeName)
	if oldPurgeName == newPurgeName {
		file.WriteString(fmt.Sprintf("跳过已经处理的文件:%v\n", info.FullPath))
		return
	}
	newFileName := strings.Join([]string{newPurgeName, info.PurgeExt}, ".")
	newFullPath := strings.Join([]string{info.PurgePath, newFileName}, "")
	file.WriteString(fmt.Sprintf("旧文件名:%s\t新文件名:%s\n", info.FullPath, newFullPath))
	err = os.Rename(info.FullPath, newFullPath)
	if err != nil {
		file.WriteString(fmt.Sprintf("重命名出错的文件filename:%v\n", info.FullPath))
	}
	file.Sync()
}
