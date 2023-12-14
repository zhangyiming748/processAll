package mediaInfo

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"strconv"
)

type AudioMedia struct {
	CreatingLibrary struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
		Url     string `json:"Url"`
	} `json:"CreatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type           string `json:"@type"`                    // General Audio
			AudioCount     string `json:"AudioCount,omitempty"`     // 音轨数
			FileExtension  string `json:"FileExtension,omitempty"`  // 扩展名
			Format         string `json:"Format,omitempty"`         // 容器格式 MPEG-4 AAC
			FileSize       string `json:"FileSize,omitempty"`       // 文件大小
			Duration       string `json:"Duration,omitempty"`       // 持续时间 秒
			OverallBitRate string `json:"OverallBitRate,omitempty"` // 总比特率
			BitRate        string `json:"BitRate,omitempty"`        // 比特率
			Channels       string `json:"Channels,omitempty"`       // 声道数
			FrameRate      string `json:"FrameRate,omitempty"`      // 帧率
			FrameCount     string `json:"FrameCount,omitempty"`     // 帧数
		} `json:"track"`
	} `json:"media"`
}

type AudioInfo struct {
	Type           string  `json:"@type"`                    // General Audio
	AudioCount     string  `json:"AudioCount,omitempty"`     // 音轨数
	FileExtension  string  `json:"FileExtension,omitempty"`  // 扩展名
	Format         string  `json:"Format,omitempty"`         // 容器格式 MPEG-4 AAC
	FileSize       uint64  `json:"FileSize,omitempty"`       // 文件大小
	Duration       float64 `json:"Duration,omitempty"`       // 持续时间 秒
	OverallBitRate string  `json:"OverallBitRate,omitempty"` // 总比特率
	BitRate        string  `json:"BitRate,omitempty"`        // 比特率
	Channels       int     `json:"Channels,omitempty"`       // 声道数
	FrameRate      float64 `json:"FrameRate,omitempty"`      // 帧率
	FrameCount     int     `json:"FrameCount,omitempty"`     // 帧数
}

/*
获取音频文件的mediainfo信息
*/
func GetAudioMedia(absPath string) AudioInfo {
	var am AudioMedia
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("C:/Program Files/MediaInfo/MediaInfo.exe", absPath, "--Output=JSON")
	}
	slog.Debug("生成的命令", slog.String("命令原文", fmt.Sprint(cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Warn("运行mediainfo命令", slog.String("命令原文", fmt.Sprint(cmd)), slog.Any("产生的错误", err))
	}
	if err = json.Unmarshal(output, &am); err != nil {
		slog.Warn("解析json", slog.Any("产生的错误", err))
	} else {
		slog.Debug("解析json成功", slog.Any("解析后的json", am))
	}
	return AudioMedia2Info(am)
}

func AudioMedia2Info(am AudioMedia) AudioInfo {
	ai := new(AudioInfo)
	ai.Type = "Audio"
	for _, kind := range am.Media.Track {
		if kind.Type == "General" {
			ai.AudioCount = kind.AudioCount
			ai.FileExtension = kind.FileExtension
			FileSize, err := strconv.ParseUint(kind.FileSize, 10, 64)
			if err != nil {
				slog.Warn("文件大小转换错误", slog.String("原始数值", kind.FileSize))
			} else {
				ai.FileSize = FileSize
			}
			if Duration, err := strconv.ParseFloat(kind.Duration, 64); err != nil {
				slog.Warn("文件时长转换错误", slog.String("原始数值", kind.Duration))
			} else {
				ai.Duration = Duration
			}
		}
		if kind.Type == "Audio" {
			ai.Format = kind.Format
			if FrameRate, err := strconv.ParseFloat(kind.FrameRate, 64); err != nil {
				slog.Warn("音频帧率转换错误")
			} else {
				ai.FrameRate = FrameRate
			}
			if FrameCount, err := strconv.Atoi(kind.FrameCount); err != nil {
				slog.Warn("音频帧数转换错误")
			} else {
				ai.FrameCount = FrameCount
			}
			if Channels, err := strconv.Atoi(kind.Channels); err != nil {
				slog.Warn("音频声道数转换错误")
			} else {
				ai.Channels = Channels
			}
		}
	}
	return *ai
}
