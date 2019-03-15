package main

import "fmt"

func main(){
backquote:=`rqweqwee/\nw`
fmt.Println(backquote)

doubelquote:="rqweqwee/\nw" //支持转义
fmt.Println(doubelquote)
}
