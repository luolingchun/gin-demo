/*
@Author  : llc
@Time    : 2020/3/1 13:28
*/

package db

import (
	"fmt"
	"gin-demo/src/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

//连接数据库
func init() {
	var err error
	//dbSource mysql "root:123@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	//dbSource postgres "host=host port=port user=gorm password=password dbname=dbname"
	dbSource := config.DBName + ".db"
	DB, err = gorm.Open(config.DBDialect, dbSource)

	if err != nil {
		fmt.Printf("database connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)

	}
}
