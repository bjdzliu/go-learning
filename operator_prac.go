package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
)

func main() {
	fmt.Println(1 << 10) // 左移10位

	a := 1
	var p *int = &a //定义指针p
	fmt.Println(*p)
}

/*
 6:0110
11:1011
-------
&  0010
|  1111
^  1101
&^ 0100   第2个数是1，强制将第一个数相应的位置改成0

*/
