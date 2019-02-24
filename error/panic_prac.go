package main

import "fmt"

func testa(){
fmt.Println("aaaaaa")
}

func testb(x int){
	fmt.Println("bbbbbb")
	//panic("this is panic test bbbb")
	var a [10]int
	a[x]=111
}

func testc(){
	fmt.Println("aaaccccc")
}

func main(){
	testa()
	testb(22)
	testc()
}