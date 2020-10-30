package template

const mainStr = `package main

import (
	"%s/routers"
)

func main() {
	r := routers.SetRouter()
	r.Run(":18088")
}
`

const controllerStr = `package controllers
`

const routerStr = `package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s\n", "hello gin")
		return
	})
	return engine
}
`

const modelStr = `package models
`

const goModuleStr = `module %s

go %s
`