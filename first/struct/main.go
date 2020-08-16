package main

import (
	"fmt"
)

//人类
type Person struct {
	Name  string
	Age   int
	Skill string
	Pt    *int
}

func (a Person) learn() {
	a.Name = "Ivana"
	fmt.Println("...", a.Name)
}
func main() {
	var p Person
	p.Name = "Ivan"
	p.learn()
	fmt.Println(p)
}
