package orientrest

import "testing"

func TestQuery(t *testing.T) {
	db := openTestDb(t, "")
	r, err := db.Command(NewQuerySQL("select * from V"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestQuery: %+v", r.Result)
	db.Close()
}

func TestPeaEatersQuery(t *testing.T) {
	db := openTestDb(t, "")
	r, err := db.Command(NewQuerySQL("select expand(in(Eat)) from Food where name = 'pea'"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestPeaEatersQuery: %+v", r)
	db.Close()
}

func TestAnimalFoodsQuery(t *testing.T) {
	db := openTestDb(t, "")
	animal_foods, err := db.Command(NewQuerySQL("select expand(out(Eat)) from Animal"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestAnimalFoodsQuery: animal_foods: %+v", animal_foods)
	animal, err := db.Command(NewQuerySQL("select name from (select expand(in('Eat')) from Food where name = 'pea')"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestAnimalFoodsQuery: animal: %+v", animal)
	db.Close()
}
