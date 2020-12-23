package main

import (
	"fmt"
)

func test() {
	fmt.Println("I am work in a single goroutine")
}

func main() {
	go test()
	//time.Sleep(time.Second)//此时阻塞了main所在的goroutine，使得上一句goroutine能够成功执行完而不至于提前结束
}
