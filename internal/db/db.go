package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-grpc/internal/rocket"
	"os"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMODE := os.Getenv("DB_SSLMODE")

	connectionString := fmt.Sprintf("host=%s port=% user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUsername, dbTable,
		dbPassword, dbSSLMODE,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil

}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}
func (s Store) DeleteRocket(id string) error {
	return nil
}
