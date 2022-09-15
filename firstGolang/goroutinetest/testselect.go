package main

//select多路复用
//不需要关闭channel

/**
 *wangdt
 */

import (
	"fmt"
	"time"
)

func main() {

	inta := make(chan int, 5)
	for in := 0; in < 5; in++ {
		inta <- in
	}

	stringa := make(chan string, 5)
	for is := 0; is < 5; is++ {
		stringa <- "hello" + fmt.Sprintf("%d", is)
	}

	for {
		select {
		case v := <-inta:
			fmt.Println("从 inta中读取数据", v)
			time.Sleep(time.Millisecond * 50)
		case s := <-stringa:
			fmt.Println("从 stringa中读取数据", s)
			time.Sleep(time.Millisecond * 50)

		default:

			fmt.Println("执行完毕")
			return //注意退出
		}

	}
}
