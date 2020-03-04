/*
@Author  : llc
@Time    : 2020/2/23 17:20
*/

package controllers

import (
	model "gin-demo/src/models"
	StatusCode "gin-demo/src/status-code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//@Summary 获取所有任务
//@Tags v1/job
//@Description 获取所有任务
//@Success 200 {object} models.Response	"ok"
//@Router /v1/job [get]
func GetJobs(c *gin.Context) {
	var res model.Response
	var j model.Job

	jobs, err := j.GetJobs()
	if err != nil {
		res.Code = StatusCode.GetError
		res.Message = "数据获取错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.Code = StatusCode.Success
	res.Message = "ok"
	res.Data = jobs
	c.JSON(http.StatusOK, res)
}

//@Summary 获取单个任务
//@Tags v1/job
//@Description 获取单个任务
//@Param id path integer true "任务id"
//@Success 200 {object} models.Response	"ok"
//@Router /v1/job/{id} [get]
func GetJob(c *gin.Context) {
	var res model.Response
	var job model.Job

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id
	err := job.GetJob()
	if err != nil {
		res.Code = StatusCode.GetError
		res.Message = "数据获取错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.Code = StatusCode.Success
	res.Message = "ok"
	res.Data = job
	c.JSON(http.StatusOK, res)
}

//@Summary 新建任务
//@Tags v1/job
//@Description 新建任务
//@Accept multipart/form-data
//@Produce application/json
//@Param name formData string true "任务名称"
//@Success 200 {object} models.Response	"ok"
//@Router /v1/job [post]
func CreateJob(c *gin.Context) {
	var res model.Response
	var job model.Job

	err := c.Bind(&job)

	if err != nil {
		res.Code = StatusCode.ArgsParseError
		res.Message = "数据解析错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	err = job.CreateJob()

	if err != nil {
		res.Code = StatusCode.CreateError
		res.Message = "数据创建错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.Code = StatusCode.Success
	res.Message = "ok"
	res.Data = job
	c.JSON(http.StatusOK, res)
}

//@Summary 更新任务
//@Tags v1/job
//@Description 更新任务
//@Accept multipart/form-data
//@Produce application/json
//@Param id path integer true "任务id"
//@Param result formData string true "任务结果"
//@Success 200 {object} models.Response	"ok"
//@Router /v1/job/{id} [put]
func UpdateJob(c *gin.Context) {
	var res model.Response
	var job model.Job

	err := c.Bind(&job)

	if err != nil {
		res.Code = StatusCode.ArgsParseError
		res.Message = "数据解析错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id
	err = job.UpdateJob()

	if err != nil {
		res.Code = StatusCode.UpdateError
		res.Message = "数据更新错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.Code = StatusCode.Success
	res.Message = "ok"
	c.JSON(http.StatusOK, res)
}

//@Summary 删除任务
//@Tags v1/job
//@Description 删除任务
//@Param id path integer true "任务id"
//@Success 200 {object} models.Response	"ok"
//@Router /v1/job/{id} [delete]
func DeleteJob(c *gin.Context) {
	var res model.Response
	var job model.Job

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	job.ID = id

	err := job.DeleteJob()
	if err != nil {
		res.Code = StatusCode.DeleteError
		res.Message = "数据删除错误::" + err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res.Code = StatusCode.Success
	res.Message = "ok"
	c.JSON(http.StatusOK, res)
}
