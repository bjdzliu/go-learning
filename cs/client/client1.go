package main

import (
	"net"
	"fmt"
)

func main(){
	//
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	//close
	defer conn.Close()
	//send  写一个分片
	conn.Write([]byte("are you ok"))

}