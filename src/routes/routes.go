package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome RentiFate"))
	})
	return router
}
