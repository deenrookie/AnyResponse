package main

import "github.com/gin-gonic/gin"

type Response struct {
	Method     string            `json:"method"`
	ReqPath    string            `json:"path"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type ApiHandler struct {
	Method  string
	Handler gin.HandlerFunc
}
