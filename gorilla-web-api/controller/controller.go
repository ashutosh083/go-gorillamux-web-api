package controller

import (
	"amazon/database"
	"amazon/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	//"go.chromium.org/luci/gae/filter/count"
	//"google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1/schema/predict/params"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "movie/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allMovies := database.Find()
	json.NewEncoder(w).Encode(allMovies)

}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "movie/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Amazon
	_ = json.NewDecoder(r.Body).Decode(&movie)
	database.Create(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "movie/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	database.Update(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "movie/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	database.Delete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "movie/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	database.DeleteAll()
	json.NewEncoder(w).Encode("")
}
