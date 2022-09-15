package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 *wangdt
 */
//读写互斥锁-- RWMutex
//允许多个读并行，但是写只能有一个

var rmmutex = sync.RWMutex{}
var wgs = sync.WaitGroup{}
var mutexa = sync.RWMutex{}

func read() {
	mutexa.RLock()
	defer mutexa.RUnlock()
	fmt.Println("------读操作")
	time.Sleep(time.Second * 2)
	wgs.Done()
}

func write() {
	mutexa.Lock()
	defer mutexa.Unlock()
	fmt.Println("写操作")
	time.Sleep(time.Second * 2)
	wgs.Done()
}
func main() {

	for i := 0; i < 10; i++ {
		wgs.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {

		wgs.Add(1)
		go write()
	}

	wgs.Wait()
}
