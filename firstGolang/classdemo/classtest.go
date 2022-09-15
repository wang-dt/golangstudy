package main

import "fmt"

/**
 *wangdt
 */
// Hello 类名的首字母大写 其他的包可以访问
// Getname 方法名的首字母大写 其他的包可以访问
type (
	Hello struct {
		name    string
		age     int
		address string
	}
)

//(this *Hello)

// this 可以任意
//(this Hello)数据copy
//(this *Hello)内存指向

func (this *Hello) getname(string) string {

	fmt.Println(this.name)
	return this.name
}

func (this *Hello) setname(name string) string {
	this.name = name
	return this.name
}

func (this *Hello) show() {
	fmt.Println("name = ", this.name)
	fmt.Println("age = ", this.age)
	fmt.Println("address = ", this.address)
}
func main() {
	h := Hello{name: "zhangsan", age: 18, address: "tongzhou"}
	h.show()
	fmt.Println("-------------")
	h.getname("lisi")
	h.show()
	h.setname("wangwu")
	h.show()
}
