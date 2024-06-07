package main

import (
	"fmt"
	"time"
)

// goroutine:複数の処理を並行して進めるための機能

func bakeCookies() {
	fmt.Println("クッキーを焼き始めました")
	time.Sleep(2 * time.Second) // クッキーを焼くのに2秒かかる
	fmt.Println("クッキーが焼き上がりました")
}

func bakeCake() {
	fmt.Println("ケーキを焼き始めました")
	time.Sleep(3 * time.Second) // ケーキを焼くのに3秒かかる
	fmt.Println("ケーキが焼き上がりました")
}

func meltChocolate() {
	fmt.Println("チョコレートを溶かし始めました")
	time.Sleep(1 * time.Second) // チョコレートを溶かすのに1秒かかる
	fmt.Println("チョコレートが溶けました")
}

func main() {
	go bakeCookies()   // クッキーを焼くGoルーチンを開始
	go bakeCake()      // ケーキを焼くGoルーチンを開始
	go meltChocolate() // チョコレートを溶かすGoルーチンを開始

	time.Sleep(4 * time.Second) // 全ての作業が終わるのを待つ
	fmt.Println("すべてのお菓子が完成しました！")
}
