package main

import "fmt"

//管道遍历

// for循环读取管道没问题
// for
func main() {

	var ch1 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}

	// 使用 range 读取管道数据 必须先关闭管道
	//要么会报错
	close(ch1)

	//for i := 0; i <= len(ch1); i++ {
	//	fmt.Println(<-ch1)
	//}

	//chan range时 没有key返回
	for v := range ch1 {
		fmt.Println(v)
	}

	// 这个位置会报错 必须写在18行
	//close(ch1)
}
