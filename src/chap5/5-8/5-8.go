package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er) // エラー時は強制終了
	}
	defer re.Body.Close() // defer を付けると、関数が終わる直前に実行

	s, er := io.ReadAll(re.Body) // サーバーから返ってきたデータ（ストリーム形式）を、最後まで一気に読み込んでバイト配列（[]byte）に変換
	if er != nil {
		panic(er)
	}

	fmt.Println(string(s)) // 読み込んだバイトデータを文字列に変換して表示
}
