package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 make(chan 型)←チャネルを作成。
	 これはゴルーチン同士で値を受け渡すために使用される

	 これを使って停止用チャネルというものを作る
	 更にこれを標準ライブラリー化したものがcontext(ctx)
	 以下がgoで慣習的に書く例
	 ctx := context.Background()
	 ctx, cancel := context.WithCancel(ctx)
	*/

	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "仕事終わったよ"
	}()

	/* range　chはcloseされたら終わる。
	 //closeは送り手が送るのがマナー
	for v := range ch {
		fmt.Println(v)
	}
	*/

	select {
	case msg := <-ch:
		fmt.Println("受信:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("タイムアウト")
	}
}
