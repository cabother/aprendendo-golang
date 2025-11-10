package externalapis

import (
	"cabother/aula/internal/dto"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetDogImage() (dto.CreateDogsDtoResponse, error) {
	url := "https://api.thedogapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "DEMO-API-KEY")

	resp, err := client.Do(req)
	if err != nil {
		return dto.CreateDogsDtoResponse{}, errors.New("error with url")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.CreateDogsDtoResponse{}, errors.New("error with url")
	}

	myResp := []dto.CreateDogApiDtoResponse{}
	err = json.Unmarshal(body, &myResp)
	if err != nil {
		return dto.CreateDogsDtoResponse{}, errors.New("error parsing dogs")
	}

	responseApi := dto.CreateDogDtoToCreateDogDtoResponse(myResp)

	return responseApi, nil
}
