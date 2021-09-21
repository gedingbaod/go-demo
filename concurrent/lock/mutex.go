package main

import (
	"fmt"
	"sync"
)

var x1 int64
var wg1 sync.WaitGroup
var lock sync.Mutex

func add1() {
	for i := 0; i < 5000; i++ {
		// 互斥锁
		lock.Lock() // 加锁
		x1 = x1 + 1
		lock.Unlock() // 解锁
	}
	wg1.Done()
}
func main() {
	wg1.Add(4)
	go add1()
	go add1()
	go add1()
	go add1()
	wg1.Wait()
	fmt.Println(x1)
}
