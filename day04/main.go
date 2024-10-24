package main

import "fmt"
// 从控制台获取输入

func main(){
	// 用户输入学生姓名，年龄，成绩，是否妹子
	var name string
	var age int
	var score float32
	var isMeizi bool
	// Scanln() 行输入 
	// fmt.Println("请输入学生姓名:")
	// fmt.Scanln(&name)
	// fmt.Println("请输入学生年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入学生成绩:")
	// fmt.Scanln(&score)
	// fmt.Println("请输入学生是否妹子:")
	// fmt.Scanln(&isMeizi)
	// fmt.Printf("学生姓名: %v,年龄: %v,成绩: %v,是否妹子: %v",name,age,score,isMeizi)
	// Scanf() 格式化输入
	fmt.Println("请输入学生姓名，年龄，成绩，是否妹子以逗号分隔:")
	fmt.Scanf("%s,%d,%f,%t",&name,&age,&score,&isMeizi)
	fmt.Printf("学生姓名: %v,年龄: %v,成绩: %v,是否妹子: %v",name,age,score,isMeizi)

}