package orientrest
/*
import "testing"

func TestInterruptCmd(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.CmdInterrupt("select * from V"); err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestGetAllCmd(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.CmdGetAll("People")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
	db.Close()
}

func TestCommand(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	//Animal
	_, err = db.Command("delete vertex Animal")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("drop class Animal")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("create class Animal extends V")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("insert into Animal set name = 'rat', specie = 'rodent'")
	if err != nil {
		t.Fatal(err)
	}
	//Food
	_, err = db.Command("delete vertex Food")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("drop class Food")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("create class Food extends V")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("insert into Food set name = 'pea', color = 'green'")
	if err != nil {
		t.Fatal(err)
	}
	//Eat
	_, err = db.Command("delete edge Eat")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("drop class Eat")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("create class Eat extends E")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Command("create edge Eat from (select from Animal where name = 'rat') to (select from Food where name = 'pea')")
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}
*/
