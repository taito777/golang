package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter:リクエストに対してのレスポンスを格納
	// http.Request:サーバへのhttpリクエストを格納

	fmt.Fprintln(w, "Hello, World!")
	// Fprintln(書き込み先, 書き込み内容)
}

func main() {
	http.HandleFunc("/hello/", Handler)
	// http.HandleFunc:指定したパスへのリクエストを処理するメソッドを指定
	// http.HundleFunc("パス", "メソッド")

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		// http.ListenAndServe(addr string, handler Handler) error:指定したアドレスでリクエストを待つようになる
		// errorが発生した場合、エラー内容を戻り値として返す

		fmt.Println("Error starting server:", err)
	}
}
