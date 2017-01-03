package main


import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:oneoaas@2A@tcp(10.0.2.47:3306)/cmdb_v15?charset=utf8&loc=Local")
	if err != nil{
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil{
		fmt.Println(err)
	}
}