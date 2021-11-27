package router

import (
	"go-postgres/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/buku", controller.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.GetBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku", controller.AddBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
