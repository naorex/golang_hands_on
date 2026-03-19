package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type SrData struct {
	msg string
	mux sync.Mutex // こちらの変数には値は設定しない（排他制御用のパラメータ）
}

func main() {
	sd := SrData{msg: "Start"}
	prmsg := func(nm string, n int) {
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	main := func(n int) {
		const nm string = "*main"
		sd.mux.Lock() // スレッドをロック
		for i := 0; i < 5; i++ {
			sd.msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
		sd.mux.Unlock() // スレッドをアンロック
	}

	hello := func(n int) {
		const nm string = "hello"
		sd.mux.Lock()
		for i := 0; i < 5; i++ {
			sd.msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock()
	}

	// 以下の2つのスレッドで、先に実行した方が終わるまで待つ動作となる
	go main(100)
	go hello(50)
	time.Sleep(5 * time.Second)
}
