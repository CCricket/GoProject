package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "/home/ivan/home/abc.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("open file err = %v", err)
		return
	}

	defer file.Close()

	str := "追，Ivana\r\n"

	writer := bufio.NewWriter(file)

	for i := 0; i < 6; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
