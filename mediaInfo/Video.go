package mediaInfo

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"strconv"
)

type VideoMedia struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                  string `json:"@type"`
			ID                    string `json:"ID"`
			VideoCount            string `json:"VideoCount,omitempty"`
			AudioCount            string `json:"AudioCount,omitempty"`
			FileExtension         string `json:"FileExtension,omitempty"`
			Format                string `json:"Format"`
			FileSize              string `json:"FileSize,omitempty"`
			Duration              string `json:"Duration"`
			OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRate        string `json:"OverallBitRate,omitempty"`
			FrameRate             string `json:"FrameRate"`
			FrameCount            string `json:"FrameCount,omitempty"`
			FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
			Extra                 struct {
				OverallBitRatePrecisionMin string `json:"OverallBitRate_Precision_Min"`
				OverallBitRatePrecisionMax string `json:"OverallBitRate_Precision_Max"`
				FileExtensionInvalid       string `json:"FileExtension_Invalid"`
			} `json:"extra,omitempty"`
			StreamOrder              string `json:"StreamOrder,omitempty"`
			MenuID                   string `json:"MenuID,omitempty"`
			FormatProfile            string `json:"Format_Profile,omitempty"`
			FormatLevel              string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC      string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames  string `json:"Format_Settings_RefFrames,omitempty"`
			FormatSettingsGOP        string `json:"Format_Settings_GOP,omitempty"`
			CodecID                  string `json:"CodecID,omitempty"`
			Width                    string `json:"Width,omitempty"`
			Height                   string `json:"Height,omitempty"`
			SampledWidth             string `json:"Sampled_Width,omitempty"`
			SampledHeight            string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio         string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio       string `json:"DisplayAspectRatio,omitempty"`
			FrameRateNum             string `json:"FrameRate_Num,omitempty"`
			FrameRateDen             string `json:"FrameRate_Den,omitempty"`
			Standard                 string `json:"Standard,omitempty"`
			ColorSpace               string `json:"ColorSpace,omitempty"`
			ChromaSubsampling        string `json:"ChromaSubsampling,omitempty"`
			BitDepth                 string `json:"BitDepth,omitempty"`
			ScanType                 string `json:"ScanType,omitempty"`
			Delay                    string `json:"Delay,omitempty"`
			DelaySource              string `json:"Delay_Source,omitempty"`
			ColourRange              string `json:"colour_range,omitempty"`
			ColourRangeSource        string `json:"colour_range_Source,omitempty"`
			FormatVersion            string `json:"Format_Version,omitempty"`
			FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
			MuxingMode               string `json:"MuxingMode,omitempty"`
			BitRateMode              string `json:"BitRate_Mode,omitempty"`
			BitRate                  string `json:"bit_rate,omitempty"`
			Channels                 string `json:"Channels,omitempty"`
			ChannelPositions         string `json:"ChannelPositions,omitempty"`
			ChannelLayout            string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
			SamplingRate             string `json:"SamplingRate,omitempty"`
			SamplingCount            string `json:"SamplingCount,omitempty"`
			CompressionMode          string `json:"Compression_Mode,omitempty"`
			VideoDelay               string `json:"Video_Delay,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

type VideoInfo struct {
	Type          string  `json:"@type"`                   // General Audio Video
	VideoCount    string  `json:"VideoCount,omitempty"`    // 视频数
	AudioCount    string  `json:"AudioCount,omitempty"`    // 音轨数
	FileExtension string  `json:"FileExtension,omitempty"` // 扩展名
	FileSize      uint64  `json:"FileSize,omitempty"`      // 文件大小
	Duration      float64 `json:"Duration,omitempty"`      // 持续时间

	VideoFormat     string  `json:"VideoFormat,omitempty"`     // 视频格式 HEVC AVC
	VideoCodecID    string  `json:"VideoCodecID,omitempty"`    // 视频编码 hvc1 avc1
	VideoWidth      int     `json:"VideoWidth,omitempty"`      // 宽度 像素
	VideoHeight     int     `json:"VideoHeight,omitempty"`     // 高度 像素
	VideoFrameRate  float64 `json:"VideoFrameRate,omitempty"`  // 视频帧率
	VideoFrameCount int     `json:"VideoFrameCount,omitempty"` // 视频帧数
	VideoBitDepth   string  `json:"VideoBitDepth,omitempty"`   // 视频位深
	VideoBitRate    string  `json:"BitRate,omitempty"`         // 视频比特率

	AudioFormat     string  `json:"AudioFormat,omitempty"`     //音频格式  AAC
	AudioCodecID    string  `json:"AudioCodecID,omitempty"`    //音频编码 mp4a-40-2
	AudioFrameRate  float64 `json:"AudioFrameRate,omitempty"`  //音频帧率
	AudioFrameCount int     `json:"AudioFrameCount,omitempty"` // 音频帧数
	AudioBitDepth   string  `json:"AudioBitDepth,omitempty"`   // 音频位深
	Channels        int     `json:"Channels,omitempty"`        // 声道数

}

/*
获取音频文件的mediainfo信息
*/
func GetVideoMedia(absPath string) VideoInfo {
	var vm VideoMedia
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("C:/Program Files/MediaInfo/MediaInfo.exe", absPath, "--Output=JSON")
	}
	slog.Debug("生成的命令", slog.String("命令原文", fmt.Sprint(cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Warn("运行mediainfo命令", slog.String("命令原文", fmt.Sprint(cmd)), slog.Any("产生的错误", err))
	}
	if err = json.Unmarshal(output, &vm); err != nil {
		slog.Warn("解析json", slog.Any("产生的错误", err))
	} else {
		slog.Debug("解析json成功", slog.Any("解析后的json", vm))
	}
	return VideoMedia2Info(vm)
}

func VideoMedia2Info(vm VideoMedia) VideoInfo {
	vi := new(VideoInfo)
	vi.Type = "Video"
	for _, kind := range vm.Media.Track {
		if kind.Type == "General" {
			vi.FileExtension = kind.FileExtension
			FileSize, err := strconv.ParseUint(kind.FileSize, 10, 64)
			if err != nil {
				slog.Warn("文件大小转换错误", slog.String("原始数值", kind.FileSize))
			} else {
				vi.FileSize = FileSize
			}
			vi.AudioCount = kind.AudioCount
		}
		if kind.Type == "Video" {
			vi.VideoCodecID = kind.CodecID
			vi.VideoFormat = kind.Format
			if Width, err := strconv.Atoi(kind.Width); err != nil {
				slog.Warn("视频宽度转换错误")
			} else {
				vi.VideoWidth = Width
			}
			if Height, err := strconv.Atoi(kind.Height); err != nil {
				slog.Warn("视频高度转换错误")
			} else {
				vi.VideoHeight = Height
			}
			vi.VideoBitDepth = kind.BitDepth
			if VideoFrameRate, err := strconv.ParseFloat(kind.FrameRate, 64); err != nil {
				slog.Warn("视频帧率转换错误")
			} else {
				vi.VideoFrameRate = VideoFrameRate
			}
			if VideoFrameCount, err := strconv.Atoi(kind.FrameCount); err != nil {
				slog.Warn("视频帧数转换错误")
			} else {
				vi.VideoFrameCount = VideoFrameCount
			}
			vi.VideoBitRate = kind.BitRate

		}
		if kind.Type == "Audio" {
			vi.AudioFormat = kind.Format
			vi.AudioCodecID = kind.CodecID
			if AudioFrameRate, err := strconv.ParseFloat(kind.FrameRate, 64); err != nil {
				slog.Warn("音频帧率转换错误")
			} else {
				vi.AudioFrameRate = AudioFrameRate
			}
			if AudioFrameCount, err := strconv.Atoi(kind.FrameCount); err != nil {
				slog.Warn("音频帧数转换错误")
			} else {
				vi.AudioFrameCount = AudioFrameCount
			}
			vi.AudioBitDepth = kind.BitDepth
			if Channels, err := strconv.Atoi(kind.Channels); err != nil {
				slog.Warn("视频声道数转换错误")
			} else {
				vi.Channels = Channels
			}
		}
	}
	return *vi
}
