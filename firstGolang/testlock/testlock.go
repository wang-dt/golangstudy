package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁 单多协程多同一个资源进行操作的时候
var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var countt = 0
var m = make(map[int]int, 0)

func testloc(num int) {
	mutex.Lock()
	defer mutex.Unlock()

	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum

	fmt.Println(countt)
	time.Sleep(time.Millisecond)
	fmt.Printf("k=%d,v=%v", num, sum)

	wg.Done()

}
func main() {

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go testloc(i)
	}

	wg.Wait()
}
