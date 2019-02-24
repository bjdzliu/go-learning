package main

import (
	"fmt"
	"runtime"
)


func main(){


	n:=runtime.GOMAXPROCS(2)

	//n:=runtime.GOMAXPROCS(1)
	fmt.Println("n=%d",n)
	for{
		go fmt.Println(0)
		fmt.Println(1)
	}




}
