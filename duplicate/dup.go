package duplicate

import (
	"log/slog"
	"processAll/util"
)

func DuplicateFromFile(fp string) {
	m := make(map[string]bool)
	src := util.ReadByLine(fp)
	dst := make([]string, 0)
	for _, v := range src {
		if _, ok := m[v]; ok {
			slog.Warn("skip", slog.String("文件名", v))
			continue
		} else {
			dst = append(dst, v)
			m[v] = true
		}
	}
	util.WriteByLine("dup.sh", dst)
}
