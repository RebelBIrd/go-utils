package gormutl

import (
	"github.com/jinzhu/gorm"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/strutl"
)

type PostgresConf struct {
	*Engine
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (conf *PostgresConf) InitPostgres() (err error) {
	sqlConfStr := strutl.ConnString("host=", conf.URL, " user=", conf.Username, " password=", conf.Password, " dbName=", conf.Database, " sslmode=disable")
	if conf.Engine.DB, err = gorm.Open("postgres", sqlConfStr); err != nil {
		logutl.Error(err)
	}
	return
}
