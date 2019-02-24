package main

import "fmt"

const a int = 1
const b = 'A' //隐式类型判断

const (
	_MAX_COUNT = 'a' //编码规范
	c          = a
	d          = a + 1
	e          = a + 2
)

const f, g, h = 1, 2, 3

const (
	a3 = 'A'
	b4
	c5 = iota //该const里的第三个值，c5为2
	d6        //d6为3
)

const (
	a1, b1 = 1, 4
	a2, b2 //对应上一行的a1，b1
)

func main() {
	fmt.Println(a2, b2, d6)
}
