package ormutl

import (
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/strutl"
	"github.com/qinyuanmao/go-utils/timeutl"
	"xorm.io/core"
	"xorm.io/xorm"
)

//mysql驱动器，带yaml配置项
type MysqlConf struct {
	Url       string `yaml:"url"`      // 数据库地址
	Username  string `yaml:"username"` // 数据库用户名
	Password  string `yaml:"password"` // 数据库密码
	Database  string `yaml:"database"` // 数据库名
	OrmEngine        // 数据库引擎
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
		dbConf := strutl.ConnString(conf.Username, ":", conf.Password, "@tcp(", conf.Url, ")/", conf.Database)
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

func (engine *MysqlConf) InitTables(initBlock func(interface{}), beans ...interface{}) {
	for _, table := range beans {
		if isExist, err := engine.IsTableExist(table); err != nil || !isExist {
			if err != nil {
				logutl.Error(err)
			}
			if err = engine.Sync2(table); err != nil {
				logutl.Error(err)
			}
			if initBlock != nil {
				initBlock(table)
			}
		} else {
			if err = engine.Sync2(table); err != nil {
				logutl.Error(err)
			}
		}
	}
}

func (conf *MysqlConf) InitMysql() {
	if conf.Engine == nil {
		dbConf := strutl.ConnString(conf.Username, ":", conf.Password, "@tcp(", conf.Url, ")/", conf.Database, "?charset=utf8&parseTime=True&Local")
		var err error
		conf.Engine, err = xorm.NewEngine("mysql", dbConf)
		if err != nil {
			conf.Engine = nil
			logutl.Error(err.Error())
		} else {
			conf.Engine.SetMapper(core.SameMapper{})
			conf.Engine.SetMaxOpenConns(25)
			conf.Engine.SetMaxIdleConns(5)
		}
	}
}

func (mEngine *MysqlConf) ResetToken(eventName, sqlString string, hours int) (err error) {
	eventName = strings.ReplaceAll(eventName, "-", "")
	if _, err = mEngine.Exec(`DROP EVENT IF EXISTS ` + eventName); err != nil {
		logutl.Error(err)
	} else if _, err = mEngine.Exec(`CREATE EVENT ` + eventName + ` on schedule at date_add(now(), interval ` + strconv.Itoa(hours) + ` hour) do ` + sqlString); err != nil {
		logutl.Error(err)
	}
	return
}

func (mEngine *MysqlConf) TimingEvent(eventName, sqlString string, doAt time.Time) (err error) {
	eventName = strings.ReplaceAll(eventName, "-", "")
	if _, err = mEngine.Exec(`DROP EVENT IF EXISTS ` + eventName); err != nil {
		logutl.Error(err)
	} else if _, err = mEngine.Exec(`CREATE EVENT ` + eventName + ` on schedule at TIMESTAMP '` + timeutl.GetTimeString2(doAt) + `' ON COMPLETION PRESERVE ENABLE do ` + sqlString); err != nil {
		logutl.Error(err)
	}
	return
}
