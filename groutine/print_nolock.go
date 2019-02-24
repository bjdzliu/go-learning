package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _,data:=range str{

fmt.Printf("%s",string(data))
	time.Sleep(time.Second)
}
fmt.Println()
}

var channel=make(chan int)

func person1(){

	Printer("hello")
	channel <- 5
}

func person2(){
	<- channel
	Printer("world")
}

func main(){
	//Printer("hello")
	//Printer("world")

go person1()
go person2()

for{}



}