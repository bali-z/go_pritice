package main

import (
	"fmt"
)

// 数组学习 
func main(){
	var arr = [3]int{3,6,9}
	var arr1 [2]int 
	var arr2 = [...]int{1,24,5,6}
	arr3 := [...]int8{3,2,1}

	// 数组需要注意的是数组的长度是数组类型的一部分:  arr 的类型是 [3]int,arr1的类型是 [2]int,arr2 的类型是 [4]int
	fmt.Printf("arr 的类型是 %T,arr1的类型是 %T,arr2 的类型是 %T \n",arr,arr1,arr2)

	test01(&arr3)

	var arr4 = [2][3]int{{1,2,3},{4,5,6}}
	//arr4 的类型是 [2][3]int
	fmt.Printf("arr4 的类型是 %T \n",arr4)
}
// 
func test01(arr *[3]int8){
	// 输出结果 &[0 0] 
	fmt.Println(arr)
	// 输出结果
	fmt.Println((*arr)[0])
// arr 的值为 0xc00008c098
// arr 的值为 0xc00008c098
// arr 的值为 0xc00008c099
// arr 的值为 0xc00008c09a
	fmt.Printf("arr 的值为 %p\n",&(*arr))
	fmt.Printf("arr 的值为 %p\n",&((*arr)[0]))
	fmt.Printf("arr 的值为 %p\n",&((*arr)[1]))
	fmt.Printf("arr 的值为 %p\n",&((*arr)[2]))

}


