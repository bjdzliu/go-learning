package main

import (
	"bufio"
	"os"
	"fmt"
)
/*
优点：
利用了缓冲

待完善的地方：
写入的字符如果超过4096，需要利用for 循环写入

*/
func main() {
	//func Copy(dst Writer, src Reader) (written int64, err error)
	path := "./file/write_prac1"
	fd, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)

	// 实现一个带缓冲的写文件
	w := bufio.NewWriterSize(fd, 4096)

	fmt.Println(w.Available()) //表示可使用的缓冲区的大小
	fmt.Println(w.Buffered())  //表示已经缓冲的数据的大小

	w.WriteString("hello world!\n")

	w.Write([]byte("hello world!\n"))

	fmt.Println(w.Available()) //表示可使用的缓冲区的大小
	fmt.Println(w.Buffered())  //表示已经缓冲的数据的大小

	//此时需要刷新一下数据才能写入文件
	w.Flush()
}