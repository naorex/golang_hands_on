package main

import (
	"fmt"
	"strconv"
	"time"
)

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>\n", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main() {

	msg := "start"

	// 共通メモリで使用されるメソッド（以降の2つのメソッド内で呼び出し）
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	// goroutine で実行するスレッド用
	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	// メインスレッド用
	main := func(n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}

	go hello(60) //=> h9 が出力される前にメインスレッドが完了するため、h8 までしか出力されない
	main(100)
}
