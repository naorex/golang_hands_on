package main

import (
	"fmt"
	"time"
)

func total(c chan int) { // チャンネル型を引数に持つ関数
	n := <-c // チャンネルに結果を渡す
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
}

func main() {
	c := make(chan int)
	go total(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
}
