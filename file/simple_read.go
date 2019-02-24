package main

/*
NewReader
NewScanner

 */

import (
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
	"io"
	"strings"
)


func main(){
	//filename:="citic-activity-promotion.log"


	f,err:=os.Open("citic-activity-promotion.log")
	if err != nil{
		fmt.Errorf("read file failed %s",err)
	}


//method1(f)
//method2(f)
method3(f)
//method4(f)
//method5(filename)
}

//输出file中的重复行
func method5(f string){
	counts := make(map[string]int)
	data,err:=ioutil.ReadFile(f)
	if err !=nil{
		fmt.Fprintf(os.Stderr,"method5:%v\n",err)
	}
for _,line:=range strings.Split(string(data),"\n"){
	counts[line]++
}

for line,n:=range counts{
	if n>1 {
		fmt.Printf("%d\t%s\n",n,line)
	}
	}
}

func method4(f *os.File){
	fmt.Println("NEW**************************************************")
	fd,err := ioutil.ReadAll(f)
	if err!=nil{
		fmt.Errorf("errinfo",err)
	}
	fmt.Println(string(fd))

	defer f.Close()
}

func method3(f *os.File){
		input:=bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
		fmt.Println(input.)
	}
fmt.Println("second read is null")
	input2:=bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input2.Text())

	}
}


func method2(f *os.File){


		r := bufio.NewReader(f)
		tempbuff:=make([]byte,4096)
		 for{
			 stringdata,err := r.ReadString('\r')
			 fmt.Println(stringdata)
			 readNum, err := r.Read(tempbuff)
			 if err!=nil && err != io.EOF{
				 fmt.Errorf("echo:",err)
			 }
			 if readNum == 0 {
				 fmt.Println("end")
				 break
				 }
		 }


}

func method1(f *os.File){

	b := make([]byte,4096)
		for{
			 n1,err:=f.Read(b)
			 if err!=nil{
			 }
			 fmt.Println(string(b))
			 fmt.Println("#############",n1)
			 if 0 ==n1 {break}
		 }


}