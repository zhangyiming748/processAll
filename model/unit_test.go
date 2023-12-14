package model

import (
	"fmt"
	"processAll/storage/mysql"
	"testing"
)

func init() {
	mysql.SetEngine()
	SyncSave()
	SyncYtdlp()
}
func TestMysql(t *testing.T) {
	s := new(Save)
	all, err := s.SumSaveAll()
	if err != nil {
		return
	} else {
		fmt.Printf("节省的空间:%v GB\n", all)
	}
}
