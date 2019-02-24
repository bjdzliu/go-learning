package main

import (
	"fmt"
)

type (
	byte     int8
	rune     int32
	ByteSize int64
)

var x1 byte = 100 //不能忽略var
var x2 = 101

var (
	g1 = 2
	g2 = 5
)

// 全局变量不能使用 x3 := 102

func test() {
	x3 := 102
	var x4 int32 = 3
	fmt.Println(x1, x2, x3, x4, g1, g2)
}

func main() {
	var a, b, c, d = 1, 2, 3, 4
	a1, b2, c3, d4 := 1, 2, 3, 5
	var x1 byte = 1
	b3 := 2
	var x2 int32 = 3
	fmt.Println(x1, b3, x2)
	fmt.Println("a,b,c,d is ", a, b, c, d)
	fmt.Println("a1,b2,c3,d4 is ", a1, b2, c3, d4)
	test()

}
