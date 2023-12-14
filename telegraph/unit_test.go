package telegraph

import (
	"processAll/util"
	"testing"
)

func TestUnit(t *testing.T) {
	f := "exam.html"
	title, srcs := Parse(ReadHtml(f))
	DownloadSrc(title, srcs)
	for _, v := range srcs {
		t.Log(v)
	}
}
func TestParse(t *testing.T) {
	f := "exam.html"
	title, srcs := Parse(ReadHtml(f))
	t.Log(title, srcs)
}

func TestFromWeb(t *testing.T) {
	uri := "https://telegra.ph/%E7"
	web, err := GetWeb(uri)
	if err != nil {
		return
	}
	title, srcs := ParseWeb(web)
	DownloadSrc(title, srcs)
}
func TestFromFile(t *testing.T) {
	fname := "/Users/zen/github/processAVIWithXorm/telegraph/exam.html"
	title, srcs := Parse(ReadHtml(fname))
	DownloadSrc(title, srcs)
}
func TestFromWebs(t *testing.T) {
	urls := util.ReadByLine("/Users/zen/github/processAVIWithXorm/telegraph/list.txt")
	for _, uri := range urls {
		web, err := GetWeb(uri)
		if err != nil {
			return
		}
		title, srcs := Parse(ReadHtml(web))
		DownloadSrc(title, srcs)
	}

}
