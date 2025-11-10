package business

import (
	"cabother/aula/internal/dto"
	"cabother/aula/internal/externalapis"
	"cabother/aula/internal/models"
	"cabother/aula/internal/repository"
	"fmt"
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

	// Cria o usuario
	lastUserID, err := repository.CreateUser(userModel)
	if err != nil {
		return fmt.Errorf("error creating user name %s", userService.Name)
	}

	// Cria o doguinho
	respDog, err := externalapis.GetDogImage()
	if err != nil {
		return err
	}

	dogModel := models.DogModel{
		Name:   respDog.Name,
		UserID: lastUserID,
		Photo:  respDog.URL,
		DogID:  int64(respDog.Id),
	}

	err = repository.CreateDog(dogModel)
	if err != nil {
		return err
	}

	// Cria os endereços
	for _, address := range userService.Addresses {
		resp, err := externalapis.FindCep(fmt.Sprintf("%d", address.Cep))
		if err != nil {
			return err
		}

		addressModel := models.AddressModel{}
		addressModel.Street = resp.Street
		addressModel.Neighborhood = resp.Neighborhood
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

func RandomCep(number int) error {
	initialValue := 14400000

	for i := 0; i < number; i++ {
		initialValue = initialValue + 1
		cepInfo, err := externalapis.FindCep(fmt.Sprintf("%d", initialValue))
		if err != nil {
			break
		}
		if cepInfo.Street != "" {
			repository.CreateAddressCep(cepInfo)
		}
		fmt.Println(cepInfo)

		// Inserir cada cep encontrado no banco de dados, na tabela address
	}

	return nil
}

func FindCep(cep string) (models.CepModel, error) {
	ceps, err := repository.GetCep(cep)
	if err != nil {
		return models.CepModel{}, err
	}

	if len(ceps) > 0 {
		response := ceps[0]
		response.Origin = "database"

		return response, nil
	}

	externalCepFound, err := externalapis.FindCep(cep)
	if err != nil {
		return models.CepModel{}, err
	}

	if externalCepFound.Cep == "" {
		return models.CepModel{}, fmt.Errorf("error search the cep %s, cep not found", cep)
	}

	address := dto.CreateAddressApi{
		Cep:          cep,
		Street:       externalCepFound.Street,
		Neighborhood: externalCepFound.Neighborhood,
		City:         externalCepFound.City,
	}

	err = repository.CreateAddressCep(address)
	if err != nil {
		return models.CepModel{}, err
	}

	response := models.CepModel{
		Cep:          externalCepFound.Cep,
		Street:       externalCepFound.Street,
		Neighborhood: externalCepFound.Neighborhood,
		City:         externalCepFound.City,
		Origin:       "api-externa",
	}

	return response, err

}
