package main

import (
	"fmt"
)

func main() {
	var a [2]string
var bb int=1
fmt.Println(bb)
	//数组个数必须是常量
	//n:=10
	//var c [n]int  // 报错

	b := [2]int{1, 2}
	var b1 [20]int
	//var gg [21]int
	fmt.Printf("%d", len(b1))
	d := [20]int{19: 2} //第19个元素为2
	e := [...]int{1, 1, 2}
	f := [...]int{20: 1}
	fmt.Println(a, b, d, e, f)
	g := [...]int{10: 1} //数组g,长度为11
	var p *[11]int = &g
	fmt.Println(p, len(g))
	//指针数组,数组中保存的类型是指针

	x, y := 1, 2
	z := [...]*int{&x, &y}
	fmt.Println("z's content", z)

	h := [5]int{}
	h[1] = 2 //根据索引赋值
	fmt.Println(h)

	i := [2][3]int{ // 创建多维数组
		{1, 2, 2},
		{4, 5, 6},
	}
	fmt.Println(i)

	//冒泡排序,从小到大排序
	a10 := [...]int{1, 2, 6, 99, 55, 3, 101, 999, 55}
	fmt.Println(a10)
	num := len(a10)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a10[i] > a10[j] {
				temp := a10[i]
				a10[i] = a10[j]
				a10[j] = temp
			}

		}

	}
	fmt.Println(a10)
}
