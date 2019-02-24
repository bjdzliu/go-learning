package main

import (
	"fmt"
)

func main() {
	c := make(chan bool) //缓存为0, 读取在写入前
	go func() {
		fmt.Println("ggggg")
		c <- true //存
	}()
	<-c //读取操作，阻塞，等待“存“

	// channel 可以迭代
	b := make(chan bool)
	go func() {
		fmt.Println("dei dai")
		b <- true
		close(b) //如果有迭代，记得关闭

	}()
	for v := range b { //循环读取chan，直到close(b)
		fmt.Println(v)
	}

}
