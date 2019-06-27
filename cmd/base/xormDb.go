package base

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

var Db *xorm.Engine

func init() {
	driverName := "mysql"
	dataSourceName := ""
	logrus.Info("初始化数据库...")
	dbEngine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = dbEngine.DB().PingContext(ctx)
	if err != nil {
		logrus.Fatal("数据库信息错误！")
	}
	logrus.Info("数据库连接成功!")

	//设置数据库最大连接数
	dbEngine.SetMaxOpenConns(2)
	//是否显示SQL语句
	dbEngine.ShowSQL(true)
	//开启debug日志
	//dbEngine.Logger().SetLevel(core.LOG_DEBUG)
	//开启自动建表
	//err = dbEngine.Sync2(new(model.User))
	if err != nil {
		logrus.Error(err.Error())
	}

	Db = dbEngine
}
