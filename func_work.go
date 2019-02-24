//分析运行结果
package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//i是值拷贝
		defer fmt.Println("defer i =", i)
		//i 是地址拷贝
		defer func() { fmt.Println("defer_closure i=", i) }()
		//func()里面的i 是地址拷贝
		fs[i] = func() { fmt.Println("closure i=", i) }

	}
	// for语句
	for _, f := range fs {
		f()
	}
}
