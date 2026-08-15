package main

import (
	"encoding/json"
//	"errors"
//	"fmt"
	"net/http"
//	"strings"
	"time"

	//"github.com/bootdotdev/learn-http-servers/internal/database"
//	"github.com/tiguco/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp2 struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
//	type parameters struct {
//		Body   string    `json:"body"`
//		UserID uuid.UUID `json:"user_id"`
//	}

	chirps, err := cfg.db.GetChirps(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}


	lstChirps := []Chirp2{}
	for _, mychirp := range chirps {
		tmpChirp := Chirp2{}
		tmpChirp.ID = mychirp.ID
		tmpChirp.CreatedAt = mychirp.CreatedAt
		tmpChirp.UpdatedAt = mychirp.UpdatedAt
		tmpChirp.Body = mychirp.Body
		lstChirps = append(lstChirps, tmpChirp)
//		fmt.Printf("id: %s \n", mychirp.ID)
//		fmt.Printf("created_at: %v \n", mychirp.CreatedAt)
//		fmt.Printf("update_at: %v \n", mychirp.UpdatedAt)
//		fmt.Printf("body: %s \n", mychirp.Body)
//		fmt.Printf("user_id: %s \n", mychirp.UserID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lstChirps)
	return

//	respondWithJSON(w, http.StatusCreated, Chirp{
//		ID:        chirp.ID,
//		CreatedAt: chirp.CreatedAt,
//		UpdatedAt: chirp.UpdatedAt,
//		Body:      chirp.Body,
//		UserID:    chirp.UserID,
//	})

}

//func validateChirp(body string) (string, error) {
//	const maxChirpLength = 140
//	if len(body) > maxChirpLength {
//		return "", errors.New("Chirp is too long")
//	}
//
//	badWords := map[string]struct{}{
//		"kerfuffle": {},
//		"sharbert":  {},
//		"fornax":    {},
//	}
//	cleaned := getCleanedBody(body, badWords)
//	return cleaned, nil
//}
//
//func getCleanedBody(body string, badWords map[string]struct{}) string {
//	words := strings.Split(body, " ")
//	for i, word := range words {
//		loweredWord := strings.ToLower(word)
//		if _, ok := badWords[loweredWord]; ok {
//			words[i] = "****"
//		}
//	}
//	cleaned := strings.Join(words, " ")
//	return cleaned
//}
