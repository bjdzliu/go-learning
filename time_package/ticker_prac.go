package main

import (
	"time"
	"fmt"
)

//d定时触发


func main(){
	ticker:=time.NewTicker(time.Second*2)
//创建定时器，每隔1秒，定时
i:=0

	for{
	<- ticker.C
	i++
fmt.Println(i)

	}



}
