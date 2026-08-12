package main

import (
	"net/http"
	"sync/atomic"
        "github.com/joho/godotenv"
	"github.com/tiguco/chirpy/internal/database"
	_ "github.com/lib/pq"
        "database/sql"
        "log"
        "os"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db  *database.Queries
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatalf("Error loading .env file: %v", enverr)
	}

	dbURL := os.Getenv("DB_URL")
        db, err := sql.Open("postgres", dbURL)

	isDev := os.Getenv("PLATFORM")

        if err != nil {
                log.Fatalf("error connecting to db: %v", err)
        }
        defer db.Close()
        dbQueries := database.New(db)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db: dbQueries,
	}

	mux := http.NewServeMux()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	if isDev == "dev" {
		mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	}else{
		mux.HandleFunc("POST /admin/reset", apiCfg.noway)
	}

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsers)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
