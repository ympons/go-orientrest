package main

import (
	"github.com/ympons/go-orientrest"
	"log"
)

func main() {
	db, _ := orientrest.OrientDB("", orientrest.Options{
		DbUser: "root",
		DbPass: "root",
	})
	if err := db.DbCreate("testdb", orientrest.STORAGE_TYPE_PLOCAL, orientrest.DB_TYPE_GRAPH); err != nil {
		log.Printf("ERROR: Creating db... %+v", err)
	}
	client, err := db.Connect(orientrest.Options{
		DbName: "testdb",
		DbUser: "admin",
		DbPass: "admin",
	})
	if err != nil {
		log.Printf("ERROR: Connecting db... %+v", err)
	}
	if info, err := client.DbInfo(client.Name); err != nil {
		log.Fatalf("ERROR: Getting info... %+v", err)
	} else {
		log.Printf("%v", info)
	}
	client.Close()
}
