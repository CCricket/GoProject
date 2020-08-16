package main

import (
	"fmt"
)

//计算 1-200的各个数的阶乘，并且把各个数的阶乘放到map中
//最后显示出来，要求使用goroutine完成

//思路
//1、编写一个函数，来计算各个数的阶乘，并放入到map中
//2、启动多个协程，统计的将结果放入到map中
// 3、map应该是全局的

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println(v)
		sum += v
	}
	c <- sum
	fmt.Println("sum = ", sum)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)

	go sum(s[len(s)/2:], c)

	fmt.Printf("len = %v  Cap = %v", len(c), cap(c))
	y := <-c
	x := <-c
	fmt.Println(x, y, x+y)
}
