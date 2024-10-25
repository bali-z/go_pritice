package main

import "fmt"

func main(){
	// 实现一个功能求和：
	var i1 int = 1
	var i2 int = 2
	var i3 int = 3
	var i4 int = 4

	var sum int = 0

	sum += i1
	sum += i2
	sum += i3
	sum += i4
	fmt.Println(sum)

	sum = 0
	val := 1
	for i := 0; i < 4; i++{
		sum += val
		val ++  
	}
	fmt.Println(sum)

	// for to while
	i := 0
	for i < 4 {
		sum += val
		val ++
		i++
	}
	fmt.Println(sum)

	// for range 
	str := "hello golang 你好"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n",str[i])
	}

	for i,value := range str{
		fmt.Printf("索引为 %d,值为 %c\n",i,value)
	}


	// break 
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i 的值为 %d,j 的值为 %d \n",i,j)
			if i == 2 && j == 2{
				break // 终止最近的循环 
			}
		}
	}

	fmt.Println("---------------ok")
	// break 配合 label 
	label1: 
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i 的值为 %d,j 的值为 %d \n",i,j)
			if i == 2 && j == 2{
				break label1// 终止最近的循环 
			}
		}
	}
	fmt.Println("---------------ok")

	// continue 
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			if i == 2 && j == 2{
				continue // 终止最近的循环 
			}
			fmt.Printf("i 的值为 %d,j 的值为 %d \n",i,j)
		}
	}

	fmt.Println("---------------ok")
	// continue 配合 lable 
	lable2: 
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			if i == 2 && j == 2{
				continue lable2 // 终止最近的循环 
			}
			fmt.Printf("i 的值为 %d,j 的值为 %d \n",i,j)
		}
	}


	// go to 

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	goto label3
	fmt.Println("5")
	fmt.Println("6")
	label3: 
	fmt.Println("7")


	// return 
	for i := 1; i <= 100; i++{
		fmt.Println(i)
		if i == 15 {
			return 
		}
	}
	fmt.Println("over -------------")
}