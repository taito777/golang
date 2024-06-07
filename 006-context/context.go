package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("長い処理が終わりました")
	case <-ctx.Done():
		fmt.Println("処理がキャンセルされました")
	}
}

func main() {
	// 3秒後にタイムアウトするcontextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go longRunningTask(ctx)

	// 4秒待ってからメインプログラムを終了
	time.Sleep(4 * time.Second)
	fmt.Println("メインプログラムが終了しました")
}
