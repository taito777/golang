package main

import (
	"flag"
	alias "fmt" // パッケージにエイリアスを設定することが可能
)

type A struct {
	ID   int
	Name string
}

func (r *A) ChangeName(name string) {
	r.Name = name
	alias.Println(*r)
}

func main() {
	b := A{33, ""}       // A structを初期化した値を変数bに入力
	c := &b              // 変数cに変数bのメモリアドレスを入力
	c.ChangeName("tabi") // ChangeName関数の使用

	for i := 0; i <= 5; i++ {
		flag.Parse()           // flag.Parse(): 引数をコード内で取得
		a := A{i, flag.Arg(i)} // flag.Arg(i): i番目の引数

		if a.ID <= 3 && a.Name != "" {
			alias.Println(a)
		} else {
			alias.Println("error")
		}

	}
}
