package main

import (
	"fmt"
	//"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func test1() {
	for i := 0; i < 100; i++ {
		fmt.Println("协线程循环1", i)
	}
	defer wg.Done()
}
func test2() {
	for i := 0; i < 100; i++ {
		fmt.Println("协线程循环2", i)
	}
	defer wg.Done()
}
func main() {

	//打印当前主机的cpy核数
	//cpunum := runtime.NumCPU()
	//fmt.Println(cpunum)
	//
	////设置goroutine最大的协程调节器个数
	//runtime.GOMAXPROCS(3)
	//fmt.Println("----------------")

	wg.Add(1)
	go test1()

	wg.Add(1)
	go test2()
	fmt.Println("主线程循环")
	for i := 0; i < 2; i++ {
		fmt.Println("主线程循环", i)
	}

	wg.Wait()
}
