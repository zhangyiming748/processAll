package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Ytdlp struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	URL        string    `xorm:"comment('下载链接') VARCHAR(255)" json:"url"`
	Status     string    `xorm:"comment('下载状态') VARCHAR(255)" json:"Status"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncYtdlp() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Ytdlp))
		if err != nil {
			slog.Error("同步数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步数据表成功")
		}
	}
}

func (y *Ytdlp) GetAll() (ytdlps []Ytdlp, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(ytdlps)
		if err != nil {
			return nil, err
		}
		return ytdlps, nil
	}
	return nil, nil
}
func (y *Ytdlp) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(y)
	}
	return 0, nil
}
func (y *Ytdlp) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(y.Id).Get(y)
	}
	return false, nil
}
func (y *Ytdlp) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(y.Id).Update(y)
	}
	return 0, nil
}

func (y *Ytdlp) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(y)
	}
	return 0, nil
}
