package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Speed struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`
	TaskId     int64     `xorm:"not null  comment('任务id') INT" json:"TaskId"`
	SrcName    string    `xorm:"comment('加速前文件名') VARCHAR(255)" json:"SrcName"`
	DstName    string    `xorm:"comment('加速后文件名') VARCHAR(255)" json:"DstName"`
	SrcType    string    `xorm:"comment('加速前文件类型') VARCHAR(255)" json:"SrcType"`
	DstType    string    `xorm:"comment('加速后文件类型') VARCHAR(255)" json:"DstTypeType"`
	SrcSize    int64     `xorm:"comment('加速前文件大小') BIGINT" json:"SrcSize"`
	DstSize    int64     `xorm:"comment('加速后文件大小') BIGINT" json:"DstSize"`
	SrcMD5     string    `xorm:"comment('加速前文件MD5') VARCHAR(255)" json:"SrcMD5"`
	DstMD5     string    `xorm:"comment('加速后文件MD5') VARCHAR(255)" json:"DstMD5"`
	Multiple   string    `xorm:"comment('加速倍数') VARCHAR(255)" json:"Multiple"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncSpeed() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Speed))
		if err != nil {
			slog.Error("同步音频加速数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步音频加速数据表成功")
		}
	}
}

func (s *Speed) GetAll() (speeds []Speed, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(speeds)
		if err != nil {
			return nil, err
		}
		return speeds, nil
	}
	return nil, nil
}

func (s *Speed) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(s)
	}
	return 0, nil
}

func (s *Speed) FindByTaskId() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Where("taskid = ?", s.TaskId).Get(s)
	}
	return false, nil
}

func (s *Speed) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(s.Id).Update(s)
	}
	return 0, nil
}

func (s *Speed) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(s)
	}
	return 0, nil
}
