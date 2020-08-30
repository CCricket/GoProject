package main

import "fmt"

// 向 intChan 放入1-8000个数
func PutNum(intChan chan int) {

	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

//判断是不是素数
func PrimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环判断是不是素数
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断num 是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数放到primeChan管道里面去
			primeChan <- num
		}
	}

	fmt.Println("有一个协程退出")

	// 向exitChan 写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	//存放结果
	primeChan := make(chan int, 2000)
	//退出管道标识
	exitChan := make(chan bool, 4)

	//开启一个协程，向intChan放入需要计算的数
	go PutNum(intChan)

	//开启4个协程
	for i := 0; i < 4; i++ {
		go PrimeNum(intChan, primeChan, exitChan)
	}

	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	//遍历结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数 = %v\n", res)
	}
}
