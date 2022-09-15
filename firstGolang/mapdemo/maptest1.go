package main

import (
	"fmt"
)

/**
 *wangdt
 */
func main() {

	var map1 = map[string]string{}

	if map1 != nil {
		fmt.Println(map1)
	} else {
		fmt.Println("kong de ")
	}

	map1 = make(map[string]string, 7)
	//1.1.9 不需要make开辟空间了

	fmt.Println(map1)

	map1["st"] = "zhangsan"
	map1["sp"] = "lisi"

	fmt.Println(map1)

	namemap(map1)
}
func namemap(map1 map[string]string) {

	fmt.Println("-------------")
	fmt.Println(map1)
	fmt.Println("-------------")

	map1["wu"] = "wangwu"
	map1["wu"] = "jjsihh"

	delete(map1, "sp")
	for s, s2 := range map1 {
		fmt.Println(s)
		fmt.Println(s2)
	}
}
