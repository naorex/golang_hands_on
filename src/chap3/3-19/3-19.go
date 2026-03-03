package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 構造体のインターフェースを定義（Data型を自作）
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// 下記のメソッドを持つため、インターフェースを実装しているとみなされ、Data型として扱う事が可能となる
type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

func (md *Mydata) PrintData() {
	// nil レシーバでの表示に対応させる
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	} else {
		fmt.Println("**This is Nil value.**")
	}
}

func main() {
	var ob *Mydata
	ob.PrintData() // => nil の状態で PrintData() を呼び出す
	ob = &Mydata{} // Mydata というポインタを代入
	ob.SetValue(map[string]string{
		"name": "Jiro",
		"data": "123 456 789",
	})
	ob.PrintData() // => SetValue()で設定した値を出力
}
