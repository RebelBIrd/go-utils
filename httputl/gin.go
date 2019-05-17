package httputl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/go-utils/logutl"
	"io/ioutil"
	"strconv"
)

func GetParam(ctx *gin.Context, key string) string{
	var value string
	value = ctx.PostForm(key)
	if value == "" {
		value = ctx.Query(key)
	}
	if value == "" {
		var values map[string]string
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		_ = json.Unmarshal(body, &values)
		value = values[key]
	}
	return value
}
func GetIntParam(ctx *gin.Context, key string) int {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		return 0
	} else {
		v, err := strconv.Atoi(vStr)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}
