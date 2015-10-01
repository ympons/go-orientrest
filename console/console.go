package main

import (
	"github.com/ympons/go-orientrest"
	"log"
)

func main() {
	a, b := orientrest.STORAGE_TYPE_LOCAL, orientrest.DB_TYPE_GRAPH
	db, _ := orientrest.OrientDB("")
	client, _ := db.Connect(gorientrest.Options{
		DbUser: "admin",
		DbPass: "admin",
	})
	client.DbCreate("mibd", orientrest.STORAGE_TYPE_LOCAL, orientrest.DB_TYPE_GRAPH)
	log.Printf("%s %s", a, b)
	client.Close()
}
