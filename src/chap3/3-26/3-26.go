package main

import (
	"fmt"
	"time"
)

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>\n", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main() {
	go hello("hello", 50) // goroutine を用いて、直下のhello()とは別スレッドで実行する
	hello("bye!", 100)
}
