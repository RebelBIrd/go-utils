package ormutl

import (
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/pageutl"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"time"
)

type MongoConf struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	session  *mgo.Session
}

func (conf *MongoConf) InitMongoDB() {
	var dialInfo = &mgo.DialInfo{
		Addrs:     []string{conf.Url},
		Direct:    false,
		Timeout:   time.Minute,
		FailFast:  false,
		Database:  conf.Database,
		Username:  conf.Username,
		Password:  conf.Password,
		PoolLimit: 4096,
	}
	var err error
	if conf.session, err = mgo.DialWithInfo(dialInfo); err != nil {
		logutl.Error(err)
	} else {
		conf.session.SetMode(mgo.Monotonic, true)
	}
}

func (conf *MongoConf) NewSession(name string) (*mgo.Session, *mgo.Collection) {
	session := conf.session.Copy()
	col := session.DB(conf.Database).C(name)
	conf.session.SetMode(mgo.Monotonic, true)
	return session, col
}

func (conf *MongoConf) Close() {
	conf.session.Close()
}

func (conf *MongoConf) Save(model interface{}) (err error) {
	t := reflect.TypeOf(model)
	s, c := conf.NewSession(t.Name())
	defer s.Close()
	if err = c.Insert(model); err != nil {
		logutl.Error(err)
	}
	return err
}

func (conf *MongoConf) Upsert(whereKey bson.M, model interface{}) (err error) {
	t := reflect.TypeOf(model)
	s, c := conf.NewSession(t.Name())
	defer s.Close()
	if _, err = c.Upsert(whereKey, model); err != nil {
		logutl.Error(err)
	}
	return
}

func (conf *MongoConf) Get(query, whereKey bson.M, model interface{}) (err error) {
	t := reflect.TypeOf(model)
	s, c := conf.NewSession(t.Name())
	defer s.Close()
	if err = c.Find(query).Select(whereKey).One(model); err != nil {
		logutl.Error(err)
	}
	return
}

func (conf *MongoConf) Find(query, whereKey bson.M, collection string) (err error, model interface{}) {
	s, c := conf.NewSession(collection)
	defer s.Close()
	if err = c.Find(query).Select(whereKey).All(model); err != nil {
		logutl.Error(err)
	}
	return
}

func (conf *MongoConf) FindPage(query, whereKey bson.M, pageIndex, pageSize int, collection string) (total, count int, model interface{}) {
	s, c := conf.NewSession(collection)
	defer s.Close()
	if err := c.Find(query).Select(whereKey).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(model); err != nil {
		logutl.Error(err)
	}
	var err error
	if total, err = c.Find(query).Count(); err != nil {
		logutl.Error(err)
	}
	count = pageutl.GetPageCount(pageSize, total)
	return
}

func (conf *MongoConf) Remove(collection string, selector interface{}) error {
	ms, c := conf.NewSession(collection)
	defer ms.Close()
	if err := c.Remove(selector); err != nil {
		logutl.Error(err)
		return err
	}
	return nil
}

func (conf *MongoConf) RemoveAll(collection string, selector interface{}) error {
	ms, c := conf.NewSession(collection)
	defer ms.Close()
	if _, err := c.RemoveAll(selector); err != nil {
		logutl.Error(err)
		return err
	}
	return nil
}
