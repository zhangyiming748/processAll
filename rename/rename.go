package rename

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"processAll/replace"
	"regexp"
	"strings"
)

func rename(root string) {
	//folders = append(folders, root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			//fmt.Println(info.Name())
			oldname := info.Name()
			oldPath := strings.Join([]string{root, oldname}, string(os.PathSeparator))
			//newName := strings.Replace(oldname, "-", "", -1)
			newName, newPath := old2new(oldname, root)
			fmt.Printf("oldName:%v\noldPath:%v\nnewName:%v\nnewPath:%v\n", oldname, oldPath, newName, newPath)
			//err = os.Rename(oldPath, newPath)
			//if err != nil {
			//	fmt.Printf("重命名出错\n")
			//	return err
			//}
		}
		//fmt.Printf("Path: %s, Size: %d bytes\n", path, info.Size())
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func old2new(oldName, oldPath string) (newName, newPath string) {

	/****************重命名逻辑*******************/
	//tme = strings.Replace(oldname, "AI换脸", "", -1)

	newName = replace.ForFileName(oldName)
	match, _ := regexp.MatchString(`[A-Z]\d{3}`, newName)
	if match {
		//slog.Info("matched!", slog.String("匹配到的字符串", match))
		//fmt.Println(match)
		re := regexp.MustCompile(`[A-Z]\d{3,4}`)
		matches := re.FindStringSubmatch(newName)
		for _, bingo := range matches {
			slog.Info("matched!", slog.String("匹配到的字符串", bingo))
			newName = strings.Replace(newName, bingo, "", 1)
		}
	}
	/****************重命名逻辑*******************/

	newPath = strings.Join([]string{oldPath, newName}, string(os.PathSeparator))

	return newName, newPath
}

/*
重命名文件删除重复内容
*/
func RenameForDup(dir, dup string) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range readDir {
		//t.Logf("%+v\n", file)
		oldName := file.Name()
		oldPath := strings.Join([]string{dir, oldName}, string(os.PathSeparator))
		newName := strings.Replace(oldName, dup, "", 1)
		newPath := strings.Join([]string{dir, newName}, string(os.PathSeparator))
		slog.Info("summary", slog.String("旧文件名", oldName), slog.String("新文件名", newPath))
		err := os.Rename(oldPath, newName)
		if err != nil {
			//t.Log(err)
			return
		}
	}
}
