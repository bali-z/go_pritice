package main

import (
	"fmt"
	"errors"
)

func main(){
	// fmt.Println("调用函数 testErrorCach ")
	// testErrorCach(10,0)
	// fmt.Println("调用结束.....")

	fmt.Println("自定义异常测试")
	_,err :=testSelfError(10,0)
	if err != nil {
		fmt.Println("出现异常 ",err)
		// 终止当前执行流
		// panic(err) 异常终止 
	}
	fmt.Println("自定义异常测试完毕")
}
// defer + recover 实现异常捕获处理
func testErrorCach(num1 int,num2 int){
	defer func(){
		err := recover()
		if nil != err {
			// 捕获到异常了 
			fmt.Println("捕获到异常了 ",err)
			fmt.Println("处理异常！！！！！！！！")
		}
	}()
	fmt.Println(num1/num2)
}

// 自定义异常
func testSelfError(num1 int,num2 int)(num int,err error) {
	if num2 == 0 {
		return 0,errors.New("被除数为 0 了")
	}
	return num1 / num2,nil 
}