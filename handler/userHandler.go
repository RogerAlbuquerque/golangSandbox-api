package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/postgresdatabase/database/models"
	"example.com/postgresdatabase/database/repositories"
)

type userHandler struct {
	Repo *repositories.UserRepository
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowerd", 405)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json", 400)
		return
	}

	err = h.Repo.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	jsonData, _ := json.MarshalIndent(user, "", "  ")

	fmt.Println(string(jsonData))

}
