package main

import "fmt"

func main() {
	// 没有缓冲区，所以出错
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
