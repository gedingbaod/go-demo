package main

import (
	"fmt"
	"time"
)

func main() {

	// 2.验证timer只能响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}

	for {
	}
}
