package main

import (
	"fmt"
)

// 定义变量
var (
	name = "zhangsan"
	age  = 18
)

// 定义常量
const (
	BEIJING  int = 1
	SHANGHAI int = 2
	SHENZHEN int = 3
)

//iota 只能出现在const中

const (
	d = iota
	b = 8
	c = iota + 1
)

//main函数

// (a string, bool2 bool)=== 参数体
//(r1, r2 int) ===返回体

func fool(a string, bool2 bool) (r1, r2 int) {
	fmt.Println(a)
	fmt.Println(bool2)

	r1 = 100
	r2 = 200

	return
}

func main() {
	var a int
	var ud int
	var du int
	var sy string = "zhangsan"
	e := 3.15
	fmt.Printf("type of a  = %T\n", a)
	fmt.Printf("type of sy  = %T\n", sy)
	fmt.Printf("type of e = %T\n", e)
	fmt.Printf("type of e = %T\n", name)
	fmt.Println(name, age)

	fmt.Println("beijing+ ", BEIJING)
	fmt.Println("shanghai+ ", SHANGHAI)
	fmt.Println("shenzhen+ ", SHENZHEN)
	fmt.Println("a+ ", d)
	fmt.Println("b+ ", b)
	fmt.Println("c+ ", c)

	ud, du = fool("zhangsan", true)

	fmt.Println(ud, du)

}
