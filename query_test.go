package orientrest
/*
import "testing"

func TestQuery(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.Query("select * from V")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestQuery: %+v", r.Results)
	db.Close()
}

func TestPeaEatersQuery(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.Query("select expand(in(Eat)) from Food where name = 'pea'")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestPeaEatersQuery: %+v", r)
	db.Close()
}

func TestAnimalFoodsQuery(t *testing.T) {
	db, err := openTestDbinfo("")
	if err != nil {
		t.Fatal(err)
	}
	animal_foods, err := db.Query("select expand(out(Eat)) from Animal")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestAnimalFoodsQuery: animal_foods: %+v", animal_foods)
	animal, err := db.Query("select name from (select expand(in('Eat')) from Food where name = 'pea')")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestAnimalFoodsQuery: animal: %+v", animal)
	db.Close()
}
*/
