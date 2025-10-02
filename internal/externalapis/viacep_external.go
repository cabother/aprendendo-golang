package externalapis

import (
	"cabother/aula/internal/dto"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func FindCep(cep string) (dto.CreateAddressApi, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return dto.CreateAddressApi{}, errors.New("error with cep")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.CreateAddressApi{}, errors.New("error with cep")
	}
	myCepInfo := dto.CreateAddressApi{}
	json.Unmarshal(body, &myCepInfo)

	return myCepInfo, nil
}
