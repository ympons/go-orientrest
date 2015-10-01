package orientrest

import (
	"os"
	"testing"
)

type Fatalistic interface {
	Fatal(args ...interface{})
}

func openTestDbinfo(info string, c ...bool) (*ODatabase, error) {
	defaultTo := func(k string, v string) string {
		if val := os.Getenv(k); val != "" {
			return val
		}
		os.Setenv(k, v)
		return v
	}
	dbname := defaultTo("ODATABASE", "testdb")
	user := defaultTo("OUSER", "admin")
	pass := defaultTo("OPASS", "admin")
	conn := true

	if c != nil {
		conn = false
	}

	db, err := OrientDB(info, Options{
		DbName: dbname,
		DbUser: user,
		DbPass: pass,
		Conn: conn,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openTestDb(t Fatalistic) *ODatabase {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	return db
}

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
