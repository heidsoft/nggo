package main

import (
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
)

type User struct {
	Id string `json:"uid"`
	Email string `json:"email"`
	Interests []string `json:"interests"`
}

func Create()  {

}

func Delete() {

}

func Update() {

}

func Query() {

}
func main() {
	cluster, _ := gocb.Connect("couchbase://127.0.0.1")
	bucket, _ := cluster.OpenBucket("default", "")

	bucket.Upsert("u:kingarthur",
		User{
			Id: "kingarthur",
			Email: "kingarthur@couchbase.com",
			Interests: []string{"Holy Grail", "African Swallows"},
		}, 0)

	// Get the value back
	var inUser User
	bucket.Get("u:kingarthur", &inUser)
	fmt.Printf("User: %v\n", inUser)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM default WHERE $1 IN interests")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"African Swallows"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}
}