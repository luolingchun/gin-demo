/*
@Author  : llc
@Time    : 2020/2/23 17:20
*/

package controllers

import (
	code "gin-demo/src/code"
	model "gin-demo/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//@Summary 获取所有任务
//@Tags v1/job
//@Description 获取所有任务
//@Success 200 {object} models.Job	"ok"
//@Router /v1/job [get]
func GetJobs(c *gin.Context) {
	var j model.Job

	jobs, err := j.GetJobs()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.GetError,
			"message": "数据获取错误::" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"message": "ok",
		"data":    jobs,
	})
}

//@Summary 获取单个任务
//@Tags v1/job
//@Description 获取单个任务
//@Param id path integer true "任务id"
//@Success 200 {object} models.Job	"ok"
//@Router /v1/job/{id} [get]
func GetJob(c *gin.Context) {
	var job model.Job

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id
	err := job.GetJob()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.GetError,
			"message": "数据获取错误::" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"message": "ok",
		"data":    job,
	})
}

//@Summary 新建任务
//@Tags v1/job
//@Description 新建任务
//@Accept multipart/form-data
//@Produce application/json
//@Param name formData string true "任务名称"
//@Success 200 {object} models.Job	"ok"
//@Router /v1/job [post]
func CreateJob(c *gin.Context) {
	var job model.Job

	err := c.Bind(&job)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.ArgsParseError,
			"message": "数据解析错误::" + err.Error(),
		})
		return
	}

	err = job.CreateJob()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.CreateError,
			"message": "数据创建错误::" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"message": "ok",
		"data":    job,
	})
}

//@Summary 更新任务
//@Tags v1/job
//@Description 更新任务
//@Accept multipart/form-data
//@Produce application/json
//@Param id path integer true "任务id"
//@Param result formData string true "任务结果"
//@Success 200 {string} string	"ok"
//@Router /v1/job/{id} [put]
func UpdateJob(c *gin.Context) {
	var job model.Job

	err := c.Bind(&job)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.ArgsParseError,
			"message": "数据解析错误::" + err.Error(),
		})
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id
	err = job.UpdateJob()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.UpdateError,
			"message": "数据更新错误::" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"message": "ok",
	})
}

//@Summary 删除任务
//@Tags v1/job
//@Description 删除任务
//@Param id path integer true "任务id"
//@Success 200 {string} string	"ok"
//@Router /v1/job/{id} [delete]
func DeleteJob(c *gin.Context) {
	var job model.Job

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id
	err := job.DeleteJob()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.DeleteError,
			"message": "数据删除错误::" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code.Success,
		"message": "ok",
	})
}
