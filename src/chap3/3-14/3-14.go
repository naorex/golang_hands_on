package main

import "fmt"

// 構造体のインターフェースを定義
// 以降、Initial() と PrintData() を持つ構造体は自動的に Data インタフェースを実装しているとみなされる
type Data interface {
	Initial(name string, data []int) // インターフェースのメソッドを定義
	PrintData() // 同上
}

// 下記のメソッドを持つため、インターフェースを実装しているとみなされ、Data 型として扱う事が可能となる
type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func main() {
	var ob Mydata = Mydata{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
}
