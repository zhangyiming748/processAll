package processVideo

import (
	"fmt"
	"path"
	"testing"
)

func TestDir(t *testing.T) {
	fp := "/Users/zen/Downloads/Telegram Desktop/水岛津实/33.mp4"
	ret := path.Dir(fp)
	t.Log(ret)
}

func TestProcessAllH265(t *testing.T) {
	root := "/Users/zen/Downloads/telegram/cache"
	pattern := "mp4"
	ProcessAllVideos2H265(root, pattern, "3")
}
func TestPanic(t *testing.T) {
	if err := recover(); err != nil {
		fmt.Println("有panic")
	}
	go func() {
		panic("panic!")
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func TestGetOutOfH265(t *testing.T) {
	GetOutOfH265("/Volumes/volume/未整理", "mp4")
}
