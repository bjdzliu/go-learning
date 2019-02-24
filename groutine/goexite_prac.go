package main

import (
	"fmt"
	"runtime"
)

func test(){
	defer fmt.Println("cccccc")
	//终止所在的协程
	runtime.Goexit()
	fmt.Println("ddddddd")
}

func main(){
go func(){
fmt.Println("aaaaa")
fmt.Println("bbbbbbbb")
}()

	for{}
}
