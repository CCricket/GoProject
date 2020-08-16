package main

import (
	"fmt"
	"strings"
)

//累加器

func addUpper() func(int) int {
	var n int = 10
	fmt.Println("n=", n)
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果name 没有指定后缀， 则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {

			return name + suffix
		}

		return name
	}
}

func makeSuffix1(name string, suffix string) string {
	if !strings.HasSuffix(name, suffix) {

		return name + suffix
	}
	return name

}

func main() {
	var a1 int64 = 999999
	var a2 int8 = int8(a1)
	fmt.Printf("a1=%T,a2 =%v\n", a1, a2)

	//基本类型转string      fmt.Sprintf（）
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar = 'h'
	var str string

	//
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T  str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str Type is %T  str = %q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str Type is %T, str = %q\n", str, str)

	for i := 1; i < 4; i++ {
		f := addUpper()
		fmt.Println(f(1))
	}

	//fmt.Println(f(2))
	//fmt.Println(f(3))
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("Ivan.go"))
	fmt.Println("文件名处理后=", f2("Ivana.jpg"))

	f3 := makeSuffix1("Ivan", ".go")
	fmt.Println("经过处理后的文件名 = ", f3)
}
