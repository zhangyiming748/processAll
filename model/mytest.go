package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Custom struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	FileName   string    `xorm:"comment('文件名') VARCHAR(255)" json:"SrcName"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncCustom() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Custom))
		if err != nil {
			slog.Error("同步测试数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步测试数据表成功")
		}
	}
}

func (c *Custom) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(c)
	}
	return 0, nil
}
func (c *Custom) InsertAll(cs []*Custom) (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(&cs)
	}
	return 0, nil
}
