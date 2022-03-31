package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql" //导入mysql驱动包"

type User struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root:howie@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	var user []User
	rows, err := db.Query("select * from  user")
	for rows.Next() {
		var Id int
		var Name string
		if err := rows.Scan(&Id, &Name); err != nil {
			continue
		}
		user = append(user, User{Id: Id, Name: Name})

	}
	fmt.Println(user)
	//[{1 111111} {2 22222} {3 33333} {4 44444}]
}
