package main

import "fmt"

/**
 * 管道写入和输出
 */
func main() {

	ch := make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 3
	//ch <- 3 当管道满了之后 写的时候会报错 deadlock-- 管道阻塞

	a := <-ch
	fmt.Println(a)
	<-ch // 取出一个值 不赋值 扔了

	b := <-ch
	fmt.Println(b)

	c := <-ch
	fmt.Println(c)

	//当管道中没有数据之后 会发生报错 deadlock-- 管道阻塞
	//d := <-ch
	//fmt.Println(d)
}
