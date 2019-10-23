package ormutl

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/qinyuanmao/go-utils/fileutl"
	"github.com/qinyuanmao/go-utils/logutl"
	"strings"
)

type SqliteConf struct {
	DbPath string
	DbName string
	OrmEngine
}

func (engine *SqliteConf) InitTables(initBlock func(interface{}), beans ...interface{}) {
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


func (conf *SqliteConf) InitSqlite() {
	if !strings.HasSuffix(conf.DbPath, "/") {
		conf.DbPath += "/"
	}
	fileutl.PathExistOrCreate(conf.DbPath)
	if conf.Engine == nil {
		var err error
		conf.Engine, err = xorm.NewEngine("sqlite", conf.DbPath + conf.DbName)
		if err != nil {
			logutl.Error(err.Error())
		} else {
			conf.SetMapper(core.SameMapper{})
			conf.SetMaxOpenConns(25)
			conf.SetMaxIdleConns(5)
		}
	}
}