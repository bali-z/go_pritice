package main 

import (
	"fmt"
	"time"
)

// 时间和内置函数 

func main(){
	now := time.Now()
	fmt.Printf("%v 对应的类型是 %T \n",now,now)

	//日期格式化 
	datastr := now.Format("2006-01-02 15:04:05")
	fmt.Println(datastr)

	// 内置函数 builtin 包内的都是
	// new 函数 返回一个指定类型对象的指针
	num := new(int) 
	fmt.Printf("num 是一个指针,num中存入的地址是 %v,num 自己的地址是 %v,num 指针指向的地址的值是 %v \n",num,&num,*num)
	// make 函数 分配内存，引用类型 指针，slice 切片,map,chan 管道,interface 等
}