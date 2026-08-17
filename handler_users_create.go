package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tiguco/chirpy/internal/auth"
	"github.com/tiguco/chirpy/internal/database"
)

type User2 struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parametersIn struct {
		Password string `json:"password"`
		Email string `json:"email"`
	}
//	type response struct {
//		User2
//	}

	decoder := json.NewDecoder(r.Body)
	paramsIn := parametersIn{}
	err := decoder.Decode(&paramsIn)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	params := database.CreateUserParams{}
	params.Email = paramsIn.Email

	hashedPass, err := auth.HashPassword(paramsIn.Password)
	params.HashedPassword = hashedPass

	user, err := cfg.db.CreateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}
	//response := database.CreateUserRow{
	response := User2{
			ID:        user[0].ID,
			CreatedAt: user[0].CreatedAt,
			UpdatedAt: user[0].UpdatedAt,
			Email:     user[0].Email}


	respondWithJSON(w, http.StatusCreated, response)
}
