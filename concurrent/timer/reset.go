package main

import (
	"fmt"
	"time"
)

func main() {
	// 5.重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

	for {
	}
}
