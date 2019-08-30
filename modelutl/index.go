package modelutl

import (
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/ormutl"
)

func Exist(id interface{}, model interface{}) bool {
	exist, err := ormutl.GetEngine().Table(model).ID(id).Exist()
	if err != nil {
		logutl.Error(err.Error())
	}
	return exist
}

func Update(id interface{}, model interface{}) (err error) {
	if Exist(id, model) {
		if _, err = ormutl.GetEngine().Table(model).ID(id).Update(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
		return nil
	} else {
		if _, err = ormutl.GetEngine().Table(model).ID(id).Insert(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
		_, _ = ormutl.GetEngine().Table(model).ID(id).Get(model)
		return nil
	}
}

func Delete(id interface{}, model interface{}) (err error) {
	if Exist(id, model) {
		if _, err = ormutl.GetEngine().Table(model).ID(id).Delete(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
		return nil
	} else {
		return nil
	}
}

func GetById(id interface{}, model interface{}) (err error) {
	if _, err = ormutl.GetEngine().Table(model).ID(id).Get(model); err != nil {
		logutl.Error(err.Error())
		return err
	}
	return nil
}
