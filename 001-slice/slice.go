package main

import "fmt"

func main() {

	//スライス作成
	l := []string{"jave", "python", "go"}

	//インデックス番号とスライスの要素を表示
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	//スライスの要素だけを表示
	for _, v := range l {
		fmt.Println(v)
	}

	//"_"を使わずに要素を表示させようとすると
	for v := range l {
		fmt.Println(v)
	}
	//インデックス番号が表示される
	//"_"でインデックス番号を埋める必要がある

	//map作成
	m := map[string]int{"apple": 100, "banana": 200}

	//mapのキーと値を表示
	for k, v := range m {
		fmt.Println(k, v)
	}

	//mapのキーを表示
	for k := range m {
		fmt.Println(k)
	}

	//mapの値を表示
	for _, v := range m {
		fmt.Println(v)
	}

	// "_"はコード内でしようしなくてもエラーが発生しない特殊な変数
	// "_"以外の変数は、定義した場合コード内で必ず使用する必要がある。（使用しない場合はエラーが発生）
}
