package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ismae1-alvarez/rest-api-go/internal/user"
	"github.com/ismae1-alvarez/rest-api-go/pkg/bootstrap"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	_ = godotenv.Load()

	// definir mi log
	l := bootstrap.InitLogger()

	db, err := bootstrap.DBConnection()

	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepo(l, db)

	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndPoints(userSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
