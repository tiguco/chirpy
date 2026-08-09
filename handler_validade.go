package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func unprofanator(msg string) string {
	re1 := regexp.MustCompile(`(?i)\bkerfuffle\b`)
	newStr1 := re1.ReplaceAllString(msg, "****")
	re2 := regexp.MustCompile(`(?i)\bsharbert\b`)
	newStr2 := re2.ReplaceAllString(newStr1, "****")
	re3 := regexp.MustCompile(`(?i)\bfornax\b`)
	newStr3 := re3.ReplaceAllString(newStr2, "****")
	return newStr3
}


func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedmsg := unprofanator(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		Cleaned_body: cleanedmsg,
	})
}
