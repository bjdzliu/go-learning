package main

import (
	"fmt"
	"strconv"
)

type (
	byte     int8
	rune     int32
	ByteSize int64
)

var x1 byte = 100 //全局不能忽略var
var x2 = 101      //类型推断

var (
	g1 = 2 //并列方式，类型推断
	g2 = 5
	g3 = 7
)

// 全局变量不能使用 x3 := 102

func test() {
	x3 := 102
	var x4 int32 = 3
	fmt.Println(x1, x2, x3, x4, g1, g2)
}

func main() {
	var a, b, c, d = 1, 2, 3, 4 //局部变量不能使用var()简写
	var a0, b0 = 100, 101       // 局部变量a0，b0 省略了类型，由系统推断
	a1, b2, c3, d4 := 1, 2, 3, 5
	var x1 byte = 1
	b3 := 2
	var x2 int32 = 3
	fmt.Println(x1, b3, x2)
	fmt.Println("a,b,c,d is ", a, b, c, d)
	fmt.Println("a1,b2,c3,d4 is ", a1, b2, c3, d4)
	fmt.Println("a0, b0 is ", a0, b0)
	test()
	var a5 float32 = 100.1 //类型转换
	b5 := int(a5)
	fmt.Println(b5)

	var a6 int = 65  //大A的ASCII码是65
	b6 := string(a6) //string转换成文本，65表示A，所以b6就是A
	fmt.Println(b6)

	var a7 int = 65        //
	b7 := strconv.Itoa(a7) //使用Itoa方法，b7=65
	fmt.Println(b7)

	b8, _ := strconv.Atoi(b7)
	fmt.Println(b8)

}
