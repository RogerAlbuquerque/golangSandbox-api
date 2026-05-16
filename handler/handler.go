package handler

import (
	"log"
	"net/http"

	"example.com/postgresdatabase/database"
	"example.com/postgresdatabase/database/repositories"
)

func Initialize() {
	database, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repositories.UserRepository{
		DB: database,
	}

	userHandler := userHandler{
		Repo: userRepo,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.CreateUser)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Teste"))
	})

	print("Server running on :8080\n")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
