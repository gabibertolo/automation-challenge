package model

import (
	"encoding/json"
	"errors"
)

type pkgDependencies struct {
	countries []Country
	people    []Person
}

func NewPackage() PackageInterface {
	return &pkgDependencies{
		people:    initializePeople(),
		countries: initializeCountries(),
	}
}

type PackageInterface interface {
	GetPeople() []Person
	GetCountries() []Country
	CreatePerson(per Person) (*Person, error)
	DeletePerson(ID int) error
	GetWorldViewModel() (*string, error)
}

func (p *pkgDependencies) GetWorldViewModel() (*string, error) {
	model := generateDiagramModel(p.countries, p.people)
	bytes, err := json.Marshal(model)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	strJson := string(bytes)
	return &strJson, nil
}

func (p *pkgDependencies) GetCountries() []Country {
	return p.countries
}

func (p *pkgDependencies) GetPeople() []Person {
	return p.people
}

func (p *pkgDependencies) CreatePerson(personRequest Person) (*Person, error) {
	for _, person := range p.people {
		if person.FirstName == personRequest.FirstName && person.LastName == personRequest.LastName {
			return nil, errors.New("person with same first name and last name already exists")
		}
	}

	if personRequest.Age != 0 && personRequest.Age < 18 {
		return nil, errors.New("user must be at least 18 years old")
	}
	lastPerson := p.people[len(p.people)-1]
	personRequest.ID = lastPerson.ID + 1
	p.people = append(p.people, personRequest)
	return &personRequest, nil
}

func (p *pkgDependencies) DeletePerson(ID int) error {

	var updatedPeople []Person
	for _, person := range p.people {
		if person.ID != ID {
			updatedPeople = append(updatedPeople, person)
		}
	}
	p.people = updatedPeople
	return nil
}

func generateDiagramModel(countries []Country, persons []Person) interface{} {
	gojsModel := make(map[string]interface{})
	var nodesModel []interface{}

	for _, country := range countries {
		countryNode := make(map[string]interface{})
		countryNode["key"] = country.ID
		countryNode["text"] = country.Name
		countryNode["isGroup"] = true
		countryNode["category"] = "country-node-template"

		nodesModel = append(nodesModel, countryNode)
	}

	for _, p := range persons {
		pNode := make(map[string]interface{})
		pNode["key"] = p.ID
		pNode["text"] = p.FirstName + " " + p.LastName
		pNode["isGroup"] = false
		pNode["group"] = p.CountryID
		pNode["category"] = "person-node-template"

		nodesModel = append(nodesModel, pNode)
	}

	gojsModel["class"] = "go.GraphLinksModel"
	gojsModel["nodeDataArray"] = nodesModel
	gojsModel["linkDataArray"] = []interface{}{}

	return gojsModel
}

func initializePeople() []Person {
	return []Person{
		{
			ID:        1,
			FirstName: "Lionel Andres",
			LastName:  "Messi",
			Age:       36,
			CountryID: 1,
		},
		{
			ID:        2,
			FirstName: "Ronaldo ",
			LastName:  "de Assis Moreira",
			Age:       43,
			CountryID: 2,
		},
		{
			ID:        3,
			FirstName: "Luis Alberto",
			LastName:  "Suárez Díaz",
			Age:       36,
			CountryID: 3,
		},
		{
			ID:        4,
			FirstName: "James David",
			LastName:  "Rodríguez Rubio",
			Age:       32,
			CountryID: 4,
		},
	}
}

func initializeCountries() []Country {
	return []Country{
		{
			ID:   1,
			Name: "Argentina",
		},
		{
			ID:   2,
			Name: "Brazil",
		},
		{
			ID:   3,
			Name: "Colombia",
		},
		{
			ID:   4,
			Name: "Uruguay",
		},
	}
}
