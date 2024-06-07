package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	messages := []string{"こんにちは", "お元気ですか？", "さようなら"}
	for _, msg := range messages {
		fmt.Println("送信:", msg)
		ch <- msg                   // メッセージをチャンネルに送信
		time.Sleep(1 * time.Second) // 1秒待つ
	}
	close(ch) // 送信が終わったらチャンネルを閉じる
}

func receiver(ch chan string) {
	for msg := range ch { // チャンネルからメッセージを受信
		fmt.Println("受信:", msg)
	}
}

func main() {
	ch := make(chan string) // チャンネルを作成

	go sender(ch)   // 送信ゴルーチンを開始
	go receiver(ch) // 受信ゴルーチンを開始

	time.Sleep(5 * time.Second) // 全てのメッセージが処理されるのを待つ
	fmt.Println("メインプログラムが終了しました")
}
