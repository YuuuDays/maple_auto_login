package main

import (
	"fmt"
)

func main() {
	/*
	 make(chan 型)←チャネルを作成。
	 これはゴルーチン同士で値を受け渡すために使用される
	*/
	ch := make(chan int)

	go func() {
		ch <- 1
		fmt.Println("A: send done")
	}()

	x := <-ch
	fmt.Println("B: received", x)
}
