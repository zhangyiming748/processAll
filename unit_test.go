package main

import (
	"log/slog"
	"processAll/model"
	"processAll/storage/mysql"
	"testing"
)

func TestUnit(t *testing.T) {
	mysql.SetEngine()
	i := new(model.Image)
	all, err := i.Sum()
	if err != nil {
		return
	}
	slog.Debug("all image", slog.Int64("共处理的图片数", all))
}
