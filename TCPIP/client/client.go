package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	//客户端可以发送单行数据，然后输出
	//os.Stdin表示标准输入
	reader := bufio.NewReader(os.Stdin)

	//从终端读取一行用户输入，并准备发送给服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read String err = ", err)
		}
		fmt.Println("line =", line)
		//再将line发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err ", err)
		}
		//fmt.Printf("line11 =%v,type = %T\n", line, line)
		//fmt.Printf("客户端发送 %d 字节数据，并退出\n", n)
		line = strings.Trim(line, "\n")
		if line == "exit" {
			break
		}
	}
	fmt.Printf("client conn = %v\n", conn)
}
