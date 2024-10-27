package main

import "fmt"

// 函数 
// 基本语法 
/**
	func 函数名称 (形参列表) (返回值类型列表){
		执行语句 ---
		return + 返回值列表
	}

	函数名：
		首字母不能是数字
		首字母大写该函数可以被本包文件和其他包文件使用 类似于public
		首字母小写该函数只能被本包使用，其他包文件不能使用 类似与 private
	
		函数不支持重载 
		支持可变参数 
		函数也是一种数据类型，可以复制给一个变量

		支持自定义数据类型 
		type 关键字 




**/


func cal (num1 int,num2 int) (int){
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func calc(num1 int,num2 int)(sum int,sub int){
	sum = num1 + num2
	sub = num1 - num2
	return 
}
type CALC func (num1 int,num2 int)(sum int,sub int)
// 函数作为参数传递
func test1(num1 int,num2 int,CALC func(int,int)(int,int))(sum int,sub int){
	return CALC(num1,num2)
}
// 函数作为参数传递
func test2(num1 int,num2 int,testFunc func(int,int)(int,int))(sum int,sub int){
	return testFunc(num1,num2)
}

// 函数可变参数 
func test(args... int)(){
	for _,arg := range args{
		fmt.Println("arg ",arg)
	}
}

func main(){
	sum := cal(1,2)
	a := cal
	fmt.Printf("a 的类型是 %T,cal 函数的类型是 %T \n",a,cal)
	fmt.Println(a(1,6))
	fmt.Println("sum = ",sum)
	test(1,2,3,4)

	type myInt int
	var num1 myInt = 30
	var num2 int = 19
	// num1 = num2 会失败 类型不一致 因为类型名不一致
	num1 = myInt(num2)
	fmt.Printf("num1 : %d" ,num1)
	sum, sub := test1(10, 5, calc)
	fmt.Println("Sum:", sum, "Sub:", sub)

	sum, sub = test2(10, 9, calc)
	fmt.Println("Sum:", sum, "Sub:", sub)
}



