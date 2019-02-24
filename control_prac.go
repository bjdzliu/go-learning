package main

import (
	"fmt"
)

func main() {
	a := 1
	if a, b := 1, 2; a > 0 { //初始化变量的作用域
		fmt.Println(a, b)
	}
	fmt.Println(a)

	for { //for语句
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a)

	a2 := 1
	switch a2 { //case 表达式是一个值

	case 0:
		fmt.Println("a2=0")
		fallthrough
	case 1:
		fmt.Println("a2=1")
	default: //如果case都不符合
		fmt.Println("None")
	}

	a1 := 5
	switch { // 使用case表达式

	case a1 >= 0:
		fmt.Println("a1=5")
		fallthrough
	case a1 >= 5:
		fmt.Println("a1=5")
	default: //如果case都不符合
		fmt.Println("None")
		goto LABEL2
	}

	/*
		switch a3 :=1{

		}
	*/

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1 //跳出LABEL1的循环
				/*
					goto LABEL1 跳到LABEL1，还是无限循环
					建议时候 goto LABEL2
				*/
			}

		}
	}

LABEL2:
	fmt.Println("over")

	/*
	   continue 直接进入下一次的循环
	   ocntinue LABEL1 继续到LABEL1 执行
	*/

LABEL3:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL3 //这里的无限循环不会执行，跳到外面执行有限循环
			// continue改成goto，程序编程无限循环
			// continue还在外层循环当中，但是goto是重新执行外层循环。
			fmt.Println(i) //这句不会被执行
		}

	}

}
