package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type RepositoryInterface interface {
	createUser(user User) (User, error)
	getUserByEmail(email string) (User, error)
	closeConnection()
}

func NewPostgresRepository() RepositoryInterface {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error connecting to the database")
	}
	r := &PostgresRepository{db}

	return r
}

type PostgresRepository struct {
	db *sqlx.DB
}

func (r PostgresRepository) closeConnection() {
	r.db.Close()
}

func (r PostgresRepository) createUser(user User) (User, error) {
	r.db.QueryRow(`INSERT INTO "user" (email, password) VALUES ($1, $2) RETURNING id`, user.Email, user.password).Scan(&user.Id)

	return user, nil
}

func (r PostgresRepository) getUserByEmail(email string) (User, error) {
	return User{}, nil
}
