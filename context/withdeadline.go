package main

import (
	"context"
	"fmt"
	"time"
	"strings"
	"reflect"
	"strconv"
)

func main() {
	hello:="hello"
	world:="world"
	str:=fmt.Sprintf("%s%s", hello, world)
	fmt.Println(str)
	join_connected:= strings.Join([]string{hello, world}, " ")
	fmt.Println(join_connected)
	fmt.Printf()

	fmt.Println(reflect.TypeOf(hello))

	string1:=strconv.FormatInt(int64(10),10)
	fmt.Println(string1)

	t1:=time.Now()
	d := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))

	//ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

    fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Unix())

		 var time1 time.Time
		 time1 = <- time.After(5 * time.Second)
		 fmt.Println("time1 value is ",time1)

	fmt.Println(time.Now().Add(50 * time.Second))

	//var time2 <- chan time.Time
	//time2 = time.After(6 * time.Second)
	//fmt.Println("time2value is ",time2)




	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():

		fmt.Println(ctx.Err())
	}



	fmt.Println(t1.Sub(time.Now()))
	fmt.Println("aa:" ,t1.Before(time.Now()))


	tick := time.Tick(1 * time.Minute)
	for _ = range tick {
		// do something
	}

	time.AfterFunc(5 * time.Minute, func() {
		fmt.Printf("expired")
	})



}
