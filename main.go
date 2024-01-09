package main

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"processAll/GBK2UTF8"
	"processAll/GetAllFolder"
	"processAll/GetFileInfo"
	"processAll/alert"
	"processAll/merge"
	"processAll/processAudio"
	"processAll/processImage"
	"processAll/processVideo"
	"processAll/sql"
	"processAll/telegraph"
	"processAll/util"
	"processAll/ytdlp"
	"runtime"
	"strings"
	"time"
)

// todo 除图片流程添加数据库代码

type summarize struct {
	Files         int64 // 总处理文件数
	SuccessAudios int64 // 总成功的音频数
	SuccessVideo  int64 // 总成功的视频数
	SuccessImage  int64 // 总成功的图片数
	Failure       int64 // 总失败数
}

/*
程序运行前选择日志和数据库配置
*/
func init() {
	setLog()
	util.SetRoot()
	sql.SetEngine()
}

func main() {
	slog.Info("当前的机器信息", slog.String("CPU名称", cpuid.CPU.BrandName), slog.Int("物理核心数", cpuid.CPU.PhysicalCores), slog.Int("每个核心的线程数", cpuid.CPU.ThreadsPerCore), slog.Int("逻辑核心数", cpuid.CPU.LogicalCores), slog.String("品牌", fmt.Sprintf("%v", cpuid.CPU.VendorID)), slog.Int64("频率", cpuid.CPU.Hz))
	if len(os.Args) > 1 {
		//go run main.go <任务> <参数1> <参数2> ...
		slog.Info("切换手动模式")
		opt := slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
		file := "Process.log"
		logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
		slog.SetDefault(logger)
		for i, v := range os.Args {
			fmt.Printf("第%d个参数%s\n", i+1, v)
		}
		m := make(map[string]string)
		m["bilibili"] = "/sdcard/Android/data/tv.danmaku.bili/download"
		m["bilibilihd"] = "/sdcard/Android/data/tv.danmaku.bilibilihd/download"
		if len(os.Args) >= 3 {
			switch os.Args[1] {
			case "merge": // main merge bilibilihd
				/*
					#!/data/data/com.termux/files/usr/bin/bash
					echo 合成bilibili缓存
					sudo /data/data/com.termux/files/home/ProcessAVI/main merge bilibilihd
				*/
				fmt.Println("main merge bilibilihd")
				merge.Merge(m[os.Args[2]])
				slog.Debug("", slog.String("运行目录", m[os.Args[2]]))
				os.Exit(0)
			}
		} else {
			slog.Warn("参数个数错误 退出", slog.String("例如", "go run main.go bilibilihd merge"))
			os.Exit(0)
		}

	}
	slog.Info("start!", slog.String("程序运行的根目录", util.GetRoot()))
	defer final()
	mission := util.GetVal("main", "mission")
	var (
		root      string
		pattern   string
		threads   string
		direction string
	)
	staterOn := util.GetVal("StartAt", "time")
	if level := util.GetVal("log", "level"); level != "Debug" {
		// Debug 模式下不等待开始运行时间
		startOn(staterOn)
	}
	start := time.Now()
	end := time.Now()
	if quiet := util.GetVal("alert", "quiet"); quiet == "yes" {
		os.Setenv("QUIET", "True")
		slog.Debug("静音模式")
	}
	defer func() {
		if email := util.GetVal("alert", "email"); email == "yes" {
			slog.Debug("发送任务完成邮件")
			sendEmail(start, end)
		}
	}()
	util.ExitAfterRun()
	switch mission {
	case "OGG":
		pattern = util.GetVal("pattern", "audio")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "audio")
		slog.Debug("开始音频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
		processAudio.AllAudios2OGG(root, pattern)
	case "i&v":
		{
			pattern = util.GetVal("pattern", "video")
			pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
			root = util.GetVal("root", "video")
			threads = util.GetVal("thread", "threads")
			slog.Debug("开始视频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
			processVideo.ProcessAllVideos2H265(root, pattern, threads)
		}
		{
			pattern = util.GetVal("pattern", "image")
			pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
			root = util.GetVal("root", "image")
			threads = util.GetVal("thread", "threads")
			pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
			slog.Debug("开始图片处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
			processImage.ProcessAllImages(root, pattern, threads)
		}
	case "video":
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "video")
		threads = util.GetVal("thread", "threads")
		slog.Debug("开始视频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
		processVideo.ProcessAllVideos2H265(root, pattern, threads)
	case "audio":
		pattern = util.GetVal("pattern", "audio")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "audio")
		slog.Debug("开始音频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
		processAudio.AllAudios2AAC(root, pattern)
	case "image":
		pattern = util.GetVal("pattern", "image")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "image")
		threads = util.GetVal("thread", "threads")
		slog.Debug("开始图片处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
		processImage.ProcessAllImages(root, pattern, threads)
	case "rotate":
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "video")
		threads = util.GetVal("thread", "threads")
		direction = util.GetVal("rotate", "direction")
		slog.Debug("开始旋转视频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads), slog.String("方向", direction))
		processVideo.RotateAllVideos(root, pattern, direction, threads)
	case "resize":
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		root = util.GetVal("root", "video")
		threads = util.GetVal("thread", "threads")
		slog.Debug("开始缩小视频处理进程", slog.String("根目录", root), slog.String("pattern", pattern), slog.String("进程数", threads))
		processVideo.ResizeAllVideos(root, pattern, threads)
	case "avmerger":
		root = util.GetVal("root", "bilibili")
		slog.Debug("开始合并哔哩哔哩进程", slog.String("根目录", root))
		merge.Merge(root)
	case "extractAAC":
		root = util.GetVal("root", "bilibili")
		slog.Debug("开始提取哔哩哔哩音频进程", slog.String("根目录", root))
		merge.ExtractAAC(root)
	case "speedUpAudio":
		root = util.GetVal("root", "audio")
		pattern = util.GetVal("pattern", "audio")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		speed := util.GetVal("speed", "audition")
		processAudio.SpeedUpAllAudios(root, pattern, speed)
		slog.Debug("开始有声小说加速处理", slog.String("根目录", root))
	case "speedUpVideo":
		root = util.GetVal("root", "video")
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		speed := util.GetVal("speed", "ffmpeg")
		processVideo.SpeedUpAllVideos(root, pattern, speed)
		slog.Debug("开始音视频同步加速处理", slog.String("根目录", root))
	case "txt":
		root = util.GetVal("root", "txt")
		pattern = util.GetVal("pattern", "txt")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		GBK2UTF8.AllGBKs2UTF8(root, pattern)
	case "telegraph":
		links := util.GetVal("Telegraph", "links")
		urls := util.ReadByLine(links)
		slog.Debug(fmt.Sprint(urls))
		for _, uri := range urls {
			telegraph.GetAndDownload(uri)
		}
		_, filename, _, _ := runtime.Caller(0)
		processImage.ProcessAllImages(path.Dir(filename), "jpg", "0")
	case "list":
		root = util.GetVal("root", "folder")
		for index, folder := range GetAllFolder.List(root) {
			for idx, file := range GetFileInfo.GetAllFiles(folder) {
				fmt.Printf("读取第%d/%d个文件夹的第%d/%d个文件:%s\n", index+1, len(folder), idx+1, len(file), file)
			}
		}
	case "v2a":
		root = util.GetVal("root", "video")
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		threads = util.GetVal("thread", "threads")
		processVideo.AllVideos2Audio(root, pattern, threads)
	case "aspect":
		root = util.GetVal("root", "video")
		pattern = util.GetVal("pattern", "video")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		threads = util.GetVal("thread", "threads")
		processVideo.FixAll4x3s(root, pattern, threads)
	case "louder":
		root = util.GetVal("root", "louder")
		pattern = util.GetVal("pattern", "audio")
		pattern = strings.Join([]string{pattern, strings.ToUpper(pattern)}, ";")
		processAudio.LouderAllAudios(root, pattern)
		slog.Debug("开始有声小说增大电平处理", slog.String("根目录", root))
	case "ytdlp":
		go runNumGoroutineMonitor()
		list := util.GetVal("yt-dlp", "links")
		ytdlp.Ytdlp(list)
	default:
		fmt.Println("参数错误")
	}
	end = time.Now()
}

/*
程序结束后发送电子邮件通知,可选追加内容
*/
func sendEmail(start, end time.Time, ss ...string) {
	i := new(alert.Info)
	i.SetUsername(util.GetVal("email", "username"))
	i.SetPassword(util.GetVal("email", "password"))
	i.SetTo(strings.Split(util.GetVal("email", "tos"), ";"))
	i.SetFrom(util.GetVal("email", "from"))
	i.SetHost(alert.NetEase.SMTP)
	i.SetPort(alert.NetEase.SMTPProt)
	i.SetSubject(strings.Join([]string{"AllInOne", util.GetVal("main", "mission"), "任务完成"}, ":"))
	text := strings.Join([]string{start.Format("任务开始时间 2006年01月02日 15:04:05"), end.Format("任务结束时间 2006年01月02日 15:04:05"), fmt.Sprintf("任务用时%.3f分", end.Sub(start).Minutes())}, "<br>")
	i.SetText(text)
	for _, s := range ss {
		i.AppendText(s)
	}
	alert.Send(i)
}

/*
等待程序开始的循环函数
*/
func startOn(t string) {
	for true {
		now := time.Now().Local().Format("15")
		if t == now {
			return
		} else {
			slog.Warn("still alive", slog.Any("time", now), slog.String("target", t))
			time.Sleep(30 * time.Minute)
		}
	}
}

/*
初始化数据库和数据表
*/

/*
设置程序运行的日志等级
*/
func setLog() {
	var opt slog.HandlerOptions
	level := util.GetVal("log", "level")
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

/*
程序运行结束后生成新的二进制可执行文件
*/
func final() {
	if runtime.GOOS == "darwin" {
		slog.Info("M1重新生成预编译文件")
		if out, err := exec.Command("zsh", "-c", "build.sh").CombinedOutput(); err != nil {
			slog.Warn("程序结束后重新编译失败")
		} else {
			slog.Debug("编译新版本二进制文件", slog.String("输出", string(out)))
		}
	}
	if runtime.GOOS == "linux" {
		slog.Info("linux64重新生成预编译文件")
		if out, err := exec.Command("bash", "-c", "build.sh").CombinedOutput(); err != nil {
			slog.Warn("程序结束后重新编译失败")
		} else {
			slog.Debug("编译新版本二进制文件", slog.String("输出", string(out)))
		}
	}
	if runtime.GOOS == "android" {
		slog.Info("android重新生成预编译文件")
		if out, err := exec.Command("bash", "-c", "build.sh").CombinedOutput(); err != nil {
			slog.Warn("程序结束后重新编译失败")
		} else {
			slog.Debug("编译新版本二进制文件", slog.String("输出", string(out)))
		}
	}
}

/*
runNumGoroutineMonitor 协程数量监控
*/
func runNumGoroutineMonitor() {
	slog.Info(fmt.Sprintf("程序初始协程数量->%d\n", runtime.NumGoroutine()))
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Printf("协程数量->%d\n", runtime.NumGoroutine())
		}
	}
}
