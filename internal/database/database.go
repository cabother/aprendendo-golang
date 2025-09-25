package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectDB() (*sql.DB, error) {
	config := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aula",
		Password: "aula",
		DBName:   "auladatabase",
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("error connecting to the database" + err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// fmt.Println("Connected to the database successfully!")
	return db, nil
}
