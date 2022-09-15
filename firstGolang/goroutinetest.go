package main

import (
	"fmt"
	"sync"
)

func putintchan(intchan chan int) {

	defer close(intchan)
	for i := 2; i < 100; i++ {
		intchan <- i
	}
	wg.Done()
}

func primechans(intchan chan int, primechan chan int, exitchan chan bool) {

	for v := range intchan {
		var flag = true
		_ = flag
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primechan <- v
		}
	}
	exitchan <- true
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {

	//intchan 创建一个管道 放1-120000个数字
	intchan := make(chan int, 1000)

	wg.Add(1)
	go putintchan(intchan)

	//primechan 接收素数
	primechan := make(chan int, 50000)
	//标识什么时候关闭primechans
	exitchan := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go primechans(intchan, primechan, exitchan)
	}

	wg.Wait()
	/**等待问题
	 **/

	//判断exitchan是否存满
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			<-exitchan
		}

		//关闭primechan
		close(intchan)
		close(exitchan)

		wg.Done()
	}()

	fmt.Println("执行完毕--")

}
