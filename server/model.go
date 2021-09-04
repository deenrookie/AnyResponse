package main

import "github.com/gin-gonic/gin"

type Response struct {
	Method     string            `form:"method" json:"method"`
	ReqPath    string            `form:"path" json:"path"`
	StatusCode int               `form:"status_code" json:"status_code"`
	Headers    map[string]string `form:"headers" json:"headers"`
	Body       string            `form:"body" json:"body"`
	Token      string            `form:"token" json:"token"`
}

type ApiHandler struct {
	Method  string
	Handler gin.HandlerFunc
}
