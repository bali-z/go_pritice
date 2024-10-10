package main

// go 数据类型

import ("fmt"
		"math")

func main(){
	// bool
	var ok bool
	ok = true
	fmt.Println(ok)

	// int 有符号整数
	var a int = 12
	var b int8 = 12
	var c int16 = 12
	var d int32 = 12
	var e int64 = 12
	
	fmt.Println(a,b,c,d,e)
	// uint 无符号整数 
	var f uint = 22 
	var g uint8 = 22
	var h uint16 = 22
	var i uint32 = 22
	var j uint64 = 22

	fmt.Println(f,g,h,i,j)

	// float
	var k float32 = 12.3
	var l float64 = 12.3
	fmt.Println(k,l)

	
	var n int16 = 34
    var m int32
    // compiler error: cannot use n (type int16) as type int32 in assignment
    //m = n
    m = int32(n)
    fmt.Printf("32 bit int is: %d\n", m)
    fmt.Printf("16 bit int is: %d\n", n)

	// 在格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字），
	// %g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），
	// %0nd 用于规定输出长度为 n 的整数，其中开头的数字 0 是必须的
	fmt.Printf("int16 %d\n", n)
	fmt.Printf("int32 %08d\n", m)
	fmt.Printf("float64 %g\n", 3.1415926)
	fmt.Printf("float64 %.2f\n", 3.1415926)
	fmt.Printf("float64 %.2e\n", 3.1415926)
	fmt.Printf("float64 %.2E\n", 3.1415926)
	fmt.Printf("float64 %.2f\n", 3.1415926)
	fmt.Printf("float64 %.2E\n", 3.1415926)
	fmt.Printf("float64 %.2e\n", 3.1415926)

	// 复数
	var complex1 complex64 = 5 + 10i;
	fmt.Printf("The value is: %v\n",complex1)
	
	var complex2 complex128 = 5 + 10i;
	fmt.Printf("The value is: %v\n",complex2)
	
	x := complex(3.1,2.2)
	fmt.Printf("x is %v\n",x)
	
	rea := real(x)
	img := imag(x)
	fmt.Printf("real(x) is %v\n",rea)
	fmt.Printf("imag(x) is %v\n",img)

	// 类型别名
	type myInt int
	var aa myInt = 12
	fmt.Println(aa)
	
	// 位运算 位清除 &^ 
	var xx uint8 = 15 // 00001111
	var yy uint8 = 4
	fmt.Printf("%08b\n", xx &^ yy);  // 00001011

	// 字符类型 byte 类型是 uint8 的别名
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3) // UTF-8 code point
	
}

// int 转 uint8
func Uint8FromInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 { // conversion is safe
        return uint8(n), nil
    }
    return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

// float64 转 int
func IntFromFloat64(x float64) int {
    if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
        whole, fraction := math.Modf(x)
        if fraction >= 0.5 {
            whole++
        }
        return int(whole)
    }
    panic(fmt.Sprintf("%g is out of the int32 range", x))
}