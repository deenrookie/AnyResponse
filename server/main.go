package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var apiResponseMap map[string]Response
var apiHandlersMap map[string]ApiHandler

func Init() {
	files := walkDir("apis")

	apiHandlersMap = make(map[string]ApiHandler)
	apiResponseMap = make(map[string]Response)

	for _, item := range files {
		content := getText(item)
		var api Response
		_ = json.Unmarshal([]byte(content), &api)
		apiResponseMap[api.ReqPath] = api
	}

	for key, value := range apiResponseMap {
		apiHandlersMap[key] = ApiHandler{
			value.Method,
			respHandler,
		}
	}
}

func main() {
	r := gin.Default()
	Init()
	registerApis(r) // 路由注册
	r.GET("/api/list", apiList)
	r.POST("/api/add", apiAdd)
	_ = r.Run("0.0.0.0:9999") // listen and serve on 0.0.0.0:8080
}
