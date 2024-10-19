package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/silvestrevivo/complete-go-webserver/internal/database"
)

func (apiCfg *apiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %s", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, DatabaseUserToUser(user))
}
