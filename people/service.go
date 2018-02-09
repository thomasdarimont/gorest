package people

import (
	"errors"
	"github.com/satori/go.uuid"
)

var errNotFound = errors.New("Person not found")

var personService = CreatePersonService()

func CreatePersonService() PersonService {

	service := &defaultPersonService{}

	tom := Person{ID: "1", Firstname: "Thomas", Lastname: "Darimont", Address: &Address{City: "Saarbrücken", State: "Saarland"}}
	anne := Person{ID: "2", Firstname: "Anne", Lastname: "Laurentius", Address: &Address{City: "Saarbrücken", State: "Saarland"}}

	service.people = append(service.people, tom)
	service.people = append(service.people, anne)

	return service
}

type PersonService interface {

	CreatePerson(p Person) Person

	DeletePerson(id string) bool

	FindAll() []Person

	FindByID(id string) (Person, error)
}

type defaultPersonService struct {
	people []Person
}

func (ps *defaultPersonService) CreatePerson(p Person) Person {

	p.ID = uuid.NewV4().String()
	ps.people = append(ps.people, p)

	return p
}

func (ps *defaultPersonService) FindAll() []Person {
	return ps.people
}

func (ps *defaultPersonService) FindByID(id string) (Person, error) {

	for _, p := range ps.people {
		if p.ID == id {
			return p, nil
		}
	}

	return Person{}, errNotFound
}

func (ps *defaultPersonService) DeletePerson(id string) bool {

	for index, item := range ps.people {
		if item.ID == id {
			ps.people = append(ps.people[:index], ps.people[index+1:]...)

			return true
		}
	}

	return false
}