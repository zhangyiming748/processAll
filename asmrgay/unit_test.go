package asmrgay

import (
	"processAll/util"
	"testing"
)

func TestDecoder(t *testing.T) {
	ans := decoder("http://www.baidu.com/s?wd=自由度")
	t.Log(ans)
}
func TestUnit(t *testing.T) {
	html := ReadHtml("C:\\Users\\zen\\go\\src\\ProcessAVI\\asmrgay\\exam.html")
	srcs := Parse(html)
	//for _, src := range srcs {
	//	t.Log(src)
	//}
	util.WriteByLine("C:\\Users\\zen\\go\\src\\ProcessAVI\\asmrgay\\exam.txt", srcs)
}
func TestEncoder(t *testing.T) {
	ret := encoder("https://www.asmrgay.com/d/asmr/%E4%B8%AD%E6%96%87%E9%9F%B3%E5%A3%B0/%E6%9D%8F%E5%90%A7%E9%AA%9A%E9%BA%A6/%E9%AA%9A%E9%BA%A6/001%E3%80%8A%E5%93%8E%E5%91%80%E6%88%91%E6%93%8D%E5%A4%A7%E8%89%B2%E7%8B%BC%E3%80%8B%E3%80%90%E5%B0%8F%E8%8E%AB%E3%80%91.mp3?sign=TsXbjWWtYVpdfY96jDbkqmaJEhzMQYWGzLLKu9vL5YI=:1772008546\n")
	t.Log(ret)
}
