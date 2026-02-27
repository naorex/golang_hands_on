package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
	Check() // インタフェースにもメソッドを追加
}

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

// 新たにメソッドを追加
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]\n", md.Name)
}

func main() {
	var ob Data = new(Mydata) // Mydata を Data型として扱う
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.Check()
}
