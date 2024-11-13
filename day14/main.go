package main

import (
	"fmt"
)


// slice 切片学习
func main(){
	var arr [6]int = [6]int{1,2,3,4,5,6}
	var intsclice []int = arr[1:4]
	fmt.Println("切片intsclice的长度为: ",len(intsclice))
	fmt.Println("切片intsclice的容量为: ",cap(intsclice))
	fmt.Println("切片intsclice的内容为: ",intsclice)

	// 切片是一个结构体 包含三个主要部分 1. 切片数组地址. 2. 切片长度。 3. 切片数组长度（切片容量）
	var slice1 = make([]int,3,6)
	slice1[0] = 1
	slice1[2] = 3
	fmt.Println("切片sclice1的长度为: ",len(slice1))
	fmt.Println("切片sclice1的容量为: ",cap(slice1))
	fmt.Println("切片sclice1的内容为: ",slice1)
	// sclice ps： 切片类似于 shell , 切片可以继续切
}