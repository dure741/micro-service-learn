package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	//循环向通道发送数据
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 0)
	go send(ch2, 10)
	//主goroutine休眠1s，保证调度成功
	time.Sleep(time.Second)
	for {
		select {
		case val := <-ch1: //从ch1读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case val := <-ch2: //从ch2读取数据
			fmt.Printf("get value %d from ch2\n", val)
		case <-time.After(2 * time.Second): //超时设置
			fmt.Println("Time out!")
			return
		}
	}
}
