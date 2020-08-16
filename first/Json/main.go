package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "Ivan",
		Age:      999,
		Birthday: "1999 - 09 - 09",
		Sal:      4500.0,
		Skill:    "ball",
	}
	//序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误， err = %v", err)
	}
	fmt.Printf("monster 序列化后 = %v\n", string(data))
}

// map序列化
func TestMap() {
	//定义map
	var a map[string]interface{}
	//make map
	a = make(map[string]interface{})
	a["name"] = "孙悟空"
	a["age"] = 500
	a["address"] = "花果山"
	//序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误， err = %v", err)
	}
	fmt.Printf("Map 序列化后 = %v\n", string(data))
}

//切片序列化
func TestSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "json"
	m1["age"] = "7"
	m1["address"] = "坪山"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "asjf"
	m2["age"] = "999"
	m2["address"] = "贵州"
	m2["hobby"] = "吹牛"
	slice = append(slice, m2)

	data1, err1 := json.Marshal(slice)
	if err1 != nil {
		fmt.Printf("序列化切片错误 err = %v", err1)
	}

	fmt.Printf("切片序列化成功  = %v\n", string(data1))
}

//基本数据类型序列化
func TestFloat64() {
	var num float64 = 3453.4567
	data, _ := json.Marshal(num)
	fmt.Printf("序列化后 data = %v\n", string(data))
}
func main() {
	//调用序列化函数
	testStruct()
	TestMap()
	TestSlice()
	TestFloat64()
}
