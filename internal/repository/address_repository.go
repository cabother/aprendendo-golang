package repository

import (
	"cabother/aula/internal/database"
	"cabother/aula/internal/models"
	"fmt"
)

func CreateAddress(address models.AddressModel) error {
	banco, err := database.ConnectDB()
	if err != nil {
		return err
	}

	execution := `insert into address(street, number, neighborhood, country, user_id, type) 
		values($1, $2, $3, $4, $5, $6)`

	result, err := banco.Exec(execution, address.Street, address.Number, address.Neighborhood, address.Country, address.UserID, address.Type)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("address not created")
	}

	return nil
}
func GetAddressByUserID(userID int64) ([]models.AddressModel, error) {
	execution := `select street, number, neighborhood, country, type from address where user_id = $1`
	banco, err := database.ConnectDB()
	if err != nil {
		return []models.AddressModel{}, err
	}
	result, err := banco.Query(execution, userID)
	if err != nil {
		return []models.AddressModel{}, err
	}

	var Addresses []models.AddressModel

	for result.Next() {
		var address models.AddressModel
		err = result.Scan(&address.Street, &address.Number, &address.Neighborhood, &address.Country, &address.Type)
		if err != nil {
			return []models.AddressModel{}, err
		}

		Addresses = append(Addresses, address)
	}

	return Addresses, nil
}
