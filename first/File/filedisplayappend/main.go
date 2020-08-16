package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "/home/ivan/home/abc.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("open file err = %v", err)
		return
	}

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {

			break
		}
		fmt.Print(str)
	}
	defer file.Close()

	str := "Ivan 继续追 Ivana \r\n"

	writer := bufio.NewWriter(file)

	for i := 0; i < 6; i++ {
		writer.WriteString(str)

	}

	writer.Flush()

}
