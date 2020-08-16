package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "/home/ivan/home/abc.txt"
	//O_WRONLY只写 ，O_TRUNC把打开文件的内容清除掉
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 06666)
	if err != nil {
		fmt.Printf("open file  err = %v", err)
		return
	}

	//	及时关闭file句柄，防止内存泄露
	defer file.Close()

	str := "你好，Ivan\r\n" //表示换行

	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，内容先写入缓存 ，再调用Flush方法，将缓存中的数据真正写入到文件中
	writer.Flush()
}
