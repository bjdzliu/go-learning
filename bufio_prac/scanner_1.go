package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main(){
	f,_:=os.Open("./bufio_prac/txt")
	buf := make([]byte,1024)
	n,_ := f.Read(buf)
	input:=string(buf[:n])
	fmt.Printf("input is %s",input)
	scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner类型中有spit 函数 【split        SplitFunc // The function to split the tokens.】
	scanner.Split(bufio.ScanWords)  //ScanWords is a split function 返回以空格分隔的字节流
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}