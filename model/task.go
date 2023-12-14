package model

import (
	"log/slog"
	"processAll/information"
	"processAll/storage/mysql"
	"processAll/util"
	"time"
)

/*
主任务表
*/

type Task struct {
	Id          int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`
	TaskType    string    `xorm:"comment('任务类型') VARCHAR(255)" json:"TaskType"`
	MachineType string    `xorm:"comment('执行的机器类型') VARCHAR(255)" json:"MachineType"`
	UpdateTime  time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime  time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime  time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncTask() {
	err := mysql.GetSession().Sync2(new(Task))
	if err != nil {
		slog.Error("同步主任务数据表出错", slog.Any("错误原文", err))
		return
	} else {
		slog.Debug("同步主任务数据表成功")
	}
}
func (t *Task) CreateOne() (int64, error) {
	t.MachineType = information.GetMachineInfo()
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(t)
	}
	return 0, nil
}
