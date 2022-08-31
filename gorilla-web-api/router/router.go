package router

import (
	"amazon/controller"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("API is live")
	})
	router.HandleFunc("/amazon/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/amazon/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/amazon/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/amazon/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/amazon/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")
	return router
}
