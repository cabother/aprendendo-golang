package service

import (
	"cabother/aula/internal/dto"
	"cabother/aula/internal/models"
	"cabother/aula/internal/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func RemoveUserByID(id int64) error {
	if id < 1 {
		return fmt.Errorf("invalid id %d", id)
	}

	err := repository.RemoveUserByID(id)
	return err
}

func CreateUser(userService dto.CreateUserService) error {

	if len(userService.Name) < 2 {
		return fmt.Errorf("invalid name %s", userService.Name)
	}

	if userService.BornDate.IsZero() || userService.BornDate.After(time.Now()) || userService.BornDate.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return fmt.Errorf("invalid date %s", userService.BornDate)
	}

	userModel := models.UserModel{
		Name:     userService.Name,
		BornDate: userService.BornDate,
		Status:   userService.Status,
	}

	lastUserID, err := repository.CreateUser(userModel)
	if err != nil {
		return fmt.Errorf("error creating user name %s", userService.Name)
	}
	for _, address := range userService.Addresses {
		url := fmt.Sprintf("https://viacep.com.br/ws/%d/json/", address.Cep)
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error with cep")
		}
		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error with cep")
		}
		MyType := dto.CreateAddressApi{}
		json.Unmarshal(Body, &MyType)
		addressModel := models.AddressModel{}
		addressModel.Street = MyType.Street
		addressModel.Neighborhood = MyType.Neighborhood
		addressModel.Number = address.Number
		addressModel.Type = address.Type
		addressModel.Country = address.Country
		addressModel.UserID = lastUserID

		err = repository.CreateAddress(addressModel)
		if err != nil {
			return fmt.Errorf("error creating user address (%s)", addressModel.GetFullAddress())
		}
	}

	return err
}
func UpdateUserByID(id int64, user dto.UpdateUserRequestBody) error {

	userModel := models.UserModel{
		Name:     user.Name,
		BornDate: user.BornDate,
		Status:   user.Status,
	}

	if len(user.Name) < 2 {
		return fmt.Errorf("invalid name %s", user.Name)
	}

	if user.BornDate.IsZero() || user.BornDate.After(time.Now()) || user.BornDate.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return fmt.Errorf("invalid date %s", user.BornDate)
	}

	if id < 1 {
		return fmt.Errorf("invalid id %d", id)
	}

	err := repository.UpdateUserByID(id, userModel)
	return err
}
func GetUserByID(id int64) (models.UserModel, error) {

	if id < 1 {
		return models.UserModel{}, fmt.Errorf("invalid id %d", id)
	}

	user, err := repository.GetUserByID(id)
	return user, err
}
func GetAllUsers() ([]models.UserModel, error) {
	users, err := repository.GetAllUsers()

	return users, err
}

func GetAllUsersBooks(name string) ([]models.UserModel, error) {
	users, err := repository.GetAllUsersByLike(name)
	if err != nil {
		return []models.UserModel{}, err
	}

	for i, user := range users {
		livros, err := repository.GetBooksByUserID(user.ID)
		if err != nil {
			return []models.UserModel{}, err
		}

		users[i].Books = livros
	}
	for i, user := range users {
		Address, err := repository.GetAddressByUserID(user.ID)
		if err != nil {
			return []models.UserModel{}, err
		}
		users[i].Address = Address
	}

	return users, nil
}
func RemoveUsersByLikeName(name string) error {
	if len(name) < 1 {
		return fmt.Errorf("invalid like name %s", name)
	}

	err := repository.RemoveUsersByLikeName(name)
	return err
}
