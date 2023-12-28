package telegraph

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"processAll/soup"
	"processAll/util"
	"strconv"
	"strings"
)

/*
读取html文件
*/
func ReadHtml(fname string) string {
	file, err := os.ReadFile(fname)
	if err != nil {
		return ""
	}
	return string(file)
}

/*
解析html文件
*/
func Parse(html string) (string, []string) {
	var srcs []string
	doc := soup.HTMLParse(html)
	title := doc.Find("h1").Text()
	title = replace(title)
	slog.Debug("获取并替换标题", slog.String("标题", title))
	imgs := doc.FindAll("img")
	for _, img := range imgs {
		src := img.Attrs()["src"]
		if strings.Contains(src, "=") {
			src = strings.Split(src, "=")[1]
		}
		src = strings.Replace(src, "https://", "http://", -1)
		srcs = append(srcs, src)
	}
	return title, srcs
}

/*
解析html网站
*/
func ParseWeb(html string) (string, []string) {

	var srcs []string
	doc := soup.HTMLParse(html)
	div := doc.Find("div", "class", "ql-editor")
	title := div.Find("h1").Text()
	title = strings.Replace(title, "\n", "", -1)
	title = strings.Replace(title, "［", "", -1)
	title = strings.Replace(title, "］", "", -1)
	title = strings.Replace(title, "\u00A0", "", -1)
	title = strings.Replace(title, "，", "", -1)
	title = strings.Replace(title, "《", "", -1)
	title = strings.Replace(title, "》", "", -1)
	title = strings.Replace(title, " ", "", -1)
	title = strings.Replace(title, "[", "", -1)
	title = strings.Replace(title, "]", "", -1)
	title = strings.Replace(title, "（", "", -1)
	title = strings.Replace(title, "）", "", -1)
	for strings.Contains(title, " ") {
		title = strings.Replace(title, " ", "", -1)
	}
	slog.Debug("获取并替换标题", slog.String("标题", title))
	imgs := doc.FindAll("img")
	for _, img := range imgs {
		src := img.Attrs()["src"]
		if strings.Contains(src, "=") {
			src = strings.Split(src, "=")[1]
		}
		src = strings.Replace(src, "https://", "http://", -1)
		srcs = append(srcs, src)
	}
	return title, srcs
}

/*
使用wget下载
*/
func DownloadSrc(title string, images []string) {
	total := len(images)
	success := 0
	for index, image := range images {
		fname := strings.Join([]string{strconv.Itoa(index + 1), "jpg"}, ".")
		if strings.HasPrefix(image, "/file") {
			image = strings.Join([]string{"http://telegra.ph", image}, "")
		}

		//title = strings.Replace(title, "", "", -1)

		os.Mkdir(title, 0777)
		dir := strings.Join([]string{title, fname}, string(os.PathSeparator))
		//"wget -e use_proxy=yes -e http_proxy=127.0.0.1:8889 -e https_proxy=127.0.0.1:8889"
		cmd := exec.Command("wget", "-e", "use_proxy=yes", "-e", "http_proxy=127.0.0.1:8889", "-e", "https_proxy=127.0.0.1:8889", "--no-check-certificate", "--tries=0", "--continue", "-O", dir, image)
		slog.Debug("wget下载前", slog.String("生成的命令", fmt.Sprint(cmd)))
		shName := strings.Join([]string{title, "sh"}, ".")
		file, err := os.OpenFile(shName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			return
		}
		defer file.Close()
		file.WriteString(fmt.Sprint(cmd))
		file.WriteString("\n")

		//output, err := cmd.CombinedOutput()
		err = util.ExecCommand(cmd)
		if err != nil {
			slog.Warn("跳过", slog.Any("当前下载文件出错", err), slog.String("文件名", fname))
			fmt.Printf("出错的命令\n%s\n", cmd)
			continue
		} else {
			//slog.Debug("下载命令结束", slog.String("命令返回", string(output)))
			success++
		}
	}
	slog.Info("全部下载完毕", slog.Int("共有文件", total), slog.Int("成功下载", success))
}
