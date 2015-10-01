package orientrest

import "testing"

func TestSumaFunction(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Command(`create function sumaFunc "return a + b" PARAMATERS [a,b]`)
	if err != nil {
		t.Fatal(err)
	}

	r, err := db.Function("sumaFunc", 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)

	db.Close()
}
