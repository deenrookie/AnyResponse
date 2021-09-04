package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
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
	for path, item := range apiHandlersMap {
		if item.Method == "GET" {
			r.GET(path, item.Handler)
		}
		if item.Method == "POST" {
			r.POST(path, item.Handler)
		}

	}
}

// 接口列表
func apiList(r *gin.Context) {
	getToken := r.Query("token")
	if getToken != token {
		r.JSON(http.StatusForbidden, gin.H{"error": "token error"})
		return
	}

	ret := make([]Response, 0)
	for _, item := range apiResponseMap {
		ret = append(ret, item)
	}
	r.JSON(200, ret)
}

// 添加接口，保存生成txt，配合gowatch自动重启
func apiAdd(c *gin.Context) {

	var form Response
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Token != token {
		c.JSON(http.StatusForbidden, gin.H{"error": "token error"})
		return
	}
	fmt.Println("token", form.Token)

	if form.ReqPath == "" {
		form.ReqPath = fmt.Sprintf("/%s", md5v1(timeNow()))
	}

	if !strings.HasPrefix(form.ReqPath, "/") {
		form.ReqPath = fmt.Sprintf("/%s", form.ReqPath)
	}

	jsonText, _ := json.Marshal(form)
	err := ioutil.WriteFile(fmt.Sprintf("apis/%v.txt", len(apiHandlersMap)+1), jsonText, 0644)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
