package main

import (
	"os"
	"fmt"
	"io"
)

func main(){
	list:=os.Args
	if len(list)!=3{
		fmt.Println("src and dst file missing")
	}
	srcfile:=list[1]
	dstfile:=list[2]
	fmt.Printf(list[0])

	if srcfile==dstfile{
		fmt.Printf("they are same")
		return
	}

	src1,err:=os.Open(srcfile)
	if err!=nil{
		fmt.Println("err",err)
		return
	}
	//新建目的文件
	df,err2:=os.Create(srcfile)
	if err!=nil{
		fmt.Println("err2",err2)
		return
	}
	//关闭文件
	defer src1.Close()
	defer df.Close()
	buf:=make([]byte,1)
	for{
		n,err:=src1.Read(buf)
		fmt.Println("读取一次")
		if err!=nil{
			fmt.Println("err",err)
			if err==io.EOF{
				break
			}
		}
		//df.WriteString(string(buf))
		df.WriteString(string(buf[:n]))
	}




}
