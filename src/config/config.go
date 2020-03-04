/*
@Author  : llc
@Time    : 2020/3/1 13:25
*/

package config

import "github.com/gin-gonic/gin"

const (
	//服务模式配置
	ServerMode = gin.DebugMode //gin.ReleaseMode
	//数据库配置
	DBDialect  = "sqlite3" //mysql, postgres or mssql
	DBUser     = "user"
	DBPassword = "password"
	DBUri      = "localhost:5432"
	DBName     = "dbname"
	//日志配置
	LOGFile = "gin-demo.log"
)
