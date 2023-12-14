package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Audio struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`
	TaskId     int64     `xorm:"not null  comment('任务id') INT" json:"TaskId"`
	SrcName    string    `xorm:"comment('源文件名') VARCHAR(255)" json:"SrcName"`
	DstName    string    `xorm:"comment('目标文件名') VARCHAR(255)" json:"DstName"`
	SrcType    string    `xorm:"comment('源文件类型') VARCHAR(255)" json:"SrcType"`
	DstType    string    `xorm:"comment('目标文件类型') VARCHAR(255)" json:"DstTypeType"`
	SrcSize    int64     `xorm:"comment('源文件大小') BIGINT" json:"SrcSize"`
	DstSize    int64     `xorm:"comment('目标文件大小') BIGINT" json:"DstSize"`
	SrcMD5     string    `xorm:"comment('源文件MD5') VARCHAR(255)" json:"SrcMD5"`
	DstMD5     string    `xorm:"comment('目标文件MD5') VARCHAR(255)" json:"DstMD5"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncAudio() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Audio))
		if err != nil {
			slog.Error("同步音频数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步音频数据表成功")
		}
	}
}

func (a *Audio) GetAll() (audios []Audio, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(audios)
		if err != nil {
			return nil, err
		}
		return audios, nil
	}
	return nil, nil
}

func (a *Audio) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(a)
	}
	return 0, nil
}

func (a *Audio) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(a.Id).Get(a)
	}
	return false, nil
}

func (a *Audio) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(a.Id).Update(a)
	}
	return 0, nil
}
func (a *Audio) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(a)
	}
	return 0, nil
}
