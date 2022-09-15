package main

import (
	"fmt"
)

func main() {

	var numbers = make([]int, 3, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 4)
	numbers = append(numbers, 6)
	fmt.Printf("len = %d,cap = %d", len(numbers), cap(numbers))
	fmt.Println(numbers)

	var sy = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sy)

	s1 := sy[0:3] // 1 2 3

	fmt.Println(s1)

	sy[2] = 100 //1 2 100 4 5 6

	fmt.Println(sy)

	//将后面的copy到前面变量

	copy(s1, sy)    // 可以将底层slice进行capy 值会进行替换
	fmt.Println(s1) // 1 2 100
	fmt.Println(sy) // 1 2 100 4 5 6

}
