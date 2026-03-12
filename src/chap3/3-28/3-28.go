package main

import "fmt"

func total(n int, c chan int) { // チャンネル型を引数に持つ関数
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t // 引数で渡されたチャンネルに結果を渡す
}

func main() {
	c := make(chan int)
	go total(100, c)
	fmt.Println("total: ", <-c) // 引数にチャンネル型を持つため、上記の goroutine の結果が終わるまで出力されない
}
