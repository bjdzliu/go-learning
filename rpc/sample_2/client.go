package main

import (
	"os"
	"fmt"
	"strconv"
	"net/rpc"
	"log"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {

	//判断参数长度
	if len(os.Args) != 4 {
		fmt.Println("parameters is not enough", os.Args[0])
	} else {
		fmt.Println("length is" + strconv.Itoa(len(os.Args)))
	}
	serverAddress := os.Args[1]
	fmt.Println("server addr is: ", serverAddress)

	//初始化一个rpc连接
	client, err := rpc.DialHTTP("tcp", serverAddress)

	if err != nil {
		log.Fatal("error happen", err)
	}
	//准备参数
	i1,_:=strconv.Atoi(os.Args[2])
	i2,_:=strconv.Atoi(os.Args[3])
	argsTwo:=ArgsTwo{i1,i2}

	var reply int
	//连接服务器
	err = client.Call("Arith.Multiply",argsTwo,&reply)
	if err!=nil{
		log.Fatal("error",err)
	}
	fmt.Println("Arith is",argsTwo.A,argsTwo.B,reply)

	i3,_:=strconv.Atoi(os.Args[2])
	i4,_:=strconv.Atoi(os.Args[3])
	QuotientTwoc:=QuotientTwo{i3,i4}

	var reply2 QuotientTwo

	err=client.Call("Arith.Divide",argsTwo,&reply2)
	if err!=nil{
		log.Fatal("error",err)
	}
	fmt.Println("Arith % is",QuotientTwoc.Quo,QuotientTwoc.Rem,reply2.Quo,reply2.Rem)

}
