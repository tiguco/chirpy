package main

import (
	"net/http"
	"log"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))


        err := cfg.db.DeleteUsers(r.Context())
        if err != nil {
                log.Printf("couldn't delete users: %w", err)
                return
        }

}

func (cfg *apiConfig) noway(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("403 Forbidden"))
}
