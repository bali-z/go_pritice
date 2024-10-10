package main 

import "fmt"

func main() {
	// 变量定义的四种方式 
	
	// 局部变量 
	// 1. 定义后赋值
	var a int
	a = 10
	fmt.Println(a)
	// 2. 定义并赋值
	var b int = 20
	fmt.Println(b)
	
	// 3. 仅声明
	var c int
	fmt.Println(c)
	
	// 4. 自动类型判断
	d := 30
	fmt.Println(d)
	
	// 多个变量定义
	var q,w,e int = 1,2,3
	
	fmt.Println(q,w,e)
	
	var f,g,h = 1,2,3
	
	fmt.Println(f,g,h)
	
	var (
		i int = 1
		j int = 2
		k int = 3
	)
	fmt.Println(i,j,k)
	
	const pi = 3.14
	fmt.Println(pi)
	
	const (
		a1 = 1
		b1 = 2
		c1 = 3
	)
	fmt.Println(a1,b1,c1)
	
	var (
		a2 int = 1
		b2 int = 2
		c2 int = 3
	)
	fmt.Println(a2,b2,c2)
}