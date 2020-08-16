package main

import (
	"fmt"
	"strconv"
)

func sum(n1 int, n2 int) int {

	//当执行到defer时，暂不执行，会将defer后面语句压入到独立的栈中
	//当函数执行完毕后，再从defer栈中按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1 = ", n1) // defer 3.ok1 n1 = 10
	defer fmt.Println("ok2 n2 = ", n2) // defer 2.ok2 n2 = 20

	res := n1 + n2
	fmt.Println("ok3 res = ", res) // 1.ok3 res = 30
	return res

}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) // 4.res = 30
	fmt.Println(f())
	fmt.Println(f1())
	fmt.Println(f2())

	str := "Ivan凤静"                  //一个中文占用3个字节
	fmt.Println("str.len", len(str)) //len()按字节遍历
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("r[%d]:%c\n", i, r[i])
	}

	Ivan := strconv.Itoa(123)
	fmt.Printf("string is %v", Ivan)

}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
