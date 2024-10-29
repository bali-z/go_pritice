package main	


import (
	"fmt"
	db "go_project/go_pritice/day09/db"
)

var val int = 20

func init(){
	fmt.Println("main init 执行了！")
}

// init 函数 类似与 React js 等 自带的初始化

// 全局匿名函数 
var Clear = func(){
	fmt.Println("clear clear 执行了！")
}


// 闭包
func getSumFunc() func( val int) (int){
	var sum int = 1
	return func(val int) int {
		sum += val
		return sum
	}
}  


func main(){
	fmt.Println("main main 执行了！")
	db.GetConnection()
	
	// 局部 匿名函数 
	var f = func (val1 int, val2 int) int {
		return val1 + val2
	}
	sum := f(1,2)
	fmt.Println(sum)
	f1 := getSumFunc()
	fmt.Println(f1(20)) // 结果为 21
	fmt.Println(f1(30)) // 结果为 51

	connect(1)
	defer disConnect(1)

	fmt.Println("使用资源 1")
	return
}


// defer 关键字 ，defer 后面的语句会延迟到函数返回前才执行，可以用于资源的自动释放


func connect(fd int ){
	fmt.Println("获取一个连接",fd)
}

func disConnect(fd int){
	fmt.Println("释放一个连接",fd)
}

