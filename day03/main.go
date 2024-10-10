package main

import ("fmt"
		"strings"
		"strconv")

// 字符串  和 C/C++不一样，Go 中的字符串是根据长度限定，而非特殊字符 \0

// 作为一种基本数据结构，每种语言都有一些对于字符串的预定义处理函数。Go 中使用 strings 包来完成对字符串的主要操作。
func main() {
	str := "Beginning of the string " +
    "second part of the string"
	fmt.Println(str)
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”

	// 1.前缀后缀 
	var st string = "This is an example of a string"
    fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", st, "Th")
    fmt.Printf("%t\n", strings.HasPrefix(st, "Th"))

	fmt.Printf("T/F? Does the string \"%s\" have suffix %s? ", st, "ing")
    fmt.Printf("%t\n", strings.HasSuffix(st, "ing"))
	
	// 2.字符串包含关系 
	fmt.Printf("T/F? Is string \"%s\" in string \"%s\"? ", "an", st)
    fmt.Printf("%t\n", strings.Contains(st, "an"))

	// 3.字符串查找
	fmt.Printf("The index of the first occurrence of \"%s\" in \"%s\" is %d\n",
    "example", st, strings.Index(st, "example"))
	fmt.Printf("The index of the last occurrence of \"%s\" in \"%s\" is %d\n",
    "example", st, strings.LastIndex(st, "example"))
	// 如果需要查询非 ASCII 编码的字符在父字符串中的位置
	fmt.Printf("%d",strings.IndexRune(st, 'e'))
	
	// 4.字符串替换
	fmt.Println(strings.Replace("foo", "o", "0", -1))
	fmt.Println(strings.Replace("foo", "o", "0", 1))

	// 5.字符串分割
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// 6.字符串连接
	fmt.Println(strings.Join([]string{"a", "b"},"-"))
	
	// 7.字符串大小写转换
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.Title("hello world"))

	// 8.统计字符出现次数
	fmt.Println(strings.Count("cheese", "e"))
	
	// 9.Repeat 重复字符串
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Repeat("a", 0))

	str = "The quick brown fox jumps over the lazy dog"
    sl := strings.Fields(str)
    fmt.Printf("Splitted in slice: %v\n", sl)
    for _, val := range sl {
        fmt.Printf("%s - ", val)
    }
    fmt.Println()
    str2 := "GO1|The ABC of Go|25"
    sl2 := strings.Split(str2, "|")
    fmt.Printf("Splitted in slice: %v\n", sl2)
    for _, val := range sl2 {
        fmt.Printf("%s - ", val)
    }
    fmt.Println()
    str3 := strings.Join(sl2,";")
    fmt.Printf("sl2 joined by ;: %s\n", str3)

	// 从字符串中读取内容
	fmt.Println(strings.TrimSpace(" \t\n Hello world \n\t"))

	// 函数 strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，
	// 然后返回指向该 Reader 的指针，从其它类型读取内容的函数还有：
	// Read() 从 []byte 中读取内容。
	// ReadByte() 和 ReadRune() 从字符串中读取下一个 byte 或者 rune。


	// 字符串与其他类型转换 
	var orig string = "666"
    var an int
    var newS string
    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)      
    an, _ = strconv.Atoi(orig)
    fmt.Printf("The integer is: %d\n", an) 
    an = an + 5
    newS = strconv.Itoa(an)
    fmt.Printf("The new string is: %s\n", newS)
	
}