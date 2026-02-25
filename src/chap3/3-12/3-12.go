package main

import (
	"fmt"
	"hello"
	"strconv"
)

// intp という型を int をベースに定義（int のエイリアスを設定し、int型の特徴を活かしつつメソッドを独自に定義したりできる）
type intp int

// intp にメソッドを追加
func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n/2); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// intp にメソッドを追加
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

// 整数を入力すると、素数かどうかチェックし、素因数分解した内容を出力する。その後数字を2倍して1を足して再度チェック
func main() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x *= 2
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
