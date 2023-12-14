package model

import (
	"log/slog"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

type AV struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`
	TaskId     int64     `xorm:"comment('任务id') INT" json:"TaskId"`
	Entry      []byte    `xorm:"comment('原始entry文件') Text" json:"Entry"`
	FinalName  string    `xorm:"comment('最终生成视频的全路径') VARCHAR(255)" json:"FinalName"`
	Video      string    `xorm:"comment('视频文件的绝对路径') VARCHAR(255)" json:"VideoName"`
	Audio      string    `xorm:"comment('音频文件的绝对路径') VARCHAR(255)" json:"AudioName"`
	FileName   string    `xorm:"comment('最终文件名') VARCHAR(255)" json:"FileName"`
	DelName    string    `xorm:"comment('标记删除的目录') VARCHAR(255)" json:"DelName"`
	FileSize   float64   `xorm:"comment('视频文件大小') Double" json:"FileSize"`
	MD5        string    `xorm:"comment('视频MD5') VARCHAR(255)" json:"MD5"`
	Code       string    `xorm:"comment('视频编码') VARCHAR(255)" json:"Code"`
	Width      int       `xorm:"comment('视频宽度') INT" json:"Width"`
	Height     int       `xorm:"comment('视频高度') INT" json:"Height"`
	VTag       string    `xorm:"comment('视频标签') VARCHAR(255)" json:"VTag"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncAV() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(AV))
		if err != nil {
			slog.Error("同步哔哩哔哩数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步哔哩哔哩数据表成功")
		}
	}
}

func (av *AV) GetAll() (AVs []AV, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(AVs)
		if err != nil {
			return nil, err
		}
		return AVs, nil
	}
	return nil, nil
}

func (av *AV) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(av)
	}
	return 0, nil
}

func (av *AV) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(av.Id).Get(av)
	}
	return false, nil
}

func (av *AV) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(av.Id).Update(av)
	}
	return 0, nil
}
func (av *AV) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(av)
	}
	return 0, nil
}
