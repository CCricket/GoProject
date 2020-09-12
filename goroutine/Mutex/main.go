package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var x = 0
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func incrementChan(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	//var m sync.Mutex
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementChan(&w, ch)
		//go increment(&w, &m)
	}
	w.Wait()
	endTime := time.Now()
	fmt.Println("use times is :",endTime.Sub(startTime))
	fmt.Println("final value of x:", x)
	fmt.Println("Max CPU :", runtime.NumCPU())
}

