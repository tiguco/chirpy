package main

import (
	"encoding/json"
	"time"
	"log"
	"net/http"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {

       type parameters struct {
	       Email string `json:"email"`
       }


       type User struct {
	       Id        uuid.UUID `json:"id"`
	       Created_at time.Time `json:"created_at"`
	       Updated_at time.Time `json:"updated_at"`
	       Email     string `json:"email"`
       }

        decoder := json.NewDecoder(r.Body)
        params := parameters{}
        err := decoder.Decode(&params)
        if err != nil {
                //respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
                log.Printf("couldn't create user: %w", err)
		return
        }


	retuser, err := cfg.db.CreateUser(r.Context(), params.Email)
        if err != nil {
                log.Printf("couldn't create user: %w", err)
                return 
        }
	newuser := User{}
	newuser.Id = retuser.ID
	newuser.Created_at = retuser.CreatedAt
	newuser.Updated_at = retuser.UpdatedAt
	newuser.Email = retuser.Email


        w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
        dat, err := json.Marshal(newuser)
        if err != nil {
                log.Printf("Error marshalling JSON: %s", err)
                w.WriteHeader(500)
                return
        }
        w.Write(dat)

        return 
}
