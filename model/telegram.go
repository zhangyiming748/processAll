package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Telegraph struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	Name       string    `xorm:"comment('文件名') VARCHAR(255)" json:"name"`
	Shell      string    `xorm:"comment('下载命令') VARCHAR(255)" json:"shell"`
	Url        string    `xorm:"comment('url')  VARCHAR(255)" json:"url"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncTelegraph() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Telegraph))
		if err != nil {
			slog.Error("同步电报图片数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步电报图片数据表成功")
		}
	}
}

func (t *Telegraph) GetAll() (telegraphs []Telegraph, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(telegraphs)
		if err != nil {
			return nil, err
		}
		return telegraphs, nil
	}
	return nil, nil
}

func (t *Telegraph) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(t)
	}
	return 0, nil
}

func (t *Telegraph) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(t.Id).Get(t)
	}
	return false, nil
}

func (t *Telegraph) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(t.Id).Update(t)
	}
	return 0, nil
}

func (t *Telegraph) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(t)
	}
	return 0, nil
}
