/*
@Author  : llc
@Time    : 2020/3/1 14:09
*/

package models

//响应结构体
type Response struct {
	Code    int         `json:"code" example:"0"`          //业务状态码
	Message string      `json:"message" example:"ok"`      //异常信息
	Data    interface{} `json:"data" swaggertype:"object"` //数据体
}
