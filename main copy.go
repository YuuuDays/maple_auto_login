// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	/*
// 	 make(chan 型)←チャネルを作成。
// 	 これはゴルーチン同士で値を受け渡すために使用される

// 	 これを使って停止用チャネルというものを作る
// 	 更にこれを標準ライブラリー化したものがcontext(ctx)
// 	 以下がgoで慣習的に書く例
// 	 ctx := context.Background()
// 	 ctx, cancel := context.WithCancel(ctx)
// 	*/

// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		cancel() // ← 停止合図（close(stop) と同じ役割）
// 	}()

// 	/* range　chはcloseされたら終わる。
// 	 //closeは送り手が送るのがマナー
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// 	*/
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("停止検知")
// 			return
// 		default:
// 			fmt.Println("処理中")
// 			time.Sleep(300 * time.Millisecond)
// 		}
// 	}
// }
