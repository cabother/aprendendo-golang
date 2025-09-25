package dto

import (
	"time"
)

// CreateUserRequestBody
// @Description: DTO para criar um usuário na camada de handler
type CreateUserRequestBody struct {
	Name     string                         `json:"name"`
	BornDate time.Time                      `json:"bornDate"`
	Status   bool                           `json:"status"`
	Address  []CreateUserAddressRequestBody `json:"addresses"`
}

type CreateUserAddressRequestBody struct {
	Number  string `json:"number"`
	Country string `json:"country"`
	Type    string `json:"type" `
	Cep     int    `json:"cep"`
}
type CreateAddressApi struct {
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
}

// CreateUserService
// @Description: DTO para criar um usuário na camada de service
type CreateUserService struct {
	Name      string
	BornDate  time.Time
	Status    bool
	Addresses []CreateAddressService
}

type CreateAddressService struct {
	Street       string
	Number       string
	Neighborhood string
	Country      string
	Type         string
	Cep          int
}

type GetBooksService struct {
	ID         int64
	Name       string
	CategoryID int64
}

// UpdateUserRequestBody
// @Description: DTO para atualizar um usuário na camada de handler
type UpdateUserRequestBody struct {
	Name     string    `json:"name"`
	BornDate time.Time `json:"bornDate"`
	Status   bool      `json:"status"`
}

type GetUserResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	BornDate time.Time `json:"bornDate"`
	Status   bool      `json:"status"`
}
type GetUsersResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	BornDate time.Time `json:"bornDate"`
	Status   bool      `json:"status"`
}

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Address struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neigborhod"`
	Country      string `json:"country"`
	Type         string `json:"type"`
}
type UsersAndBooksResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	UserBooks   []Book    `json:"books"`
	UserAddress []Address `json:"address"`
}

type GetUsersAndBooksResponse struct {
	Users []UsersAndBooksResponse `json:"users"`
}
