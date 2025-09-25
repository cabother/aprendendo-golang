package repository

import (
	"cabother/aula/internal/database"
	"cabother/aula/internal/models"
	"fmt"
	"strconv"
)

func RemoveUserByID(id int64) error {
	execution := `delete from users where id = $1`
	banco, _ := database.ConnectDB()
	result, err := banco.Exec(execution, id)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("id %d not found", id)
	}

	return nil
}

func CreateUser(user models.UserModel) (int64, error) {
	banco, err := database.ConnectDB()
	if err != nil {
		return 0, err
	}

	status := strconv.FormatBool(user.Status)
	execution := `insert into users(name, born_date, status) 
		values('` + user.Name + `','` + user.BornDate.Format("2006-01-02") + `',` + status + `) RETURNING id`

	var id int64
	err = banco.QueryRow(execution).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateUserByID(id int64, user models.UserModel) error {
	banco, err := database.ConnectDB()
	if err != nil {
		return err
	}

	status := strconv.FormatBool(user.Status)
	query := `update users set name = $1, born_date = $2, status = $3 where id = $4`

	result, err := banco.Exec(query, user.Name, user.BornDate.Format("2006-01-02"), status, id)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not updated")
	}

	return nil
}

func GetUserByID(id int64) (models.UserModel, error) {
	execution := `Select id, name, born_date, status from users where id = $1`
	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution, id)
	if err != nil {
		return models.UserModel{}, err
	}

	var user models.UserModel
	for result.Next() {
		err = result.Scan(&user.ID, &user.Name, &user.BornDate, &user.Status)
		if err != nil {
			return models.UserModel{}, err
		}
	}

	if user.ID == 0 {
		return models.UserModel{}, fmt.Errorf("id %d not found", id)
	}

	return user, nil
}

func GetAllUsers() ([]models.UserModel, error) {
	execution := `Select id, name, born_date, status from users`
	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution)
	if err != nil {
		return []models.UserModel{}, err
	}

	var users []models.UserModel

	for result.Next() {
		var user models.UserModel
		err = result.Scan(&user.ID, &user.Name, &user.BornDate, &user.Status)
		if err != nil {
			return []models.UserModel{}, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return []models.UserModel{}, fmt.Errorf("not found")
	}

	return users, nil
}
func GetAllUsersByLike(filter string) ([]models.UserModel, error) {
	value := "%" + filter + "%"
	execution := `Select id, name, born_date, status from users where name like $1`
	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution, value)
	if err != nil {
		return []models.UserModel{}, err
	}

	var users []models.UserModel

	for result.Next() {
		var user models.UserModel
		err = result.Scan(&user.ID, &user.Name, &user.BornDate, &user.Status)
		if err != nil {
			return []models.UserModel{}, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return []models.UserModel{}, fmt.Errorf("not found")
	}

	return users, nil
}
func RemoveUsersByLikeName(name string) error {
	filter := "%" + name + "%"
	execution := `delete from users where name like $1`
	banco, _ := database.ConnectDB()
	result, err := banco.Exec(execution, filter)
	if err != nil {
		return err
	}

	rowsAffected, errRowsAffected := result.RowsAffected()
	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rowsAffected == 0 {
		return fmt.Errorf("like name %s not found", filter)
	}

	return nil
}
