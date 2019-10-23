package ormutl

import (
	"github.com/go-xorm/xorm"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/sliceutl"
	"log"
	"reflect"
	"strings"
)

type OrmEngine struct {
	*xorm.Engine
}

func (engine *OrmEngine) Exist(id interface{}, model interface{}) bool {
	exist, err := engine.Table(model).ID(id).NoCache().NoAutoCondition().Exist(model)
	if err != nil {
		logutl.Error(err.Error())
	}
	return exist
}

func (engine *OrmEngine) Save(id interface{}, model interface{}) (err error) {
	if _, err = engine.UseBool(engine.GetBoolField(model)...).Insert(model); err != nil {
		logutl.Error(err)
		return err
	}
	if _, err = engine.ID(id).UseBool(engine.GetBoolField(model)...).NoAutoCondition().Get(model); err != nil {
		logutl.Error(err)
		return err
	}
	return
}

func (engine *OrmEngine) SaveArray(m interface{}) (err error) {
	models := sliceutl.InterfaceSlice(m)
	for i := 0; i < len(models); i = i + 99 {
		if i+99 >= len(models) {
			_, err = engine.Insert(models[i:])
			if err != nil {
				logutl.Error(err)
				var session = engine.NewSession()
				defer session.Close()
				_ = session.Begin()
				for j := i; j < len(models); j++ {
					if _, err := engine.Insert(models[j]); err != nil {
						logutl.Error(err)
						_ = session.Rollback()
					}
				}
				if err = session.Commit(); err != nil {
					logutl.Error(err)
				}
			}
		} else {
			_, err = engine.Insert(models[i : i+99])
			if err != nil {
				logutl.Error(err)
				var session = engine.NewSession()
				defer session.Close()
				_ = session.Begin()
				for j := i; j < i+99; j++ {
					if _, err := session.Insert(models[j]); err != nil {
						logutl.Error(err)
						_ = session.Rollback()
					}
				}
				if err = session.Commit(); err != nil {
					logutl.Error(err)
				}
			}
		}
	}
	return
}

func (engine *OrmEngine) Update(id interface{}, model interface{}) (err error) {
	if engine.Exist(id, model) {
		if _, err = engine.NoAutoCondition().UseBool(engine.GetBoolField(model)...).ID(id).Update(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
	} else {
		if _, err = engine.Insert(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
	}
	if _, err = engine.NoAutoCondition().ID(id).Get(model); err != nil {
		logutl.Error(err.Error())
		return err
	}
	return
}

func (engine *OrmEngine) Delete(id interface{}, model interface{}) (err error) {
	if engine.Exist(id, model) {
		if _, err = engine.Table(model).ID(id).Delete(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
		return nil
	} else {
		return nil
	}
}

func (engine *OrmEngine) GetById(id interface{}, model interface{}) (err error) {
	if _, err = engine.Table(model).ID(id).NoAutoCondition().Get(model); err != nil {
		logutl.Error(err.Error())
		return err
	}
	return nil
}

func (engine *OrmEngine) Remove(tableName string, where []string, values []interface{}) (err error) {
	w := strings.Join(where, " And ")
	if _, err := engine.Exec("Delete From `"+tableName+"` Where "+w, values); err != nil {
		logutl.Error(err)
	}
	return
}

func (engine *OrmEngine) GetBoolField(model interface{}) (fields []string) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Type.Kind() == reflect.Bool {
			result = append(result, t.Field(i).Name)
		}
	}
	return result
}
