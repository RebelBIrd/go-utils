package httputl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qinyuanmao/go-utils/logutl"
	"io/ioutil"
	"net/http"
	"strconv"
)

type BaseGroup struct {
	Path       string
	Routers    []BaseRouter
	Middleware gin.HandlerFunc
	Func       func(group *gin.RouterGroup)
}

type BaseRouter struct {
	Type       MethodType
	Path       string
	Middleware gin.HandlerFunc
	Handler    gin.HandlerFunc
}

func StartServer(groups []BaseGroup, routers []BaseRouter, port int, init func(engine *gin.Engine)) {
	engine := gin.Default()
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, Resp404Failed())
	})
	for _, group := range groups {
		group.Func = func(g *gin.RouterGroup) {
			for _, router := range group.Routers {
				router.addPath(g)
			}
		}
		group.Func(engine.Group(group.Path, group.Middleware))
	}
	for _, router := range routers {
		router.addEnginePath(engine)
	}
	if init != nil {
		init(engine)
	}
	if err := engine.Run(":" + strconv.Itoa(port)); err != nil {
		logutl.Error(err.Error())
	}
}

func (router *BaseRouter) addPath(group *gin.RouterGroup) {
	switch router.Type {
	case POST:
		if router.Middleware == nil {
			group.POST(router.Path, router.Handler)
		} else {
			group.POST(router.Path, router.Middleware, router.Handler)
		}
	case GET:
		if router.Middleware == nil {
			group.GET(router.Path, router.Handler)
		} else {
			group.GET(router.Path, router.Middleware, router.Handler)
		}
	case PUT:
		if router.Middleware == nil {
			group.PUT(router.Path, router.Handler)
		} else {
			group.PUT(router.Path, router.Middleware, router.Handler)
		}
	case DELETE:
		if router.Middleware == nil {
			group.DELETE(router.Path, router.Handler)
		} else {
			group.DELETE(router.Path, router.Middleware, router.Handler)
		}
	}
}

func (router *BaseRouter) addEnginePath(engine *gin.Engine) {
	switch router.Type {
	case POST:
		if router.Middleware == nil {
			engine.POST(router.Path, router.Handler)
		} else {
			engine.POST(router.Path, router.Middleware, router.Handler)
		}
	case GET:
		if router.Middleware == nil {
			engine.GET(router.Path, router.Handler)
		} else {
			engine.GET(router.Path, router.Middleware, router.Handler)
		}
	case PUT:
		if router.Middleware == nil {
			engine.PUT(router.Path, router.Handler)
		} else {
			engine.PUT(router.Path, router.Middleware, router.Handler)
		}
	case DELETE:
		if router.Middleware == nil {
			engine.DELETE(router.Path, router.Handler)
		} else {
			engine.DELETE(router.Path, router.Middleware, router.Handler)
		}
	}
}

func GetParam(ctx *gin.Context, key string) string {
	var value string
	value = ctx.PostForm(key)
	if value == "" {
		value = ctx.Query(key)
	}
	if value == "" {
		value = ctx.Param(key)
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

func GetInt64Param(ctx *gin.Context, key string) int64 {
	vStr := GetParam(ctx, key)
	if vStr == "" {
		return 0
	} else {
		v, err := strconv.ParseInt(vStr, 10, 64)
		if err != nil {
			logutl.Error(err.Error())
		}
		return v
	}
}
