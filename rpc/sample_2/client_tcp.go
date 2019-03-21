package main

import (
	"fmt"
	"os"
	"net/rpc"
	"log"
	"strconv"
)



func main(){
	fmt.Println("call rpc server")
	if len(os.Args) !=4{
		fmt.Println("parameters not enough")
		os.Exit(1)
	}
	service:=os.Args[1]
	client,err:=rpc.Dial("tcp",service)
	if err!=nil{
		log.Fatal("dial error")
	}
	num1:=os.Args[2]
	i1,error1:=strconv.Atoi(num1)
	if error1!=nil{
		fmt.Println("get num1 error",error1)
		os.Exit(1)
	}
	num2:=os.Args[3]
	i2,error2:=strconv.Atoi(num2)
	if error2!=nil {
		fmt.Println("get num2 error",error2)
		os.Exit(1)
	}
aa:=Putnum{i1,i2}
var reply int
err1:=client.Call("Num.Method1",aa,&reply)

	if err1!=nil{
		log.Fatal("jisuan shibai",err1)
	}
	fmt.Println("result is",reply)
}


type Putnum struct{
	A,B int
}