package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Image struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	TaskId     int64     `xorm:"not null pk  comment('任务id') INT" json:"TaskId"`
	ToType     bool      `xorm:"comment('静态0fales｜动态1true')BOOL" json:"ToType"`
	SrcName    string    `xorm:"comment('源文件名') VARCHAR(255)" json:"SrcName"`
	DstName    string    `xorm:"comment('目标文件名') VARCHAR(255)" json:"DstName"`
	SrcType    string    `xorm:"comment('源文件类型') VARCHAR(255)" json:"SrcType"`
	DstType    string    `xorm:"comment('目标文件类型') VARCHAR(255)" json:"DstTypeType"`
	SrcSize    int64     `xorm:"comment('源文件大小') BIGINT" json:"SrcSize"`
	DstSize    int64     `xorm:"comment('目标文件大小') BIGINT" json:"DstSize"`
	SavedSize  float64   `xorm:"comment('节省的磁盘空间(MB)') Float" json:"SavedSize"`
	SrcMD5     string    `xorm:"comment('源文件MD5') VARCHAR(255)" json:"SrcMD5"`
	DstMD5     string    `xorm:"comment('目标文件MD5') VARCHAR(255)" json:"DstMD5"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncImage() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Image))
		if err != nil {
			slog.Error("同步图片数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步图片数据表成功")
		}
	}
}

func (i *Image) GetAll() (images []Image, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(images)
		if err != nil {
			return nil, err
		}
		return images, nil
	}
	return nil, nil
}

func (i *Image) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(i)
	}
	return 0, nil
}

func (i *Image) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(i.Id).Get(i)
	}
	return false, nil
}

func (i *Image) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(i.Id).Update(i)
	}
	return 0, nil
}
func (i *Image) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(i)
	}
	return 0, nil
}
