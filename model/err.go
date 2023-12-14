package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Err struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	FullPath   string    `xorm:"comment('出错文件的全路径') VARCHAR(255)" json:"FullPath"`
	Reason     string    `xorm:"comment('错误文本') VARCHAR(255)" json:"Reason"`
	TTY        string    `xorm:"comment('控制台内容') TEXT" json:"tty"`
	Warn       string    `xorm:"comment('日志内容') TEXT" json:"content"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncErr() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Err))
		if err != nil {
			slog.Error("同步错误数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步错误数据表成功")
		}
	}
}

func (e *Err) GetAll() (es []Err, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(es)
		if err != nil {
			return nil, err
		}
		return es, nil
	}
	return nil, nil
}

func (e *Err) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(e)
	}
	return 0, nil
}
func (e *Err) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(e)
	}
	return 0, nil
}
