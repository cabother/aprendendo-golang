package models

import "fmt"

type AddressModel struct {
	Street       string
	Number       string
	Neighborhood string
	Country      string
	UserID       int64
	Type         string
}

func (a AddressModel) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.Number, a.Neighborhood, a.Country)
}
