package main

import (
	"github.com/ympons/go-orientrest"
	"log"
)

func main() {
	client, err := orientrest.New("")
	if err != nil {
		log.Fatal(err)
	}

	admin, err := client.Auth("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	_, err = admin.DbCreate("testdb", orientrest.DB_TYPE_GRAPH, orientrest.STORAGE_TYPE_PLOCAL)
	if err != nil {
		log.Fatal(err)
	}

	admin.Close()
}
