package main

/*
import package external_package
import package channel ,

 */

import (
	"fmt"
	"sort"
	."helloworld/external_package" //(operate var exter_package )
	"time"
	"helloworld/external_p2" //(operate var exter_package  )
)

/*
在package external_p2 中也导入了 helloworld/external_package，也对包内的变量操作：exter_package
结论：包内的变量是共享的exter_package

 */
func main() {

	go func(){
		Out()
		fmt.Println("in first GR:")
		Setvalue()
		Out()
	}()

	go func(){
		Setvalue()
Out()
	}()
	time.Sleep(10*time.Second)

	external_p2.Output2()

	time.Sleep(10*time.Second)

	var m map[int]string = make(map[int]string)
	// 长度99
	m2 := make(map[int]string, 99)
	m[1] = "ok"
	fmt.Println(m)
	fmt.Println(m2)
	//通过key删除
	delete(m, 1)
	fmt.Println(m)
	Out()

	// 两层map嵌套，初始化,m3的value是map
	var m3 map[int]map[int]string
	//
	m3 = make(map[int]map[int]string)
	a, ok := m3[2][1]
	if !ok {
		//初始化
		m3[2] = make(map[int]string)
	}
	m3[2][1] = "ok"
	a = m3[2][1]
	fmt.Println(a, ok)

	//存放map的slice,slice长度是5
	sm := make([]map[int]string, 5)
	//忽略key的值，对slice迭代
	for _, v := range sm {
		//v是map，对v初始化，才可以赋值
		//或者，达到更改slice的目的 	sm[i]=make(map[int]string)
		//		sm[i][1]="ok"
		v = make(map[int]string)
		v[1] = "ok"
		v[2] = "len2"
		fmt.Println(v)

	}
	//v是一个copy，不会影响slice
	fmt.Println(sm)

	// sort map ms
	//key放到slice里，然后排序
	ms := map[int]string{1: "a", 2: "b", 99: "c"}
	slicex := make([]int, len(ms))
	i := 0
	for k, _ := range ms {
		slicex[i] = k
		i++
	}
	sort.Ints(slicex)
	fmt.Println(slicex)

}
