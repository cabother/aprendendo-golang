package externalapis

import (
	"cabother/aula/internal/dto"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func FindCep(x int) (dto.CreateAddressApi, error) {
	number := x
	url := fmt.Sprintf("https://viacep.com.br/ws/%d/json/", number)
	resp, err := http.Get(url)
	if err != nil {
		return dto.CreateAddressApi{}, errors.New("error with cep")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.CreateAddressApi{}, errors.New("error with cep")
	}
	myType := dto.CreateAddressApi{}
	json.Unmarshal(body, &myType)
	return myType, nil
}
