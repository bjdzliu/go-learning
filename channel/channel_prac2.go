package main

import (
	"time"
	"fmt"
)
type Cake2 struct{state string}
func main(){
	icechan:=make(chan Cake2,1)
	go icer2(icechan)

	go func(){

/*		x,ok:=<- icechan
		if !ok{
			fmt.Println("no vaule")
		}*/
		icechan <- Cake2{"okfunc" }
		a:=1+3
		fmt.Println("get value from chan",a)
		<-icechan
	}()
	time.Sleep(5*time.Second)
}

func icer2(iced chan Cake2){
	iced <- Cake2{"ok" }
	b:=2+5
	fmt.Println("b value is",b)
	x:=<-iced //释放掉
	fmt.Println("after release",x.state)


}