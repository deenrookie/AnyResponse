package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func respHandler(c *gin.Context) {
	reqPath := c.Request.URL.Path
	respStruct := apiResponseMap[reqPath]
	if respStruct.StatusCode != 0 {
		fmt.Println(reqPath)
		for key, value := range respStruct.Headers {
			c.Header(key, value)
		}
		c.String(respStruct.StatusCode, respStruct.Body)
	}
}

// 根据apis目录下的json生成接口
func registerApis(r *gin.Engine) {
	files := walkDir("apis")
	fmt.Println(files)

	apiHandlersMap := make(map[string]ApiHandler)
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
	for path, item := range apiHandlersMap {
		if item.Method == "GET" {
			r.GET(path, item.Handler)
		}
		if item.Method == "POST" {
			r.POST(path, item.Handler)
		}

	}
}
