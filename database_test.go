package orientrest

import (
	"os"
	"log"
	"testing"
)

type Fatalistic interface {
	Fatal(args ...interface{})
}

var defaultTo = func(k string, v string) string {
	if val := os.Getenv(k); val != "" {
		return val
	}
	os.Setenv(k, v)
	return v
}

func openTestDb(t Fatalistic, server string) *Database {
	dbname := defaultTo("ODATABASE", "testdb")
	user := defaultTo("OUSER", "root")
	pass := defaultTo("OPASS", "root")

	client, err := New("")
	if err != nil {
		t.Fatalf(err)
	}

	db, err := client.Open(dbname, user, pass)
	if err != nil {
		t.Fatalf(err)
	}
	return db
}

func openTestAdmin(t Fatalistic, server string) *Admin {
	user := defaultTo("OUSER", "root")
	pass := defaultTo("OPASS", "root")

	client, err := New("")
	if err != nil {
		t.Fatalf(err)
	}

	admin, err := client.Auth(user, pass)
	if err != nil {
		t.Fatalf(err)
	}

	return admin
}

func openTestDb(t Fatalistic) *ODatabase {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	return db
}

/*
func TestOpenDb(t *testing.T) {
	urlFunc := func(url string) {
		db, err := openTestDbinfo(url)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
	}
	urlFunc("http://127.0.0.1:2480")
}

func TestCreateDb(t *testing.T) {
	db, err := openTestDbinfo("", false)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.DbCreate("testdb", STORAGE_TYPE_PLOCAL, DB_TYPE_GRAPH); err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestDropDb(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.DbDrop("testdb"); err != nil {
		t.Fatal(err)
	}
}

func TestInfoDb(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}

	r, err := db.DbInfo("testdb")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	db.Close()
}

func TestListDbs(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.DbList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	db.Close()
}

func TestAvailableLangs(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.DbAvailableLangs("testdb")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	db.Close()
}
*/
