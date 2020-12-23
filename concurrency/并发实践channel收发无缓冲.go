package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	//使用for 循环从channel中读取数据
	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	//创建没有缓冲的channel
	ch := make(chan string)
	go printInput(ch)
	//从命令行中读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("End the game!")
		}
	}
	//程序最后关闭 channel
	defer close(ch)
}
