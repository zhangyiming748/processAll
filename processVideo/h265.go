package processVideo

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"processAll/GetAllFolder"
	"processAll/GetFileInfo"
	"processAll/mediaInfo"
	"processAll/model"
	"processAll/replace"
	"processAll/util"
	"runtime"
	"strings"
)

func ProcessAllVideos2H265(root, pattern, threads string) {
	folders := GetAllFolder.List(root)
	folders = append(folders, root)
	for i, folder := range folders {
		slog.Debug(fmt.Sprintf("获取全部子文件夹,正在处理第个 %d/%d 文件夹", i+1, len(folders)))
		ProcessVideos2H265(folder, pattern, threads)
		runtime.GC()
	}
}

func ProcessVideos2H265(dir, pattern, threads string) {
	infos := GetFileInfo.GetAllFileInfo(dir, pattern)
	for i, info := range infos {
		slog.Debug(fmt.Sprintf("获取全部文件,正在处理第个 %d/%d 文件", i+1, len(dir)))
		ProcessVideo2H265(info, threads)
		runtime.GC()
	}
}

func ProcessVideo2H265(in GetFileInfo.BasicInfo, threads string) {
	mi, ok := in.MediaInfo.(mediaInfo.VideoInfo)
	if ok {
		slog.Debug("断言视频mediainfo结构体成功", slog.Any("MediainfoVideo结构体", mi))
	} else {
		slog.Warn("断言视频mediainfo结构体失败")
	}
	slog.Info("获取帧数", slog.Int("当前视频帧数", mi.VideoFrameCount))
	//slog.Debug("文件信息", slog.Any("info", in))
	if strings.Contains(in.FullPath, "h265") {
		slog.Debug("跳过当前已经在h265目录中的文件", slog.String("文件名", in.FullPath))
		return
	}
	prefix := in.PurgePath // 输入文件的纯路径
	slog.Debug("perfix", slog.String("perfix", prefix))

	slog.Debug("fullname", slog.String("fullname", in.FullName))
	middle := "h265"
	if err := os.Mkdir(strings.Join([]string{prefix, middle}, string(os.PathSeparator)), 0777); err != nil {
		if strings.Contains(err.Error(), "file exists") {
			slog.Debug("输出文件夹已存在")
		}
	} else {
		slog.Debug("创建输出文件夹")
	}
	dstPurgeName := replace.ForFileName(in.PurgeName) // 输入文件格式化后的新文件名
	out := strings.Join([]string{in.PurgePath, middle, string(os.PathSeparator), dstPurgeName, ".mp4"}, "")
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("出现错误", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
		}
	}()

	slog.Debug("", slog.String("out", out), slog.String("extName", in.PurgeExt))
	mp4 := strings.Replace(out, in.PurgeExt, "mp4", -1)
	slog.Debug("调试", slog.String("输入文件", in.FullPath), slog.String("输出文件", mp4))
	cmd := exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-map_chapters", "-1", "-threads", threads, mp4)
	// info := GetFileInfo.GetVideoFileInfo(in.FullPath)

	//	if fmt.Sprintf("%v", cpuid.CPU.VendorID) == "Intel" {
	//		bitRate, err := GetFileInfo.GetBitRate(in.FullPath)
	//		slog.Debug("bitrate in h265", slog.String("bitrate", bitRate))
	//		if err != nil {
	//			goto AGAIN
	//		} else {
	//			b := strings.Join([]string{bitRate, "k"}, "")
	//			slog.Debug("获取的比特率", slog.String("bitrate", b))
	//			if b == "" || b == "0" || b == "k" {
	//				slog.Warn("获取比特率失败", slog.String("bitrate", b))
	//				goto AGAIN
	//			} else {
	//				//b := strings.Join([]string{mi.VideoBitRate, "k"}, "")
	//				//ffmpeg -i in.mp4 -c:v hevc_qsv -c:a aac -ac 1 -b:v 100k,
	//				cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "hevc_qsv", "-b:v", b, "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-map_chapters", "-1", "-threads", threads, mp4)
	//				slog.Info("检测到Intel系统", slog.String("命令改变为", fmt.Sprint(cmd)))
	//			}
	//		}
	//	}
	//AGAIN:
	bitRate, _ := GetFileInfo.GetBitRate(in.FullPath)
	slog.Debug("bitrate in h265", slog.String("bitrate", bitRate))

	b := strings.Join([]string{bitRate, "k"}, "")
	slog.Debug("获取的比特率", slog.String("bitrate", b))
	cmd = exec.Command("ffmpeg", "-threads", threads, "-i", in.FullPath, "-c:v", "libx265", "-b:v", b, "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", "-map_chapters", "-1", "-threads", threads, mp4)

	if mi.VideoWidth > 1920 && mi.VideoHeight > 1920 {
		slog.Warn("视频大于1080P需要使用其他程序先处理视频尺寸", slog.Any("原视频", in))
		ResizeVideo(in, threads)
		return
	} else if mi.VideoFormat == "HEVC" {
		if mi.VideoCodecID == "hvc1" {
			slog.Debug(fmt.Sprintf("跳过hevc/hvc1文件"), slog.String("文件名", in.FullPath))
			return
		} else {
			addTag(in)
			slog.Debug("添加标签", slog.String("文件名", in.FullPath))
		}
	}
	slog.Info("生成的命令", slog.String("command", fmt.Sprint(cmd)))
	slog.Info("视频信息", slog.Int("视频帧数", mi.VideoFrameCount), slog.String("比特率", mi.VideoBitRate))
	err := util.ExecCommand(cmd)
	if err != nil {
		return
	}
	slog.Debug("视频编码运行完成")
	if s_size, d_size, diff, err := util.GetDiffFileSize(in.FullPath, mp4); err != nil {
		slog.Warn("文件优化大小计算出错")
	} else {
		save := new(model.Save)
		save.SrcName = in.FullPath
		save.DstName = mp4
		save.SrcSize = s_size
		save.DstSize = d_size
		save.Size = diff
		_, err := save.InsertOne()
		if err != nil {
			fmt.Println("节省的空间 记录插入失败")
		}
	}
	if err == nil {
		if err = os.RemoveAll(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
		} else {
			slog.Debug("删除成功", slog.Any("源文件", in.FullName))
		}
	}
	s := new(model.Save)
	all, err := s.SumSaveAll()
	if err != nil {
		return
	} else {
		//fmt.Printf("节省的空间:%v GB\n", all)
		slog.Info("转码总共节省的空间", slog.String("GB", all))
	}
	slog.Debug("本次转码完成")
}
func addTag(in GetFileInfo.BasicInfo) {
	prefix := strings.Trim(in.FullPath, in.FullName) // 带 /
	dst := strings.Join([]string{prefix, "tag"}, "")
	os.Mkdir(dst, 0777)
	target := strings.Join([]string{dst, in.FullName}, string(os.PathSeparator))
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-c:v", "copy", "-c:a", "copy", "-ac", "1", "-c:s", "copy", "-tag:v", "hvc1", "-map_chapters", "-1", target)
	err := util.ExecCommand(cmd)
	if err == nil {
		if err = os.RemoveAll(in.FullPath); err != nil {
			slog.Warn("删除失败", slog.Any("源文件", in.FullPath), slog.Any("错误", err))
		} else {
			slog.Debug("删除成功", slog.Any("源文件", in.FullName))
		}
	}
}
