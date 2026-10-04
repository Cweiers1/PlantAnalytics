package server

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func homeHandler(w http.ResponseWriter, r *http.Request, pool pgxpool.Pool) {
	w.Write([]byte("Test"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func nodeIngestHandler(w http.ResponseWriter, r *http.Request) {

}

func getPlantHandler(w http.ResponseWriter, r *http.Request) {

}

func addPlantHandler(w http.ResponseWriter, r *http.Request) {

}

func Start(pool *pgxpool.Pool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, pool)
	})
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Println("listening on :5000")
	log.Fatal(server.ListenAndServe())
}
