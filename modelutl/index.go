package modelutl

import (
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/ormutl"
	"github.com/qinyuanmao/go-utils/sliceutl"
	"log"
	"reflect"
)

func Exist(id interface{}, model interface{}) bool {
	exist, err := ormutl.GetEngine().Table(model).ID(id).NoCache().NoAutoCondition().Exist(model)
	if err != nil {
		logutl.Error(err.Error())
	}
	return exist
}

func Save(id interface{}, model interface{}) (err error) {
	if _, err = ormutl.GetEngine().Insert(model); err != nil {
		logutl.Error(err)
		return err
	}
	if _, err = ormutl.GetEngine().ID(id).NoAutoCondition().Get(model); err != nil {
		logutl.Error(err)
		return err
	}
	return
}

func SaveArray(m interface{}) (err error) {
	models := sliceutl.InterfaceSlice(m)
	for i := 0; i < len(models); i = i + 99 {
		if i+99 >= len(models) {
			_, err = ormutl.GetEngine().Insert(models[i:])
			if err != nil {
				logutl.Error(err)
			}
		} else {
			_, err = ormutl.GetEngine().Insert(models[i : i+99])
			if err != nil {
				logutl.Error(err)
			}
		}
	}
	return
}

func Update(id interface{}, model interface{}) (err error) {
	if Exist(id, model) {
		if _, err = ormutl.GetEngine().NoAutoCondition().UseBool(GetBoolField(model)...).ID(id).Update(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
	} else {
		if _, err = ormutl.GetEngine().Insert(model); err != nil {
			logutl.Error(err.Error())
			return err
		}
	}
	if _, err = ormutl.GetEngine().NoAutoCondition().ID(id).Get(model); err != nil {
		logutl.Error(err.Error())
		return err
	}
	return
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
	if _, err = ormutl.GetEngine().Table(model).ID(id).NoAutoCondition().Get(model); err != nil {
		logutl.Error(err.Error())
		return err
	}
	return nil
}

func GetBoolField(model interface{}) (fields []string) {
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
