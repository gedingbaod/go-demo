package main

import (
	"fmt"
	"time"
)

func main() {
	// 4.停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}
	for {
	}
}
