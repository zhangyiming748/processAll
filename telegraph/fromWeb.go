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
	"time"
)

func GetAndDownload(website string) {
	withProxy, err := soup.GetWithProxy(website, "http://127.0.0.1:8889")
	if err != nil {
		slog.Warn("Get 失败", slog.Any("错误内容", err))
		return
	}
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("出问题的网页", slog.String("网页原文", withProxy))
		}
	}()
	slog.Debug("Get 成功", slog.String("网页内容", withProxy))
	doc := soup.HTMLParse(withProxy)
	div := doc.Find("article")
	h1 := div.Find("h1")
	title := replace(h1.Text())
	slog.Debug("ql-editor节点", slog.Any("div", div))
	imgs := div.FindAll("img")
	for i, img := range imgs {
		src := img.Attrs()["src"]
		if strings.Contains(src, "telegra.ph") {
			src = strings.Replace(src, "https://", "http://", -1)
		} else if strings.Contains(src, "23img.com") {
			src = strings.Replace(src, "https://wsrv.nl/?url=", "", -1)
			src = strings.Replace(src, "https://", "http://", -1)
		} else {
			src = strings.Join([]string{"http://telegra.ph", src}, "")
		}
		fmt.Println(i + 1)
		slog.Debug("获取图片地址", slog.String("src", src), slog.String("所属标题", title))
		os.Mkdir(title, 0777)
		fname := strings.Join([]string{strconv.Itoa(i + 1), "jpg"}, ".")
		dir := strings.Join([]string{title, fname}, string(os.PathSeparator))
		cmd := exec.Command("wget", "-e", "use_proxy=yes", "-e", "http_proxy=127.0.0.1:8889", "-e", "https_proxy=127.0.0.1:8889", "--no-check-certificate", "--tries=0", "--continue", "-O", dir, src)
		err = util.ExecCommand(cmd)
		if err != nil {
			time.Sleep(time.Second)
			err = util.ExecCommand(cmd)
			if err != nil {
				time.Sleep(time.Second)
				err = util.ExecCommand(cmd)
				if err != nil {
					time.Sleep(time.Second)
					slog.Warn("下载失败", slog.String("文件地址", src), slog.String("下载命令", fmt.Sprint(cmd)))
				}
			}
		}
	}
}
