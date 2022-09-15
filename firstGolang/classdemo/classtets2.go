package main

import "fmt"

/**
 *wangdt
 */
type Human struct {
	name string
	sex  string
}

func (receiver *Human) Eet() {
	fmt.Println("Human.Eet")
}

func (receiver *Human) Work() {
	fmt.Println("Human.Work")
}

type Supman struct {
	Human //Supman 继承了Human类的方法
	level int
}

// 重写Human的Eet方法
func (receiver *Supman) Eet() {
	fmt.Println("Supman.Eet")

}

// 新方法
func (receiver *Supman) Fly() {
	fmt.Println("Supman.Fly")
}

func (receiver *Supman) Print() {
	fmt.Println("nane = ", receiver.name)
	fmt.Println("sex = ", receiver.sex)
	fmt.Println("level = ", receiver.level)

}

func main() {

	var h Human
	h.name = "张三"
	h.sex = "男"

	h.Eet()
	h.Work()

	var s Supman
	s.name = "超人"
	s.sex = "男"
	s.level = 88

	s.Eet()
	s.Work()
	s.Fly()
	s.Print()
}
