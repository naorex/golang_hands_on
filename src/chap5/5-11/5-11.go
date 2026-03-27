package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Tel  string `json:"tel"`
}

func main() {
	// 本来は http.Get で取得するはずのデータを文字列で代用
	s := []byte(`{
		"0": {"name": "Taro", "mail": "taro@yamada", "tel": "090-999-999"},
		"1": {"name": "Hanako", "mail": "hanako@flower", "tel": "080-888-888"},
		"2": {"name": "Sachiko", "mail": "sachi@happy", "tel": "070-777-777"}
	}`)

	// 構造体のマップとして受け取る
	var data map[string]Person
	er := json.Unmarshal(s, &data)
	if er != nil {
		fmt.Println("Unmarshal error:", er)
		return
	}

	for k, p := range data {
		fmt.Printf("ID:%s Name:%s Mail:%s Tel:%s\n", k, p.Name, p.Mail, p.Tel)
	}
}
