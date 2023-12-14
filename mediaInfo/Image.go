package mediaInfo

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"strconv"
)

type ImageMedia struct {
	CreatingLibrary struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
		Url     string `json:"Url"`
	} `json:"CreatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type          string `json:"@type"`                   // General Image
			ImageCount    string `json:"ImageCount,omitempty"`    // 图片数
			FileExtension string `json:"FileExtension,omitempty"` // jpg
			Format        string `json:"Format,omitempty"`        // JPEG
			FileSize      string `json:"FileSize,omitempty"`      // 文件大小 字节
			Width         string `json:"Width,omitempty"`         // 图片宽度
			Height        string `json:"Height,omitempty"`        // 图片高度
			BitDepth      string `json:"BitDepth,omitempty"`      // 位深
		} `json:"track"`
	} `json:"media"`
}
type ImageInfo struct {
	Type          string `json:"@type"`                   // General Audio
	ImageCount    string `json:"ImageCount,omitempty"`    // 图片数
	FileExtension string `json:"FileExtension,omitempty"` // 扩展名
	Format        string `json:"Format"`                  // 容器格式 MPEG-4 AAC
	FileSize      uint64 `json:"FileSize,omitempty"`      // 文件大小
	Width         int    `json:"Width"`                   // 宽度 像素
	Height        int    `json:"Height"`                  // 高度 像素
	BitDepth      string `json:"BitDepth"`                // 位深
}

/*
获取音频文件的mediainfo信息
*/
func GetImageMedia(absPath string) ImageInfo {
	var im ImageMedia
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("C:/Program Files/MediaInfo/MediaInfo.exe", absPath, "--Output=JSON")
	}
	slog.Debug("生成的命令", slog.String("命令原文", fmt.Sprint(cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Warn("运行mediainfo命令", slog.String("命令原文", fmt.Sprint(cmd)), slog.Any("产生的错误", err))
	}
	if err = json.Unmarshal(output, &im); err != nil {
		slog.Warn("解析json", slog.Any("产生的错误", err))
	} else {
		slog.Debug("解析json成功", slog.Any("解析后的json", im))
	}
	return ImageMedia2Info(im)
}

func ImageMedia2Info(im ImageMedia) ImageInfo {
	ii := new(ImageInfo)
	ii.Type = "Image"
	for _, kind := range im.Media.Track {
		if kind.Type == "General" {
			ii.ImageCount = kind.ImageCount
			ii.FileExtension = kind.FileExtension
			FileSize, err := strconv.ParseUint(kind.FileSize, 10, 64)
			if err != nil {
				slog.Warn("文件大小转换错误", slog.String("原始数值", kind.FileSize))
			} else {
				ii.FileSize = FileSize
			}
		}
		if kind.Type == "Image" {
			ii.ImageCount = kind.Format
			ii.Format = kind.Format
			if Width, err := strconv.Atoi(kind.Width); err != nil {
				slog.Warn("图片视频宽度转换错误")
			} else {
				ii.Width = Width
			}
			if Height, err := strconv.Atoi(kind.Height); err != nil {
				slog.Warn("图片高度转换错误")
			} else {
				ii.Height = Height
			}
			ii.BitDepth = kind.BitDepth
		}
	}
	return *ii
}
