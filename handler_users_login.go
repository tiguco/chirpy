package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/tiguco/chirpy/internal/auth"
//	"github.com/tiguco/chirpy/internal/database"
        "github.com/google/uuid"
        "time"
)

type User3 struct {
        ID        uuid.UUID `json:"id"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
        Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	//params := database.CreateUserParams{}

        type parametersIn struct {
                Password string `json:"password"`
                Email string `json:"email"`
        }
	params := parametersIn{}
	err := decoder.Decode(&params)
	fmt.Println("pass: " + params.Password)
	fmt.Println("email: " + params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	dbhashedPass, err := cfg.db.GetHashedPasswordFromEmail(r.Context(), params.Email)

	fmt.Println("dbhashedpass: " + dbhashedPass)

	passou, err := auth.CheckPasswordHash(params.Password, dbhashedPass)

	if passou {
		user, err := cfg.db.GetUserFromEmail(r.Context(), params.Email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve user", err)
			return
		}

//		response := database.CreateUserRow{
//			ID:        user.ID,
//			CreatedAt: user.CreatedAt,
//			UpdatedAt: user.UpdatedAt,
//			Email:     user.Email}


               response := User2{
                        ID:        user.ID,
                        CreatedAt: user.CreatedAt,
                        UpdatedAt: user.UpdatedAt,
                        Email:     user.Email}


		respondWithJSON(w, http.StatusOK, response)

	}else{
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
	}
}
