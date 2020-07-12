/*
@Author  : llc
@Time    : 2020/2/23 17:11
*/

package main

import (
	"gin-demo/src/config"
	_ "gin-demo/src/docs"
	. "gin-demo/src/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"net/http"
	"os"
)

// @title 任务管理 API
// @version 1.0
// @description 任务管理系统

// @tag.name v1/job
// @tag.description 任务管理
func main() {
	//服务模式
	gin.SetMode(config.ServerMode)
	//日志
	file, _ := os.Create(config.LOGFile)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) //保存到文件并在控制台中输出
	gin.DefaultErrorWriter = io.MultiWriter(file, os.Stdout)
	//初始化路由
	r := InitRouter()
	//跨域
	r.Use(cors.Default())
	//swagger文档
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	//跟路由重定向到swagger文档
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	//启动服务
	_ = r.Run(":8080")
}
