package main

import (
	"github.com/gin-gonic/gin"
)

var apiResponseMap map[string]Response

func main() {
	r := gin.Default()
	registerApis(r) // 路由注册
	_ = r.Run("0.0.0.0:9999") // listen and serve on 0.0.0.0:8080
}
