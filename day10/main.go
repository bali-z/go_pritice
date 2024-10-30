package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串 基础操作 
func main(){
	// 1. 统计字符串的长度，按照字节数统计 len(str)
	str := "zhangdezhong 好帅！"
	fmt.Printf("字符串str的字节数是 %d \n",len(str))
	// 2. 字符串遍历 
	//（1）for range 遍历
	for i,c := range str {
		fmt.Printf("str 第 %d 位字符为 %c \n",i,c)
	} 
	// (2) 切片方式遍历 rune(str) 
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("str 第 %d 位为 %c \n",i,r[i])
	}

	// 3. 字符串转整数&整数转字符串
	str1 := "456"
	num1,_ := strconv.Atoi(str1)
	fmt.Println("字符串转整数结果 ",num1)
	fmt.Println("整数转字符串结果 ",strconv.Itoa(23))

	// 4. 判断字串是否存在于目标串中
	exist := strings.Contains("golanggolang","lang")
	fmt.Println("字符串 'golanggolang' 是否包含字串 'lang'",exist)

	// 5. 不区分大小写比较两个字符串是否相同
	fmt.Println("字符串 'Go' 和 字符串 'gO' 不区分大小写是否相同",strings.EqualFold("Go","gO"))

	// 6. 获取字符串在目标串第一次出现的下标
	fmt.Println("字符串 'lang' 在 字符串 'golanggolang' 中第一次出现的位置为 ",strings.Index("golanggolang","lang"))

	// 7. 字符串的替换
	str2 := "golanggolanggolanggo"
	// 全部替换 
	fmt.Println("'golanggolanggolanggo' 替换 'go' to 'java'",strings.Replace(str2,"go","java",-1))
	// 替换前两个 
	fmt.Println("'golanggolanggolanggo' 替换 'go' to 'python'",strings.Replace(str2,"go","python",2))

	// 8. 字符串分割
	strArr := strings.Split("java--python--golang","--")
	for _,s := range strArr{
		fmt.Printf("字符串分割结果 %s \n",s)
	}
	// 9. 大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))

	// 10. trim 测试发现，这个会自动识别有的有的话就直接trim了 
	fmt.Println(strings.TrimSpace("     go lang   "))

	// 11. trim 指定字符串
	fmt.Println(strings.Trim("~~~~~go~~lang~~~","~~~"))

	// 12. trim 左侧指定字符串 & trim 右侧指定字符串
	fmt.Println(strings.TrimLeft("~~~~~golang~~~","~~"))
	fmt.Println(strings.TrimRight("~~~~~golang~~~","~~"))

	// 13. 判断字符串是否以某个串开头
	fmt.Println(strings.HasPrefix("http://oracle.google.com","http://"))

	// 14. 判断字符串是否以某个串结尾
	fmt.Println(strings.HasSuffix("http://oracle.google.com","com"))

	return 
}