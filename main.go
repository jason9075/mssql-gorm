package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var server = "localhost"
var port = 1433
var user = "sa"
var password = "zaq12345"
var database = "astraDB"

type Member struct {
	gorm.Model
	ID   uint
	Name string
	Age  uint
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := gorm.Open("mssql", connString)
	if err != nil {
		fmt.Println(err)
	}

	// 防止 gorm 產生 SQL 時，自動在 table 加上 s
	// https://gorm.io/docs/conventions.html#Pluralized-Table-Name
	db.SingularTable(true)

	db.LogMode(true)

	var name string
	var id, age uint

	//  --- Create ---

	// db.Exec("INSERT INTO Member (ID, Name, Age) VALUES (?, ?, ?)", 2, "Jason Kuan", 32)

	// --- Read ---

	row := db.Raw("SELECT ID, Name, Age FROM Member WHERE ID = ?", 2).Row()
	row.Scan(&id, &name, &age)
	fmt.Printf("ID: %d, Name: %s, Age %d \n", id, name, age)

	// --- Update ---

	db.Exec("UPDATE Member SET Age = ? WHERE ID = ?", 40, 2)
	row = db.Raw("SELECT ID, Name, Age FROM Member WHERE ID = ?", 2).Row()
	row.Scan(&id, &name, &age)
	fmt.Printf("ID: %d, Name: %s, Age %d \n", id, name, age)

	// --- Delete ---

	db.Exec("DELETE Member WHERE ID = ?", 2)
}
