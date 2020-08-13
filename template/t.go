package template

const mainStr = `package main

func main() {
	
}
`

const controllerStr = `package controllers
`

const routerStr = `package routers

import (
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
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