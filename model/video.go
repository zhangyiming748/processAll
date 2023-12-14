package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type Video struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	TaskId     int64     `xorm:"not null  comment('任务id') INT" json:"TaskId"`
	Frame      string    `xorm:"comment('预计文件帧数') VARCHAR(255)" json:"frame"`
	SrcName    string    `xorm:"comment('源文件名') VARCHAR(255)" json:"SrcName"`
	DstName    string    `xorm:"comment('目标文件名') VARCHAR(255)" json:"DstName"`
	SrcType    string    `xorm:"comment('源文件类型') VARCHAR(255)" json:"SrcType"`
	DstType    string    `xorm:"comment('目标文件类型') VARCHAR(255)" json:"DstTypeType"`
	SrcSize    int64     `xorm:"comment('源文件大小') BIGINT" json:"SrcSize"`
	DstSize    int64     `xorm:"comment('目标文件大小') BIGINT" json:"DstSize"`
	SavedSize  float64   `xorm:"comment('节省的磁盘空间(MB)') Float" json:"SavedSize"`
	SrcMD5     string    `xorm:"comment('源文件MD5') VARCHAR(255)" json:"SrcMD5"`
	DstMD5     string    `xorm:"comment('目标文件MD5') VARCHAR(255)" json:"DstMD5"`
	SrcCode    string    `xorm:"comment('源编码') VARCHAR(255)" json:"SrcCode"`
	DstCode    string    `xorm:"comment('目标编码') VARCHAR(255)" json:"DstCode"`
	SrcWidth   int       `xorm:"comment('源宽度') INT" json:"SrcWidth"`
	DstWidth   int       `xorm:"comment('目标宽度') INT" json:"DstWidth"`
	SrcHeight  int       `xorm:"comment('源高度') INT" json:"SrcHeight"`
	DstHeight  int       `xorm:"comment('目标高度') INT" json:"DstHeight"`
	SrcVTag    string    `xorm:"comment('源视频标签') VARCHAR(255)" json:"SrcVTag"`
	DstVTag    string    `xorm:"comment('目标视频标签') VARCHAR(255)" json:"DstVTag"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncVideo() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Video))
		if err != nil {
			slog.Error("同步视频数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步视频数据表成功")
		}
	}
}

func (v *Video) GetAll() (videos []Video, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(videos)
		if err != nil {
			return nil, err
		}
		return videos, nil
	}
	return nil, nil
}
func (v *Video) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(v)
	}
	return 0, nil
}
func (v *Video) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(v.Id).Get(v)
	}
	return false, nil
}
func (v *Video) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(v.Id).Update(v)
	}
	return 0, nil
}
func (v *Video) UpdateByTaskId() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Where("task_id = ?", v.TaskId).Update(v)
	}
	return 0, nil
}
func (v *Video) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(v)
	}
	return 0, nil
}
