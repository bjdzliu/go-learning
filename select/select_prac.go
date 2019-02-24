package main

import (
	"fmt"
	"time"
)
func main(){
	c:=make(chan int)
	o:=make(chan bool)

	go func(){
		for {
			select{
			case v:=<-c:
				fmt.Println(v)
			case <-time.After(5*time.Second):
				fmt.Println("timeout")
			o<-true
			break
			}
		}

	}()
c<-666
<-o  //读通道o的数据，没有数据时，阻塞

}
