package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Louder struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`
	TaskId     int64     `xorm:"not null  comment('任务id') INT" json:"TaskId"`
	SrcName    string    `xorm:"comment('增大电平前文件名') VARCHAR(255)" json:"SrcName"`
	DstName    string    `xorm:"comment('增大电平后文件名') VARCHAR(255)" json:"DstName"`
	SrcType    string    `xorm:"comment('增大电平前文件类型') VARCHAR(255)" json:"SrcType"`
	DstType    string    `xorm:"comment('增大电平后文件类型') VARCHAR(255)" json:"DstTypeType"`
	SrcSize    int64     `xorm:"comment('增大电平前文件大小') BIGINT" json:"SrcSize"`
	DstSize    int64     `xorm:"comment('增大电平后文件大小') BIGINT" json:"DstSize"`
	SrcMD5     string    `xorm:"comment('增大电平前文件MD5') VARCHAR(255)" json:"SrcMD5"`
	DstMD5     string    `xorm:"comment('增大电平后文件MD5') VARCHAR(255)" json:"DstMD5"`
	Multiple   string    `xorm:"comment('增大电平倍数') VARCHAR(255)" json:"Multiple"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncLouder() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Louder))
		if err != nil {
			slog.Error("同步音频放大数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步音频放大数据表成功")
		}
	}
}

func (l *Louder) GetAll() (louders []Louder, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(louders)
		if err != nil {
			return nil, err
		}
		return louders, nil
	}
	return nil, nil
}

func (l *Louder) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(l)
	}
	return 0, nil
}

func (l *Louder) FindByTaskId() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Where("taskid = ?", l.TaskId).Get(l)
	}
	return false, nil
}

func (l *Louder) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(l.Id).Update(l)
	}
	return 0, nil
}

func (l *Louder) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(l)
	}
	return 0, nil
}
