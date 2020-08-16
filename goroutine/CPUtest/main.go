package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU() // 逻辑ＣＰＵ
	fmt.Println("cpuNum = ", cpuNum)
	//设置num - 1个CPU运行go程序
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("num = ",cpuNum)

}
