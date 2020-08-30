package main

import (
	"fmt"
	"sync"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Groutine", i)
	//time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		//wg.Add(1)
		go process(i, &wg)
	}
	wg.Add(3)
	wg.Wait()
	fmt.Println("All go routine finshed executing")
}
