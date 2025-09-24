package apis

import (
	"github.com/gorilla/mux"
)

func APIS() *mux.Router {
	r := mux.NewRouter()

	// here we can register all service level interface
	RegisterCache()

	// here we will define all APIs
	r.HandleFunc("/api/countries/search/{name}", GetCountry).Methods("GET", "OPTIONS")

	return r
}
