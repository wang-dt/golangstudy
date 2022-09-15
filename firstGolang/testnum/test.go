package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 *wangdt
 */
var wg = sync.WaitGroup{}

func test(n int) {
	star := time.Now().Unix()
	for num := (n-1)*30 + 1; num < n*30; num++ {
		flss := true
		_ = flss
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				flss = false
				break
			}
			if flss {
				fmt.Println(num, "是素数")
			}
		}
	}
	end := time.Now().Unix()

	fmt.Println(end - star)

	defer wg.Done()

}
func main() {
	for i := 1; i <= 4; i++ {
		go wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
