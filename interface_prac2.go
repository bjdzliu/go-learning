package main

import (
	"fmt"
)

type empty interface {
}

//定义一个接口
type USB interface {
	Name() string //Name()方法，返回string
	Connecter
}

//interface嵌套，但是不能递归嵌套
type Connecter interface {
	Connect()
}

type PhoneConnectter struct {
	name string
}

//为struct创建方法,实现了USB接口中Name()的方法
func (pc PhoneConnectter) Name() string {
	return pc.name
}

//为struct创建方法,实现了Connecter接口中的Connect方法
//内部方法，访问内部变量
func (pc PhoneConnectter) Connect() {
	fmt.Println("connected", pc.name)
}

//TVConnecter 无法转换成USB,没有实现USB中Name()的方法
type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("connected", tv.name)
}
func main() {
	//a实现了USB接口，可以转化为Connecter接口
	pc := PhoneConnectter{"initName"}
	var a Connecter
	//pc值拷贝
	a = Connecter(pc)
	a.Connect()
	pc.name = "ceshi"
	//pc的name值，并没有修改a对象中name的值
	a.Connect()

	//tv := TVConnecter{"tvconnect"}
	//var a USB
	//a=USB(tv)
	//a.connect()

	//只有当接口存储的类型和对象都是nil时，接口才是nil
	var testnil interface{}
	var p *int = nil //类型是指针
	testnil = p
	fmt.Println(testnil == nil)
	fmt.Println(p == nil)

}

//没有显示声明实现空接口
//usb.(PhoneConnectter)判断usb(a)是否是PhoneConnectter类型

func Disconnect(usb interface{}) { //实现空接口
	//pc是转换后的结果。如果ok是false，则pc为nil或空

	if pc, ok := usb.(PhoneConnectter); ok { //简单的类型断言
		fmt.Println("disconnect", pc.name)
		return
	}
	switch v := usb.(type) {
	case PhoneConnectter:
		fmt.Println("Disconnected", v.name)
	default:
		fmt.Println("Unknown")
	}

	fmt.Println("unknow device")
}
