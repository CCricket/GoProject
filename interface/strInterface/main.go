package main

import (
	"fmt"
)

//Describer 接口
type Describer interface {
	Describe()
}

// Person 结构体
type Person struct {
	name string
	age  int
}

// Describe 实现该方法
func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

//Address 结构体
type Address struct {
	state   string
	country string
}

// Describe 方法
func (a *Address) Describe() {
	fmt.Printf("state %s Country %s\n", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Ivan", 24}
	d1 = p1
	d1.Describe()
	p2 := Person{"Ivana", 25}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"ShenZhen", "China"}

	//d2 = a
	d2 = &a
	d2.Describe()
}
