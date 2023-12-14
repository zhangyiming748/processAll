package mediaInfo

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"strconv"
)

type GeneralMedia struct {
	CreatingLibrary struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
		Url     string `json:"Url"`
	} `json:"CreatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                  string `json:"@type"`                   // General
			FileExtension         string `json:"FileExtension,omitempty"` // 文件扩展名
			FileSize              string `json:"FileSize,omitempty"`      // 文件大小 字节
			StreamSize            string `json:"StreamSize,omitempty"`
			FileModifiedDate      string `json:"FileModifiedDate,omitempty"`      // 文件最后修改时间
			FileModifiedDateLocal string `json:"FileModifiedDateLocal,omitempty"` // 文件最后修改本地时间
		} `json:"track"`
	} `json:"media"`
}
type GeneralInfo struct {
	Type          string `json:"@type"`                   // General
	FileExtension string `json:"FileExtension,omitempty"` // 扩展名
	FileSize      uint64 `json:"FileSize,omitempty"`      // 文件大小
}

/*
获取其他文件的mediainfo信息
*/
func GetGeneralMedia(absPath string) GeneralInfo {
	var gm GeneralMedia
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("C:/Program Files/MediaInfo/MediaInfo.exe", absPath, "--Output=JSON")
	}
	slog.Debug("生成的命令", slog.String("命令原文", fmt.Sprint(cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Warn("运行mediainfo命令", slog.String("命令原文", fmt.Sprint(cmd)), slog.Any("产生的错误", err))
	}
	if err = json.Unmarshal(output, &gm); err != nil {
		slog.Warn("解析json", slog.Any("产生的错误", err))
	} else {
		slog.Debug("解析json成功", slog.Any("解析后的json", gm))
	}
	return GeneralMedia2Info(gm)
}

func GeneralMedia2Info(gm GeneralMedia) GeneralInfo {
	gi := new(GeneralInfo)
	gi.Type = "General"
	for _, kind := range gm.Media.Track {
		if kind.Type == "General" {
			gi.FileExtension = kind.FileExtension
			FileSize, err := strconv.ParseUint(kind.FileSize, 10, 64)
			if err != nil {
				slog.Warn("文件大小转换错误", slog.String("原始数值", kind.FileSize))
			} else {
				gi.FileSize = FileSize
			}
		}
	}
	return *gi
}
