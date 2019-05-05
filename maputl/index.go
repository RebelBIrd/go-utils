package maputl

import (
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/qinyuanmao/go-utils/logutl"
	"reflect"
)

/**
 * struct转map
 * @param obj待转的struct
 * @return map
 */
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/**
 * map转struct
 * @param mp map
 * @param obj 待转的struct，必须为指针
 * @return error
 */
func Map2Struct(mp map[string]interface{}, obj interface{}) (err error) {
	err = mapstructure.Decode(mp, obj)
	if err != nil {
		logutl.Error(err.Error())
	}
	return
}

/**
 * 取出map[string]interface{}中的value，注意：只能取基础类型
 * @param m map[string]interface{} 待取值的map
 * @param key map的key值
 * @param value 函数外绑定的值，必须是指针
 * @return error 如果错误则返回
 */
func GetMapValue(m map[string]interface{}, key string, value interface{}) error {
	if m[key] == nil {
		fmt.Println("Param not found.")
		return errors.New("Param not found.")
	}
	if "*"+reflect.TypeOf(m[key]).String() == reflect.TypeOf(value).String() {
		switch m[key].(type) {
		case string:
			v := value.(*string)
			*v = m[key].(string)
		case float64:
			v := value.(*float64)
			*v = m[key].(float64)
		case bool:
			v := value.(*bool)
			*v = m[key].(bool)
		case uint8:
			v := value.(*uint8)
			*v = m[key].(uint8)
		case uint16:
			v := value.(*uint16)
			*v = m[key].(uint16)
		case uint32:
			v := value.(*uint32)
			*v = m[key].(uint32)
		case uint64:
			v := value.(*uint64)
			*v = m[key].(uint64)
		case int8:
			v := value.(*int8)
			*v = m[key].(int8)
		case int16:
			v := value.(*int16)
			*v = m[key].(int16)
		case int32:
			v := value.(*int32)
			*v = m[key].(int32)
		case int64:
			v := value.(*int64)
			*v = m[key].(int64)
		case float32:
			v := value.(*float32)
			*v = m[key].(float32)
		case complex64:
			v := value.(*complex64)
			*v = m[key].(complex64)
		case complex128:
			v := value.(*complex128)
			*v = m[key].(complex128)
		case int:
			v := value.(*int)
			*v = m[key].(int)
		case uint:
			v := value.(*uint)
			*v = m[key].(uint)
		case uintptr:
			v := value.(*uintptr)
			*v = m[key].(uintptr)
		default:
			v := value.(*map[string]interface{})
			*v = m[key].(map[string]interface{})
		}
		return nil
	} else {
		fmt.Println("Param type error.")
		return errors.New("Param type error.")
	}
}
