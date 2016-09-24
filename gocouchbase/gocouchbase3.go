package main

import (
	"fmt"
	"github.com/couchbase/gocb"
)

func main() {
	var bucket *gocb.Bucket
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")

	myQuery := gocb.NewN1qlQuery("UPDATE default USE KEYS ? SET data = ARRAY c FOR c IN data WHEN c.id !=? END ")
	var myParams []interface{}
	myParams = append(myParams,"firewall")
	myParams = append(myParams,1471596117)
	rows,err := bucket.ExecuteN1qlQuery(myQuery,myParams)
	if err != nil{
		fmt.Printf(err.Error())
	}
	fmt.Printf("Results: %+v\n", rows)
}
