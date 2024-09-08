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
	router.HandleFunc("/post", rentiflat.RentiFlatCreatePost)
	router.HandleFunc("/search", rentiflat.SearchFlat)
	return router
}
