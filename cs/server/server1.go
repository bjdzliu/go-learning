package main

import (
	"net"
	"fmt"
)

func main(){
//listen
listener,err:=net.Listen("tcp","127.0.0.1:8080")
if err!=nil{
	fmt.Println("err=",err)
	return
}

defer listener.Close()

//wait
conn,err:=listener.Accept()
if err!=nil{
	fmt.Println("err=",err)
}

//recv request
buf:=make([]byte,1204)
//put in buf ; n is length
n,err:=conn.Read(buf)
if err!=nil{
	fmt.Println(err)
	return
}

fmt.Println("buf=",string(buf[:n]))
}

