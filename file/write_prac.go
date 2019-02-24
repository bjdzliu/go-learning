package main

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

/*
缺点：没有利用缓冲

*/


func check(e error){
	if e !=nil{
		panic(e)
	}
}

func checkFileIsExist(filename string) bool{
	var exist=true
	if _,err:=os.Stat(filename);os.IsNotExist(err){
exist=false
	}
	return exist
}

func main(){
	var wireteString="测试文件内容n"
	var filename="./file/citic-activity-promotion.log"

	var f *os.File
	var err1 error

	if checkFileIsExist(filename){
		//os.Open()  只读
		f,err1=os.OpenFile(filename,os.O_APPEND,0666)
		fmt.Println("文件存在")

	}else {
		f,err1=os.Create(filename)
		fmt.Println("文件不存在")
	}
	check(err1)
	n,err1:=io.WriteString(f,wireteString)
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)

//second method
var d1=[]byte(wireteString)
err2:=ioutil.WriteFile("./file./output1.txt",d1,0666)
	check(err2)

//third method
f,err3:=os.Create("./file./output1.txt")
	check(err3)
defer f.Close()
n2,err3:=f.Write(d1)

	check(err3)
	fmt.Printf("写入 %d 个字节n", n2)
	n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()


}