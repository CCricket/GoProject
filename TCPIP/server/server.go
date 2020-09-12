package main

import (
	"fmt"
	"io"
	"net" //做网络socket开发时，net包含有我们需要的所有方法和函数
)

func Process(conn net.Conn) {
	//关闭conn
	defer conn.Close()

	for {
		//创建一个新切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送消息
		//如果客户端没有write【发送】，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 输入 \n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端已退出")
			return
		}
		//显示客户端发送的内容到服务端的终端
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听...")
	//TCP：表示使用的网络协议是tcp
	//0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()

	//循环等待客户端链接   处理多个不同客户端链接
	for {
		fmt.Println("等待客户端来链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
		} else {
			fmt.Printf("Accept() suc  conn = %v 客户端IP = %v\n", conn, conn.RemoteAddr().String())
		}
		go Process(conn)
	}
	//fmt.Printf("listen suc = %v\n", listen)
}
