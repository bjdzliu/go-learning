package main

import "fmt"

type Humer interface {
	sayhi()

}

type Personer interface {
	Humer
	sing(lrc string)

}
type Student struct{
	name string
	id int
}
//让Student实现sayhi方法
func (temp *Student) sayhi(){
	fmt.Printf("student %d,%s\n",temp.id,temp.name)
}

func (tmp *Student) sing(n string){
	fmt.Printf("person %v\n",n)
}

func main(){

	//humer interface have sayhi
	var i  Humer
	var p Personer
	s:=&Student{id:10,name:"mike"}
	//student struct实现了Humer和Personer的全部方法，即实现了Humer和Personer接口。可以赋值给Humer和Personer interface
/*	为什么要赋值interface？
	比如：写一个函数，参数是接口，只要是实现了接口方法的对象，都可以传递
鸭子类型
*/

	i=s
	i.sayhi()
	p=s
	p.sayhi()

	p=&Student{id:10,name:"mike"}
	p.sayhi()
	p.sing("ceshi")

	//超集可转化成子集,接口之间赋值
	i=p
	i.sayhi()

}