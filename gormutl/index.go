package gormutl

import (
	"github.com/jinzhu/gorm"
)

type Engine struct {
	*gorm.DB
}

func (engine *Engine) CloseDB() error {
	return engine.Close()
}

func (engine *Engine) InitTable(initBlock func(interface{}), tables ...interface{}) {
	for _, table := range tables {
		if !engine.HasTable(table) {
			engine.CreateTable(&table)
			initBlock(table)
		} else {
			engine.AutoMigrate(table)
		}
	}
}
