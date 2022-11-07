package main

import (
	// "database/sql"
	// _ "github.com/denisenkom/go-mssqldb" // It is driver
	"errors"
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

	//  --- Create ---

	// member := Member{ID: 1, Name: "Jason Kuan", Age: 30}
	// result := db.Create(&member)
	// fmt.Println(result.Value)

	// --- Read ---

	var member Member
	// result := db.Find(&member, "Age = ?", 30)
	result := db.Model(&Member{}).First(&member)
	fmt.Printf("is not found: %t \n", errors.Is(result.Error, gorm.ErrRecordNotFound))
	fmt.Println(result.Value)
}
