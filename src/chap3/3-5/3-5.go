package main

import "fmt"

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar) // [10 20 30]
	initial(&ar)
	fmt.Println(ar) // [0 0 0]
}

func initial(ar *[]int) {
	// ポインタをスライスで操作する関数
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0 // ar のポインタにあるスライスの [i] を 0 に置換
	}
}
