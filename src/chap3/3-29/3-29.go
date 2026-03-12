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
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z) // goroutine で速く計算が終わったものが出力される（例：55 500500 5050）
}
