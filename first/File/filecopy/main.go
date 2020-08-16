package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	readFile := "/home/ivan/home/abc.txt"
	writeFile := "/home/ivan/home/ivan.txt"

	data, err := ioutil.ReadFile(readFile)
	fmt.Print(string(data))
	if err != nil {
		return
	}

	err = ioutil.WriteFile(writeFile, data, 0777)

	if err != nil {
		return
	}
	/*
			//打开文件
		readF, err := os.OpenFile(readFile, os.O_RDONLY, 0777)
		writeF, err1 := os.OpenFile(writeFile, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Open file err = %v", err)
			return
		}

		if err1 != nil {
			fmt.Printf("Open File Err  %v=", err)
			return
		}
		//关闭文件句柄
		defer readF.Close()
		defer writeF.Close()

		//获取 读取  写  文件句柄
		readbuf := bufio.NewReader(readF)
		writebuf := bufio.NewWriter(writeF)
		for {
			read, readerr := readbuf.ReadString('\n')
			if readerr == io.EOF {
				break
			}
			fmt.Print(read)
			writebuf.WriteString(read)
			writebuf.Flush()
		}
	*/
}
