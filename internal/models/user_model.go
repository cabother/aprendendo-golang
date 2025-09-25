package models

import (
	"time"
)

type UserModel struct {
	ID       int64
	Name     string
	BornDate time.Time
	Status   bool
	Books    []BookModel
	Address  []AddressModel
}
