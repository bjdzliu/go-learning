package main

import (
	"net"
	"log"
	"os"
	"io"
	"time"
)

func main(){
	conn,err:=net.Dial("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)

	}
	defer conn.Close()
	go mustCopy(os.Stdout,conn)
	time.Sleep(5*time.Second)
	mustCopy(conn,os.Stdin)
}
func mustCopy(dst io.Writer,src io.Reader){
	if _,err:=io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}

}