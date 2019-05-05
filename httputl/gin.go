package httputl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetParam(ctx *gin.Context, key string) interface{} {
	var value interface{}
	value = ctx.PostForm(key)
	if value == "" {
		value = ctx.Query(key)
	}
	if value == "" {
		var values map[string]interface{}
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		_ = json.Unmarshal(body, &values)
		value = values[key]
	}
	return value
}
