package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("type a number")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {  // index と 要素を取り出す記法。"_" は未使用でもエラーとならない特殊な変数名
		n, err := strconv.Atoi(v)
		if err != nil {
			goto err
		}
		t += n
		}
	fmt.Println("total:", t)
	return

	err:
	fmt.Println("ERROR!!")
}
