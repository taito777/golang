package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func goodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good, World!")
}

func main() {
	http.HandleFunc("/hello/", helloHandler) // HandleFunc:指定したパスへのリクエストを処理するメソッドを指定
	http.HandleFunc("/good/", goodHandler)   // http.HundleFunc("パス", "メソッド")
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
