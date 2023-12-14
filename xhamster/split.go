package xhamster

import (
	"os"
	"strings"
)

func SplitLinks(fp string) {
	b, err := os.ReadFile(fp)
	if err != nil {
		return
	}
	c := string(b)
	cs := strings.Replace(c, "https", "\nhttps", -1)
	//cs = strings.Replace(c, "\n", "", 1)
	bs := []byte(cs)
	os.WriteFile("new.txt", bs, 0777)
}
