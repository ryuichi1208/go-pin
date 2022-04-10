package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/cooking")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	result := db.QueryRow("select * from items")
	var name string
	var price int

	result.Scan(&name, &price)
	fmt.Println(name)

	// a := []string{"test1", "test2", "test3", "test4", "test5"}
	// for _, v := range a {
	// 	res, err := db.Exec(`insert into items(name, price) values(?,?)`, v, 100)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(res)
	// }
	// res, err := db.Exec(`insert into items(name, price) values("aaa",200)`)
	// fmt.Println(res, err)

	query := `
	INSERT INTO items(
		name,
		price
	) VALUES
	`

	n := []string{
		"aaa",
		"bbb",
		"ccc1",
		"ccc2",
		"ccc3",
		"ccc5",
		"ccc6",
	}
	var values []string
	for _, a := range n {
		val := fmt.Sprintf("('%s', %d)", a, 111)
		values = append(values, val)
	}
	query += strings.Join(values, ",")
	fmt.Println(query)
	_, err = db.Exec(query)
	if err != nil {
		panic(err)

	}
}
