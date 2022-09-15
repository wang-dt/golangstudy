package main

import "fmt"

/**
 *wangdt
 */
// channel单向通道
// chan<-只写通道
func test(intc chan<- int) {

	for i := 0; i < 10; i++ {
		intc <- i
	}
}

func tests(intb <-chan int) {

	for i2 := range intb {
		fmt.Println(i2)
	}
}

func main() {

	//	只能写
	intc := make(chan<- int, 5)
	test(intc)

	//只能读
	intb := make(<-chan int, 5)
	tests(intb)

}
