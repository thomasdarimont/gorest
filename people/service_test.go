package people

import "testing"

func TestCreateAndFindPerson(t *testing.T) {

	ps := CreatePersonService()

	p := Person{Firstname: "Ralf", Lastname: "Steinbach", Address: &Address{City: "Saarbrücken", State: "Germany"}}
	newPerson := ps.CreatePerson(p)

	if newPerson.Firstname != p.Firstname || newPerson.Lastname != p.Lastname || newPerson.ID == "" {
		t.Fatal("new person is wrong")
	}
}

func TestFindByID(t *testing.T) {

	ps := CreatePersonService()

	person, e := ps.FindByID("1")
	if e != nil {
		t.Fatal("Found nothing")
	}

	if person.Firstname != "Thomas" || person.ID != "1" {
		t.Fatal("Should have found tom")
	}
}