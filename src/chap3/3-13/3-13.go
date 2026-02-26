package main

import (
	"fmt"
	"hello"
	"strconv"
)

type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n/2); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x % n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

// もっとも大きな素数に値を変更する関数
func (num *intp) doPrime() { // intp のポインタが num に渡される
	pf := num.PrimeFactor() // ポインタからのメソッド呼び出しは (*num).PrimeFactor() のような書き方にする必要がない
	*num = intp(pf[len(pf) - 1])
}

// 整数を入力すると素数判定と素因数を表示する。そして一番大きな素数に変更し、1加算する
func main() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x.doPrime()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
