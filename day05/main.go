package main

import "fmt"

// 分支结构
func main() {
	// 实现功能: 如果口罩的库存小于30个提示库存不足
	count := 30
	if count <= 30 {
		fmt.Println("对不起，口罩存量不足！")
	} else {
		fmt.Println("库存充足")
	}

	score := 90
	
	// 单分支效果 
	if score >= 90{
		fmt.Println("您的成绩为A级别")
	}
	if score >= 80 && score < 90{
		fmt.Println("您的成绩为B级别")
	}
	if score >= 70 && score < 80{
		fmt.Println("您的成绩为C级别")
	}
	if score >= 60 && score < 70 {
		fmt.Println("您的成绩为D级别")
	}
	if score < 60 {
		fmt.Println("您的成绩为E级别")
	}
	// 多分支 
	if score >= 90{
		fmt.Println("您的成绩为A级别")
	}else if score >= 80 && score < 90{
		fmt.Println("您的成绩为B级别")
	}else if score >= 70 && score < 80{
		fmt.Println("您的成绩为C级别")
	}else if score >= 60 && score < 70 {
		fmt.Println("您的成绩为D级别")
	}else if score < 60 {
		fmt.Println("您的成绩为E级别")
	}

	// switch 分支 
	switch score/10 {
	case 10 , 9:
		fmt.Println("您的等级为A级别")
	case 8 :
		fmt.Println("您的等级为A级别")
	case 7 :
		fmt.Println("您的等级为A级别")
	case 6 :
		fmt.Println("您的等级为A级别")
	case 5 :
		fmt.Println("您的等级为A级别")
	case 4 :
		fmt.Println("您的等级为A级别")
	case 3 :
		fmt.Println("您的等级为A级别")
	case 2 :
		fmt.Println("您的等级为A级别")
	case 1 :
		fmt.Println("您的等级为A级别")
	default:
		fmt.Println("您的等级为F级别")
	}
	
	a := 2
	switch{
	case a == 1 :
		fmt.Println("aaa")
	case a == 2 :
		fmt.Println("bbb")
	}

	switch b := 2; {
	case b == 1 : 
		fmt.Println("aaa")
	case b == 2 : 
		fmt.Println("bbb")
	}

	switch score/10 {
	case 10 , 9:
		fmt.Println("您的等级为A级别")
		fallthrough // 穿透一层！ 
	case 8 :
		fmt.Println("您的等级为B级别")
		fallthrough // 再穿透一层!
	case 7 :
		fmt.Println("您的等级为A级别")
	case 6 :
		fmt.Println("您的等级为A级别")
	case 5 :
		fmt.Println("您的等级为A级别")
	case 4 :
		fmt.Println("您的等级为A级别")
	case 3 :
		fmt.Println("您的等级为A级别")
	case 2 :
		fmt.Println("您的等级为A级别")
	case 1 :
		fmt.Println("您的等级为A级别")
	default:
		fmt.Println("您的等级为F级别")
	}
}