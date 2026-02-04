package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("type a number")
	ar := strings.Split(x, " ") // 半角スペースで分割
	t := 0
	for i := 0; i < len(ar); i++ {
		n, err := strconv.Atoi(ar[i])
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
