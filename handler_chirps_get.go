package main

import (
	"net/http"
	"github.com/google/uuid"
        "time"
)


//import (
//        "encoding/json"
//        "errors"
//        "net/http"
//        "strings"
//        "time"
//        //"github.com/bootdotdev/learn-http-servers/internal/database"
//        "github.com/tiguco/chirpy/internal/database"
//        "github.com/google/uuid"
//)
// 



type Chirp2 struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}


func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
        
    idStr := r.PathValue("chirpID")
    
    if idStr == "" {
        http.Error(w, "Missing resource ID", http.StatusBadRequest)
        return
    }


    // Parse the UUID
    id, err := uuid.Parse(idStr)
    if err != nil {
	    http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	    return
    }

    chirp, err := cfg.db.GetChirp(r.Context(), id)

    myChirp := Chirp2{
                ID:        chirp.ID,
                CreatedAt: chirp.CreatedAt,
                UpdatedAt: chirp.UpdatedAt,
                Body:      chirp.Body,
                UserID:    chirp.UserID,
    }

    if err != nil {
	    respondWithError(w, http.StatusNotFound, "Chirp not found", err)
	    return
    }

    respondWithJSON(w, http.StatusOK, myChirp)

}
