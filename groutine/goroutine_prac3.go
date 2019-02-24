package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 10) //缓存为10
	go func() {
		fmt.Println("ggggg")

		<-c //缓存为0时，读操作一直等待一个写操作 (阻塞的)

	}()

	c <- true //缓存为10时,这里写操作完成，直接结束

}
