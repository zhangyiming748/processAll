package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type File struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	Name       string    `xorm:"comment('文件名') VARCHAR(255)" json:"name"`
	Type       string    `xorm:"comment('文件类型') VARCHAR(255)" json:"type"`
	Size       int64     `xorm:"comment('文件大小') BIGINT" json:"size"`
	MD5        string    `xorm:"comment('文件MD5') VARCHAR(255)" json:"MD5"`
	IsVideo    bool      `xorm:"comment('是否为视频文件') BOOL" json:"isvideo"`
	Frame      string    `xorm:"comment('文件大致帧数') VARCHAR(255)" json:"frame"`
	Code       string    `xorm:"comment('编码') VARCHAR(255)" json:"code"`
	Width      int       `xorm:"comment('宽度') INT" json:"width"`
	Height     int       `xorm:"comment('高度') INT" json:"height"`
	VTag       string    `xorm:"comment('视频标签') VARCHAR(255)" json:"vtag"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncFile() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(File))
		if err != nil {
			slog.Error("同步文件数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步文件数据表成功")
		}
	}
}

func (f *File) GetAll() (files []File, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(files)
		if err != nil {
			return nil, err
		}
		return files, nil
	}
	return nil, nil
}

func (f *File) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(f)
	}
	return 0, nil
}

func (f *File) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(f.Id).Get(f)
	}
	return false, nil
}

func (f *File) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(f.Id).Update(f)
	}
	return 0, nil
}

func (f *File) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(f)
	}
	return 0, nil
}
func (f *File) InsertAll(fs []*File) (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(&fs)
	}
	return 0, nil
}
