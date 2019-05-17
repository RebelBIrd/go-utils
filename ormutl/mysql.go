package ormutl

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/strutl"
)

//mysql驱动器，带yaml配置项
type MysqlConf struct {
	Url      string `yaml:"url"`      // 数据库地址
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Database string `yaml:"database"` // 数据库名
	*xorm.Engine                      // 数据库引擎
}

var mEngine *MysqlConf

func GetEngine() *MysqlConf {
	if mEngine == nil {
		logutl.Error("Mysql doesn't init.")
		return nil
	}
	return mEngine
}

func InitMysql(conf MysqlConf) {
	if mEngine == nil {
		dbConf := strutl.ConnString(conf.Username, ":", conf.Password, "@tcp(", conf.Url, ")/", conf.Database, "?charset=utf8&parseTime=True&Local")
		var err error
		conf.Engine, err = xorm.NewEngine("mysql", dbConf)
		if err != nil {
			mEngine = nil
			logutl.Error(err.Error())
		} else {
			mEngine = &conf
			mEngine.SetMapper(core.SameMapper{})
			mEngine.SetMaxOpenConns(25)
			mEngine.SetMaxIdleConns(5)
		}
	}
}

func (mEngine *MysqlConf) InitTables(initBlock func(interface{}), beans ...interface{}) {
	for _, table := range beans {
		if isExist, err := mEngine.IsTableExist(table); err != nil || !isExist {
			_ = mEngine.CreateTables(table)
			if initBlock != nil {
				initBlock(table)
			}
		} else {
			_ = mEngine.Sync(table)
		}
	}
}
