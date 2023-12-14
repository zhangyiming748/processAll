package util

import (
	"errors"
	"log/slog"
	"os"
)

const (
	GB = 1024 * 1024 * 1024
)

/*
计算所提供文件大小 字节
*/

func GetSize(fp string) (uint64, error) {
	if file, err := os.Open(fp); err != nil {
		return 0, err
	} else if info, err := file.Stat(); err != nil {
		return 0, err
	} else {
		defer file.Close()
		size := uint64(info.Size())
		return size, nil
	}
}

/*
计算所给定的两个文件大小差 返回GB
*/
func GetDiffSize(src, dst uint64) (float64, error) {
	if dst >= src {
		slog.Warn("处理后的文件比源文件更大,放弃", slog.Uint64("源文件大小", src), slog.Uint64("目标文件大小", dst))
		return 0, errors.New("处理后的文件比源文件更大,放弃")
	}
	save := float64(src-dst) / GB
	return save, nil
}

/*
计算所给定的两个文件名 返回GB
*/
func GetDiffFileSize(src, dst string) (srcSize uint64, dstSize uint64, save float64, err error) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		slog.Warn("获取文件差值出错", slog.Any("错误原文", err))
	//		return
	//	}
	//}()
	srcSize, err = GetSize(src)
	dstSize, err = GetSize(dst)
	save, err = GetDiffSize(srcSize, dstSize)
	if err != nil {
		return 0, 0, 0, errors.New("获取文件差值出错")
	}
	return srcSize, dstSize, save, nil
}
