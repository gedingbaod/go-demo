package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m1 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)
			value, _ := m1.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
