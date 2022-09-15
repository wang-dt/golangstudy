package main

import "fmt"

/**
 *wangdt
 */

// 切片
func main() {

	//声明一个slice
	//slice1 := []int{1, 2, 3}

	//声明一个切片 但是没有给slice分配空间
	//var slice1 []int

	//声明一个切片 用make 给slice分配空间 //[]int为切片类型,初始化值为0
	//var slice1 = make([]int, 2)

	//声明一个切片 用make给slice分配空间 //[]int为切片类型
	slice1 := make([]int, 4)

	fmt.Printf("len = %d slice = %v\n", len(slice1), slice1)

	if slice1 == nil {
		fmt.Println("kongqie")
	} else {
		fmt.Println("youzhi")
	}
}
