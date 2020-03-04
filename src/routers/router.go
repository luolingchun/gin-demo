/*
@Author  : llc
@Time    : 2020/2/23 17:21
*/

package routers

import (
	. "gin-demo/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/job", GetJobs)
		v1.POST("/job", CreateJob)
		v1.GET("/job/:id", GetJob)
		v1.PUT("/job/:id", UpdateJob)
		v1.DELETE("/job/:id", DeleteJob)
	}

	//v2 := r.Group("/v2")
	//{
	//	v2.GET("/job", GetJobs2)
	//}

	return r
}
