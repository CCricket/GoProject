package main

import (
	"fmt"
	"strconv"
	"time"
)

/*MPG :
M表示操作系统的主线程（物理线程）
P表示协程执行需要的上下文
G表示协程
*/
//在主线程（可以理解为进程）中，开启一个goroutine，该协程每隔1s输出"" hello world"
//在主线程中也每隔1s输出""hello golang"，输出10次后，退出程序
//要求主线程和goroutine同时执行

//创建函数，每秒输出helloworld
func TestGoroutine() {

	for i := 1; i <= 12; i++ {
		fmt.Println("TestGoroutine---> hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func main() {

	go TestGoroutine() //go 表示开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main---> hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	//go TestGoroutine() //主线程执行完后就不会执行此函数

}
