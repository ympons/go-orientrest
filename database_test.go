package orientrest

import (
	"os"
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
		t.Fatal(err)
	}

	db, err := client.Open(dbname, user, pass)
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func openTestAdmin(t Fatalistic, server string) *Admin {
	user := defaultTo("OUSER", "root")
	pass := defaultTo("OPASS", "root")

	client, err := New("")
	if err != nil {
		t.Fatal(err)
	}

	admin, err := client.Auth(user, pass)
	if err != nil {
		t.Fatal(err)
	}

	return admin
}

func TestAuthAdmin(t *testing.T) {
	admin := openTestAdmin(t, "")
	admin.Close()
}

func TestOpenDb(t *testing.T) {
	db := openTestDb(t, "")
	db.Close()
}

func TestCreateDb(t *testing.T) {
	admin := openTestAdmin(t, "")
	d, err := admin.DbCreate("testdb", DatabaseType(DB_TYPE_GRAPH), StoreType(STORAGE_TYPE_PLOCAL))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", d)
	admin.Close()
}

func TestDropDb(t *testing.T) {
	admin := openTestAdmin(t, "")
	if err := admin.DbDrop("testdb"); err != nil {
		t.Fatal(err)
	}
	admin.Close()
}

func TestInfoDb(t *testing.T) {
	admin := openTestAdmin(t, "")
	r, err := admin.DbInfo("testdb")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	admin.Close()
}

func TestListDbs(t *testing.T) {
	admin := openTestAdmin(t, "")
	r, err := admin.DbList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	admin.Close()
}

func TestAvailableLangs(t *testing.T) {
	admin := openTestAdmin(t, "")
	r, err := admin.DbAvailableLangs("testdb")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	admin.Close()
}
