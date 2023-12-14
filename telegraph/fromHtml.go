package telegraph

import (
	"log/slog"
	"processAll/soup"
	"processAll/util"
)

func GetWeb(html string) (string, error) {
	// html := util.GetVal("Telegraph", "url")
	proxy := util.GetVal("Telegraph", "proxy")
	if proxy == "" {
		proxy = "http://127.0.0.1:8889"
	}
	withProxy, err := soup.GetWithProxy(html, proxy)
	if err != nil {
		slog.Warn("Get 失败", slog.Any("错误内容", err))
		return "", err
	} else {
		slog.Debug("Get 成功", slog.String("网页内容", withProxy))
		return withProxy, nil
	}
}
