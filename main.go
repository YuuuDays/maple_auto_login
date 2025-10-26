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
		ch <- 2
		ch <- 3
		close(ch)
	}()

	// range　chはcloseされたら終わる。
	// closeは送り手が送るのがマナー
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("owari ★")
}
