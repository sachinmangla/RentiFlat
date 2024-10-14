package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sachinmangla/rentiflat/rentiflat"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/sachinmangla/rentiflat/docs"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome RentiFlat"))
	}).Methods("GET")

	router.HandleFunc("/add-owner", rentiflat.OwnerDetailCreatePost).Methods("POST")

	router.HandleFunc("/login", rentiflat.Login).Methods("POST")

	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.RentiFlatCreatePost)).ServeHTTP(w, r)
	}).Methods("POST")

	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.SearchFlat)).ServeHTTP(w, r)
	}).Methods("GET")

	router.HandleFunc("/update/{flat_id}", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.UpdateFlatDetail)).ServeHTTP(w, r)
	}).Methods("PUT")

	router.HandleFunc("/delete/{flat_id}", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.DeleteFlatPost)).ServeHTTP(w, r)
	}).Methods("DELETE")

	return router
}
