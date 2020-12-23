package main

import (
	"fmt"
	"time"
)

func setVTO1(v *int) {
	*v = 1
}

func setVTo2(v *int) {
	*v = 2
}

func main() {
	v := new(int)
	go setVTO1(v)
	go setVTo2(v)
	fmt.Println(*v)
	go func(name string) {
		fmt.Println("hello " + name)
	}("dury")
	time.Sleep(time.Second)
}
