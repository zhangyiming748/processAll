package GetFileInfo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/zhangyiming748/filetype"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
)

/*
判断当前文件扩展名是否在选定的扩展名列表中
*/
func In(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

/*
通过文件头判断文件类型
选择合适的结构体
*/
func SelectTypeByHead(fp string) string {
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	if filetype.IsVideo(head) {
		slog.Debug("File is a video")
		return "Video"
	} else if filetype.IsAudio(head) {
		slog.Debug("File is a audio")
		return "Audio"
	} else if filetype.IsImage(head) {
		slog.Debug("File is a image")
		return "Image"
	} else {
		slog.Debug("File is a general")
		return "General"
	}
}

/*
获取文件MD5
*/
func GetMD5(fp string) string {
	pFile, err := os.Open(fp)
	if err != nil {
		slog.Warn("获取md5打开文件出错", slog.String("文件名", fp), slog.Any("错误文本", err))
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil))
}

/*
通过扩展名判断文件类型
选择合适的结构体
*/
func SelectType(ext string) string {
	switch ext {
	case "jpeg":
		return "Image"
	case "JPEG":
		return "Image"
	case "jpg":
		return "Image"
	case "JPG":
		return "Image"
	case "png":
		return "Image"
	case "PNG":
		return "Image"
	case "webp":
		return "Image"
	case "WEBP":
		return "Image"
	case "tif":
		return "Image"
	case "TIF":
		return "Image"

	case "mp3":
		return "Audio"
	case "MP3":
		return "Audio"
	case "aac":
		return "Audio"
	case "AAC":
		return "Audio"
	case "m4a":
		return "Audio"
	case "M4A":
		return "Audio"
	case "flac":
		return "Audio"
	case "FLAC":
		return "Audio"
	case "wma":
		return "Audio"
	case "WMA":
		return "Audio"
	case "wav":
		return "Audio"
	case "WAV":
		return "Audio"
	case "ogg":
		return "Audio"
	case "OGG":
		return "Audio"

	case "webm":
		return "Video"
	case "WEBM":
		return "Video"
	case "mkv":
		return "Video"
	case "MKV":
		return "Video"
	case "m4v":
		return "Video"
	case "M4V":
		return "Video"
	case "mp4":
		return "Video"
	case "MP4":
		return "Video"
	case "mov":
		return "Video"
	case "MOV":
		return "Video"
	case "avi":
		return "Video"
	case "AVI":
		return "Video"
	case "wmv":
		return "Video"
	case "WMV":
		return "Video"
	case "ts":
		return "Video"
	case "TS":
		return "Video"
	case "rmvb":
		return "Video"
	case "RMVB":
		return "Video"
	case "flv":
		return "Video"
	case "FLV":
		return "Video"
	case "vob":
		return "Video"
	case "VOB":
		return "Video"
	default:
		return "General"
	}
}

func GetBitRate(fp string) (string, error) {
	ff, _ := exec.Command("ffmpeg", "-i", fp).CombinedOutput()
	re := regexp.MustCompile(`bitrate:\s*(\d+)\s*kb/s`)
	matches := re.FindAllStringSubmatch(string(ff), -1)
	if len(matches) > 0 {
		str := fmt.Sprintf("比特率字段:%s\n", matches)
		fmt.Println(str)
		num := fmt.Sprint(matches[0][1])
		slog.Debug(num)
		return num, nil
	}

	return "", nil
}
