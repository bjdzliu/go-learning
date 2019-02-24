package main

import (
	"fmt"
)

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

func main() {
	a := PhoneConnectter{"initName"}
	a.Connect()
	Disconnect(a)

}

//没有显示声明实现接口USB
//usb.(PhoneConnectter)判断usb(a)是否是PhoneConnectter类型

func Disconnect(usb USB) { //PhoneConnectter实现了usb接口

	if pc, ok := usb.(PhoneConnectter); ok { //简单的类型断言
		fmt.Println("disconnect", pc.name)
		return
	}
	fmt.Println("unknow device")
}
