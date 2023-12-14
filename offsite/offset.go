package offsite

import (
	"fmt"
	"os"
	"path"
	"processAll/GetFileInfo"
	"strconv"
	"strings"
)

func offset(dir string, off int) {
	names := GetFileInfo.GetAllFiles(dir)
	for _, name := range names {
		suffix := path.Ext(name)
		prefix := strings.Trim(name, suffix)
		old, _ := strconv.Atoi(prefix)
		newer := strconv.Itoa(old + off)
		nfname := strings.Join([]string{newer, suffix}, "")
		before := strings.Join([]string{dir, name}, string(os.PathSeparator))
		after := strings.Join([]string{dir, nfname}, string(os.PathSeparator))
		fmt.Printf("旧文件名:%s\t新文件名:%s\n", before, after)
		os.Rename(before, after)
	}
}
func addZero(dir string) {
	names := GetFileInfo.GetAllFiles(dir)
	for _, name := range names {
		//fmt.Println(name)
		suffix := path.Ext(name)
		prefix := strings.Trim(name, suffix)
		if len(prefix) == 3 {
			fmt.Printf("三位数%v\n", name)
			newName := strings.Join([]string{"0", prefix}, "")
			fmt.Printf("三位数补0后%v\n", newName)
			newName = strings.Join([]string{newName, suffix}, "")
			before := strings.Join([]string{dir, name}, string(os.PathSeparator))
			after := strings.Join([]string{dir, newName}, string(os.PathSeparator))
			fmt.Printf("源文件名%s\t新文件名%s\n", before, after)
			os.Rename(before, after)
		}
	}
}
