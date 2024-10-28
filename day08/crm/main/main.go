package main	

import (
	"fmt"
	"go_project/go_pritice/day08/crm/dbutils"
)

func main(){
	fmt.Println("开始调用dbutils包下的获取数据库链接方法")
	// dbutils.getConnection()
	dbutils.GetConnection()
}