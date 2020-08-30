package main

import (
	"fmt"
)

// VowelsFinder 定义接口
type VowelsFinder interface {
	FindVowels() []rune
}

//MyString 定义string类型的结构体
type MyString string

//FindVowels 实现接口中的方法
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	fmt.Printf("Vowels are %c\n", name.FindVowels())
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())
	fmt.Printf("类型断言 %T\n", v)
}
