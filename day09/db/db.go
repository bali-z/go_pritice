package db

import (
	"fmt"
)

var val int = 10


func init(){
	fmt.Println("db init 执行了！")
}

func GetConnection(){
	fmt.Println("db getConnection 执行了！")
}