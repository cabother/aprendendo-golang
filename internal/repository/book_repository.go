package repository

import (
	"cabother/aula/internal/database"
	"cabother/aula/internal/models"
)

func GetBooksByUserID(userID int64) ([]models.BookModel, error) {
	execution := `select id, name from books where id in (select book_id from user_books where user_id = $1)`

	banco, _ := database.ConnectDB()
	result, err := banco.Query(execution, userID)
	if err != nil {
		return []models.BookModel{}, err
	}

	var books []models.BookModel
	for result.Next() {
		var temporaryBook models.BookModel

		err = result.Scan(&temporaryBook.ID, &temporaryBook.Name)
		if err != nil {
			return []models.BookModel{}, err
		}

		books = append(books, temporaryBook)
	}

	return books, nil
}
