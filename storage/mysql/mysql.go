package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"os"
	"processAll/util"
	"strings"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var MyEngine *xorm.Engine

func SetEngine() {
	user := util.GetVal("mysql", "user")
	ip := util.GetVal("mysql", "ip")
	port := util.GetVal("mysql", "port")
	passwd := util.GetVal("mysql", "passwd")
	database := util.GetVal("mysql", "database")

	uri := strings.Join([]string{ip, port}, ":")
	src := strings.Join([]string{user, ":", passwd, "@tcp(", uri, ")/", database, "?charset=utf8mb4"}, "")
	slog.Debug("数据库链接", slog.String("参数", src))

	MyEngine, _ = xorm.NewEngine("mysql", src)
	ch := make(chan struct{}, 1)
	go success(ch)
	select {
	case <-ch:
		MyEngine.SetMapper(names.GonicMapper{})
		MyEngine.SetTZDatabase(time.Local)
		if util.GetVal("log", "level") == "Debug" {
			MyEngine.SetLogLevel(log.LOG_DEBUG)
			//MyEngine.ShowSQL(true)
		}
	case <-time.After(5 * time.Second):
		slog.Warn("数据库连接超时")
		if err := util.SetVal("mysql", "switch", "off"); err != nil {
			slog.Warn("数据库连接失败后关闭失败,退出程序")
			os.Exit(-1)
		}
	}
}

func GetSession() *xorm.Session {
	return MyEngine.NewSession()
}
func success(ch chan struct{}) {
	if err := MyEngine.Ping(); err != nil {
		slog.Error("创建数据库引擎失败", slog.Any("错误信息", err))
	} else {
		slog.Debug("创建数据库引擎成功", slog.Any("MyEngine", MyEngine))
		ch <- struct{}{}
	}
}
