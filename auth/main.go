package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var Service UserService

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := NewPostgresRepository()
	defer r.closeConnection()
	Service = NewUserService(r)

	http.HandleFunc("/auth/healthcheck", healthcheck)
	http.HandleFunc("/auth/login", login)
	http.HandleFunc("/auth/register", register)

	http.ListenAndServe(":8081", nil)
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{ \"status\": \"ok\"}")
}

func login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u, _ := Service.NewUser(UserDTO{})
	json.NewEncoder(w).Encode(u)
}

func register(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u UserDTO
	err := json.NewDecoder(req.Body).Decode(&u)
	fmt.Printf("%+v\n", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, error := Service.NewUser(u)
	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
