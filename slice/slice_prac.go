package main


import (
	"fmt"
	"time"
)

func main() {
	var s1 []int //声明slice
	var s_string []string
	fmt.Printf("s_string",s_string)
	fmt.Println("s1",s1)

	a := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11}  //数组定义
	//a2:=[...]int{2,2,2,2,2} //数组定义
	s2 := a[5:10] //切片
	fmt.Println("a array's slice is",s2)
	fmt.Println("s2'len :",len(s2),"s2's cap :",cap(s2))



	//错误演示
	var x []int
//	go func(){x = make([]int,10)}()
	go func(){x=make([]int,1000000)}()
	time.Sleep(5*time.Second)
	x[99999]=1
	fmt.Println(x[99999])

	s3 := a[5:] //切片，取到结尾 。区别于[:5]
	fmt.Println(s3)
	s4 := make([]int, 3, 10)      //初始分配10个内存地址。以10递增
	fmt.Println("s4 len:",len(s4), "s4 cap :",cap(s4)) //cap容量长度
	s5 := make([]int, 3, 6)
	fmt.Printf("%p\n", s5)
	s5 = append(s5, 1, 2, 3)
	fmt.Printf("%v %p\n", s5, s5) //没有超过容量，地址不变
	s5 = append(s5, 1, 2, 3)
	fmt.Printf("%v %p\n", s5, s5) //超过容量，地址改变

	scores := new([]int)  //返回一个pointer
	//same as var scores *[]int = new([]int)
	fmt.Println("score pointer's value is ",*scores)

	var v1  []int = make([]int, 100)
	fmt.Println("v1:",v1)  //v1 is slice  len(v1)=100

	var v2 []int
	//fmt.Printf(v2) s2==nil

	var v3= []int{}
	fmt.Printf("%T,%T",v3,v2)



	p2 := new([10]int)

	fmt.Println("p2 is ",p2)


	//多个slice
	b := []int{1, 2, 3, 4, 5}
	s6 := b[2:5]
	s7 := b[1:3]
	fmt.Println("s6 cap:",cap(s6),s6, s7)
	s6 = append(s6, 1, 2, 1, 2, 3, 4)
	s6[0] = 99 //s6有新的地址
	fmt.Println(s6, s7)

	slice := []string{10: "1"}
	fmt.Println(len(slice), slice)

	//reslice
	slice3 := []int{1, 2, 3, 4}
	slice4 := slice3[1:4]
	fmt.Println(slice3, slice4)

}
