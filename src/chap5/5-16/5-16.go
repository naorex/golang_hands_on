package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// 必要性で言えば以下の構造体定義をしなくても問題無いが、普通 DB から取り出したレコードを扱う際には構造体で受けた方が、後々の拡張が容易
type Mydata struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\"" + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

func main() {
	con, er := sql.Open("sqlite3", "../data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	q := "SELECT * FROM mydata"
	rs, er := con.Query(q) // 戻り値は Rows という構造体
	if er != nil {
		panic(er)
	}
	for rs.Next() { // Rows からレコードを取り出す
		var md Mydata
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println((md.Str()))
	}
}
