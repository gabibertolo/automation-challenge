package model

type Country struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Person struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty" json:"age,omitempty"`
	CountryID int    `json:"country_id,omitempty" json:"country_id,omitempty"`
}
