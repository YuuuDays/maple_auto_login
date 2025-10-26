package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 make(chan 型)←チャネルを作成。
	 これはゴルーチン同士で値を受け渡すために使用される
	*/
	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(200 * time.Microsecond)
		a <- "Aのmessage"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		b <- "Bのmessage"
	}()

	/* range　chはcloseされたら終わる。
	 //closeは送り手が送るのがマナー
	for v := range ch {
		fmt.Println(v)
	}
	*/

	select {
	case x := <-a:
		fmt.Println(x)
	case y := <-b:
		fmt.Println(y)
	default:
		fmt.Println("mada")
	}
}
