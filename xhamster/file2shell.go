package xhamster

import (
	"processAll/util"
	"strings"
)

func File2Shell() {
	wget := []string{"#!/bin/bash"}
	links := util.ReadByLine("list.txt")
	for _, link := range links {
		prefix := "ytdlp --proxy http://172.26.0.1:8889"
		cmd := strings.Join([]string{prefix, link}, " ")
		wget = append(wget, cmd)
	}
	util.WriteByLine("list.sh", wget)
}
