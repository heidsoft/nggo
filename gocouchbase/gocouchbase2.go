package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/gocb"
)

var bucket *gocb.Bucket

type Person struct {
	FirstName        string            `json:"firstname,omitempty"`
	LastName         string            `json:"lastname,omitempty"`
	SocialNetworking *SocialNetworking `json:"socialNetworking,omitempty"`
}

type SocialNetworking struct {
	Twitter string `json:"twitter,omitempty"`
	Website string `json:"website,omitempty"`
}

func updateDocument(documentId string) {
	fmt.Println("Update document by id...")
	var person Person
	_, error := bucket.Get(documentId, &person)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	jsonPerson, _ := json.Marshal(&person)
	fmt.Println(string(jsonPerson))
}


func getDocument(documentId string) {
	fmt.Println("Getting the full document by id...")
	var person Person
	_, error := bucket.Get(documentId, &person)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	jsonPerson, _ := json.Marshal(&person)
	fmt.Println(string(jsonPerson))
}

func createDocument(documentId string, person *Person) {
	fmt.Println("Upserting a full document...")
	_, error := bucket.Upsert(documentId, person, 0)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	getDocument(documentId)
	getSubDocument(documentId)
}

func getSubDocument(documentId string) {
	fmt.Println("Getting part of a document by id...")
	fragment, error := bucket.LookupIn(documentId).Get("socialNetworking").Execute()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	var socialNetworking SocialNetworking
	fragment.Content("socialNetworking", &socialNetworking)
	jsonSocialNetworking, _ := json.Marshal(&socialNetworking)
	fmt.Println(string(jsonSocialNetworking))
	upsertSubDocument(documentId, "thepolyglotdeveloper.com")
}

func upsertSubDocument(documentId string, website string) {
	fmt.Println("Upserting part of a document...")
	_, error := bucket.MutateIn(documentId, 0, 0).Upsert("socialNetworking.website", website, true).Execute()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	getDocument(documentId)
}

type Datacenter struct{
	Name        string            `json:"name"`
	Description string            `json:"description"`
}
type DcMeta struct {
	Class       string            `json:"class"`
	Description string            `json:"description"`
	Data 	    []*Datacenter     `json:"data"`
}

func main() {
	fmt.Println("Starting the app...")
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")
	//person := Person{FirstName: "liu", LastName: "bin", SocialNetworking: &SocialNetworking{Twitter: "heidsoft"}}
	//createDocument("liubin", &person)

	//cas,err := bucket.Upsert("jake", person, 0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//person.FirstName = "wen"
	//_,err = bucket.Replace("jake",person,cas,0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}

	//bucket.Remove("datacenter_meta",0)
	//dcs := []*Datacenter{} //空切片
	//
	//dc1 := new(Datacenter)
	//dc1.Name = "上海"
	//dc1.Description = "上海数据中心"
	//dcs = append(dcs,dc1)
	//
	//dc2 := new(Datacenter)
	//dc2.Name = "北京"
	//dc2.Description = "北京数据中心"
	//dcs = append(dcs,dc2)
	//
	//dc3 := new(Datacenter)
	//dc3.Name = "新加坡"
	//dc3.Description = "新加坡数据中心"
	//dcs = append(dcs,dc3)
	//
	//dcMeta := new(DcMeta)
	//dcMeta.Description = "数据中心元数据"
	//dcMeta.Class = "datacenter"
	//dcMeta.Data = dcs
	//
	//fmt.Println("插入数据中心元数据....")
	//cas,err := bucket.Upsert("datacenter_meta", dcMeta, 0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//
	//dc4 := new(Datacenter)
	//dc4.Name = "纽约"
	//dc4.Description = "纽约数据中心"
	//dcs = append(dcs,dc4)
	//dcMeta.Data = dcs
	//
	//fmt.Println("更新数据中心元数据....")
	//_,err = bucket.Replace("datacenter_meta",dcMeta,cas,0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}

	var myDcMeta DcMeta
	_, error := bucket.Get("datacenter_meta", &myDcMeta)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	jsonMyDcMeta, _ := json.Marshal(&myDcMeta)
	fmt.Println(string(jsonMyDcMeta))
}