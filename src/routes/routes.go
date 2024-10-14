package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sachinmangla/rentiflat/rentiflat"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome RentiFate"))
	})
	router.HandleFunc("/add-owner", rentiflat.OwnerDetailCreatePost)
	router.HandleFunc("/login", rentiflat.Login)
	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.RentiFlatCreatePost)).ServeHTTP(w, r)
	})
	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.SearchFlat)).ServeHTTP(w, r)
	})
	router.HandleFunc("/update/{flat_id}", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.UpdateFlatDetail)).ServeHTTP(w, r)
	})
	router.HandleFunc("/delete/{flat_id}", func(w http.ResponseWriter, r *http.Request) {
		rentiflat.Authenticate(http.HandlerFunc(rentiflat.DeleteFlatPost)).ServeHTTP(w, r)
	})
	return router
}
