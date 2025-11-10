package dto

type CreateDogsDtoResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CreateDogApiDtoResponse struct {
	Breeds []BreedsApiResponse `json:"breeds"`
	URL    string              `json:"url"`
}

type BreedsApiResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateDogDtoToCreateDogDtoResponse(dogsDto []CreateDogApiDtoResponse) CreateDogsDtoResponse {
	firstDog := dogsDto[0]

	return CreateDogsDtoResponse{
		Id:   firstDog.Breeds[0].ID,
		Name: firstDog.Breeds[0].Name,
		URL:  firstDog.URL,
	}
}
