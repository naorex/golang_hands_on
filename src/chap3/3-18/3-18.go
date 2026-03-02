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
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// 下記のメソッドを持つため、インターフェースを実装しているとみなされ、Data型として扱う事が可能となる
type Yourdata struct {
	Name string
	Mail string
	Age int
}

func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

func (md *Yourdata) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", md.Name, md.Age)
	fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {

	// 構造体を格納する配列を用意
	ob := [2]Data{}

	// Mydata構造体を設定
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})

	// Yourdata構造体を設定
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@mume.mo",
		"age": "34",
	})

	// 同じメソッドを用いて中身の異なるメソッドが動作（ポリモーフィズム）
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
