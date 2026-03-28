package main

import (
	"database/sql"
	"fmt"
	"hello"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\"" + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}
var qry string = "SELECT * FROM mydata WHERE id = ?" // プレースホルダ

func main() {
	con, er := sql.Open("sqlite3", "../data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	// 以下、継続実行して、id 入力あるたびにクエリでDBから値を取得・表示する
	for true {
		s := hello.Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs, er := con.Query(qry, n) // 入力値を引数に入れてプレースホルダに渡す
		if er != nil {
			panic(er)
		}
		for rs.Next() {
			var md Mydata
			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
			if er != nil {
				panic(er)
			}
			fmt.Println(md.Str())
		}
	}
	fmt.Println("***end***")
}
