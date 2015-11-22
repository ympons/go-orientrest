package orientrest

import "testing"

func TestCommand(t *testing.T) {
	db := openTestDb(t, "")

	//Animal
	_, err := db.Command(NewCommandSQL("delete vertex Animal"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("drop class Animal"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("create class Animal extends V"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("insert into Animal set name = 'rat', specie = 'rodent'"))
	if err != nil {
		t.Fatal(err)
	}

	//Food
	_, err = db.Command(NewCommandSQL("delete vertex Food"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("drop class Food"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("create class Food extends V"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("insert into Food set name = 'pea', color = 'green'"))
	if err != nil {
		t.Fatal(err)
	}

	//Eat
	_, err = db.Command(NewCommandSQL("delete edge Eat"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("drop class Eat"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("create class Eat extends E"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command(NewCommandSQL("create edge Eat from (select from Animal where name = 'rat') to (select from Food where name = 'pea')"))
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestInterruptCmd(t *testing.T) {
	db := openTestDb(t, "")
	if err := db.CmdInterrupt("select * from V"); err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestGetAllCmd(t *testing.T) {
	db := openTestDb(t, "")
	r, err := db.CmdGetAll("People")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	db.Close()
}
