package main

import (
	"fmt"
	"sync"
)

/**
 *wangdt
 */

// 三个方法使用 add done wait
var wg = sync.WaitGroup{}

// 写数据
func write(ch chan int) {
	defer close(ch)

	for i := 0; i < 10; i++ {
		fmt.Println("写入数据成功")
		ch <- i
	}
	wg.Done()
}

// 读数据
func read(ch chan int) {

	for v := range ch {
		fmt.Println("读取数据成功", v)
	}
	wg.Done()
}
func main() {

	// 有缓冲通道chan int, 10
	var ch = make(chan int, 10)

	wg.Add(1)
	go write(ch)

	wg.Add(1)
	go read(ch)

	wg.Wait()
}
