package handler

import (
	"cabother/aula/internal/dto"
	"cabother/aula/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	CodeErr string `json:"codeErr"`
}

func RemoveUsers(c *gin.Context) {
	id := c.Param("id")

	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := gin.H{"error": fmt.Sprintf("invalid id %s", id)}
		c.JSON(http.StatusBadRequest, res)

		return
	}
	err = service.RemoveUserByID(idNumber)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error removing id %s", id), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id %s removido", id)})
}
func NewUser(c *gin.Context) {
	receivedBody := dto.CreateUserRequestBody{}

	err := c.BindJSON(&receivedBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	user := dto.CreateUserService{
		Name:      receivedBody.Name,
		BornDate:  receivedBody.BornDate,
		Status:    receivedBody.Status,
		Addresses: []dto.CreateAddressService{},
	}
	for _, addressReceived := range receivedBody.Address {
		currentAddressService := dto.CreateAddressService{}
		currentAddressService.Cep = addressReceived.Cep
		currentAddressService.Number = addressReceived.Number
		currentAddressService.Type = addressReceived.Type
		currentAddressService.Country = addressReceived.Country
		user.Addresses = append(user.Addresses, currentAddressService)
	}

	err = service.CreateUser(user)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error creating user %v", user), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user %v criado", user)})
}
func UpdateUser(c *gin.Context) {
	var receivedBody dto.UpdateUserRequestBody

	err := c.BindJSON(&receivedBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user := dto.UpdateUserRequestBody{
		Name:     receivedBody.Name,
		BornDate: receivedBody.BornDate,
		Status:   receivedBody.Status,
	}

	id := c.Param("id")
	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := gin.H{"error": fmt.Sprintf("invalid id %s", id)}
		c.JSON(http.StatusBadRequest, res)

		return
	}

	err = service.UpdateUserByID(idNumber, user)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error updating id %s", id), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user %v atualizado com sucesso", user)})
}

func GetUsersByID(c *gin.Context) {
	// receber o id
	id := c.Param("id")

	// converte o id que vai ser string para int
	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := gin.H{"error": fmt.Sprintf("invalid id %s", id)}
		c.JSON(http.StatusBadRequest, res)

		return
	}

	// chamar o serviço
	userModel, err := service.GetUserByID(idNumber)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error geting id %s", id), "error": err.Error()}
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, res)
			return
		}

		c.JSON(http.StatusInternalServerError, res)

		return
	}

	// converter o user de resposta para o tipo dto
	userResponse := dto.GetUserResponse{
		ID:       userModel.ID,
		Name:     userModel.Name,
		BornDate: userModel.BornDate,
		Status:   userModel.Status,
	}

	// retornar o user dto
	c.JSON(http.StatusOK, userResponse)
}

func GetAllUsers(c *gin.Context) {
	usersModel, err := service.GetAllUsers()
	if err != nil {
		res := gin.H{"message": "error geting users", "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	list := []dto.GetUsersResponse{}

	for _, item := range usersModel {
		if strings.Contains(item.Name, "a") {
			list = append(list, dto.GetUsersResponse{
				ID:       item.ID,
				Name:     item.Name,
				BornDate: item.BornDate,
				Status:   item.Status,
			})
		}

	}
	c.JSON(http.StatusOK, list)
}

func GetAllUsersAndBooks(c *gin.Context) {
	x, y := c.GetQuery("name")
	usersModel, err := service.GetAllUsersBooks(x)
	if y == false {

	}
	if err != nil {
		res := gin.H{"message": "error geting users", "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	response := dto.GetUsersAndBooksResponse{}
	for _, userModel := range usersModel {
		userAndAddressResponseItem := dto.UsersAndBooksResponse{}
		userAndAddressResponseItem.ID = userModel.ID
		userAndAddressResponseItem.Name = userModel.Name
		userAndAddressResponseItem.UserAddress = []dto.Address{}
		userAndAddressResponseItem.UserBooks = []dto.Book{}

		for _, addressModel := range userModel.Address {
			addressDto := dto.Address{}
			addressDto.Country = addressModel.Country
			addressDto.Neighborhood = addressModel.Neighborhood
			addressDto.Number = addressModel.Number
			addressDto.Street = addressModel.Street
			addressDto.Type = addressModel.Type

			userAndAddressResponseItem.UserAddress = append(userAndAddressResponseItem.UserAddress, addressDto)
		}

		for _, bookModel := range userModel.Books {
			bookDto := dto.Book{}
			bookDto.ID = bookModel.ID
			bookDto.Name = bookModel.Name
			userAndAddressResponseItem.UserBooks = append(userAndAddressResponseItem.UserBooks, bookDto)
		}

		if len(userAndAddressResponseItem.UserAddress) > 0 || len(userAndAddressResponseItem.UserBooks) > 0 {
			response.Users = append(response.Users, userAndAddressResponseItem)
		}
	}

	c.JSON(http.StatusOK, response)
}
func RemoveUsersByLikeName(c *gin.Context) {
	x, y := c.GetQuery("name")
	if y == false {
	}
	err := service.RemoveUsersByLikeName(x)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error removing user like name %s", x), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user like name %s removido", x)})
}

//http://google.com?name=renato -> Query Param
//http://google.com/name/renato -> Path Param

func RandomCep(c *gin.Context) {
	x := c.Param("number")
	numero, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("Erro ao converter a string:", err)
		return
	}
	err = service.RandomCep(numero)
	if err != nil {
		res := gin.H{"message": fmt.Sprintf("error search the cep", x), "error": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
}
