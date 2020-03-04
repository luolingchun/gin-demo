/*
@Author  : llc
@Time    : 2020/3/4 14:59
*/

package models

import (
	"fmt"
	orm "gin-demo/src/db"
)

//初始化数据库
func init() {
	fmt.Println("初始化数据库...")
	if ok := orm.DB.HasTable(&Job{}); ok {
		fmt.Println("数据库表已存在：job")
		return
	} else {
		orm.DB.CreateTable(&Job{})
	}

}
