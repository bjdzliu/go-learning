package main

import (
	"os"
	"bufio"
	"io"
	"time"
	"log"
	"fmt"
)

/*
利用 bufReader.ReadString 读取文件
 */

func readBufio(path string) {
	file, err := os.Open(path)
	//file 实现了read方法
	if err != nil {
		panic(err)
	}
	defer file.Close()
  //NewReader的参数file传递给了形参io.Reader，io.Reader是接口类型(含有read方法)，io.Reader是鸭子类型
  //NewReader 返回一个*reader(bufio里面定义的sturct，有属性，有方法）
	bufReader := bufio.NewReader(file)
  //  bufReader2 :=bufio.NewScanner(file)
	buf := make([]byte, 4096)

//读取数据的方式,每次读取4096，循环读取
for {

	string_line, err := bufReader.ReadString('\r')
	fmt.Println("readString content is", string_line)

	fmt.Fprintf(os.Stdout, "an %s\n", "flag")
	//Fprintf 输出，并将内容写入到os.Stderr

	//Read返回字节数到readNum
	//因为已经读完了，bufReader里面是空的
	readNum, err := bufReader.Read(buf)
	fmt.Println("Read content is", readNum)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if readNum == 0 { break; }
}
	}


func main() {
	//size is 26MB
	pathName := "./file/bigfile"
/*	start := time.Now()
	readCommon(pathName)
	timeCommon := time.Now()
	log.Printf("read common cost time %v\n", timeCommon.Sub(start))
*/
	readBufio(pathName)
	timeCommon := time.Now()
	timeBufio := time.Now()
	log.Printf("read bufio cost time %v\n", timeBufio.Sub(timeCommon))

/*	readIOUtil(pathName)
	timeIOUtil := time.Now()
	log.Printf("read ioutil cost time %v\n", timeIOUtil.Sub(timeBufio))*/


}



