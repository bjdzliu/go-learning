package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func writeFile(path string){
f,err:=os.Create(path)
if err!=nil{
	fmt.Println("err=",err)
return
}
defer f.Close()

var buf string
for i:=0;i<10;i++{
	//得到一个字符串，字符串放到buf中
	buf=fmt.Sprintf("i=%d\n",i)
	//将buf写入文件
	fmt.Println("buf=",buf)
	// n返回字节长度和错误
	n,err:=f.WriteString(buf)
	if err!=nil{
		fmt.Println("err")
		return
	}
	fmt.Println("n=",n)
}

}


func Read(path string){
	f,err:=os.Open(path)
	if err!=nil{
		fmt.Println("err")
		return
	}
	//关闭文件
	defer f.Close()
buf:=make([]byte,1024)  //按照2k大小读取
// n 返回文件内容字节的长度
//文件只读取一次
n,err:=f.Read(buf)
	fmt.Println("只读取一次")
	if err!=nil{
		fmt.Println("err")
		return
	}
fmt.Println("buf",string(buf[:n]))


}

//每次读取一行
func Readeveryline(path string){
	//打开文件
	f,err:=os.Open(path)
	if err!=nil{
		fmt.Println("err")
		return
	}
	defer f.Close()
	//新建一个缓冲区，文件全部放在缓冲区
	r:=bufio.NewReader(f)
	for {
		//按行读，遇到/n算一行
		buf,err:=r.ReadBytes('\n')
		if err!=nil{
			//遇到结尾，退出循环
			if err ==io.EOF{
				break
			}
		}
		fmt.Printf("%v",string(buf))
	}


}

//读大文件

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := readLine(r)
		if err != nil {
			break
		}
	}

}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func main(){
	path:="./file_operatoin/demo.txt"
	//writeFile(path)
	//fmt.Println("###########")
	//Read(path)
	ReadLine(path)
}
