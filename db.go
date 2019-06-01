package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type users struct {
	id     int
	name   string
	score  int
	groups string
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/msu-go?charset=utf8")
	//Select from database
	rows, _ := db.Query("select * from users")
	for rows.Next() {
		var get users
		err = rows.Scan(&get.id, &get.name, &get.score, &get.groups)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(get.id, " ", get.name, " ", get.score, " ", get.groups)
	}
	rows.Close()
	// Insert data into table
	result, error := db.Exec(
		"insert into users(`name`,`score`,`groups`) values (?,?,?)", "Jakhongir", "15", "942-15")
	lastId, err := result.LastInsertId()
	if error != nil {
		log.Fatal(error)
	}
	//Prepare statement
	fmt.Println(lastId)
	stmt, err := db.Prepare("update users set name = ?, score = ? , groups = ? where id = ?")
	result1, err := stmt.Exec("here we go", lastId, lastId, lastId)
	result1.RowsAffected()

}
