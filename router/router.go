package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

// Init initialize the routes
func Init() {
	router = mux.NewRouter()
	setRoutes()
}

// GetRouter returns the mux router
func GetRouter() *mux.Router {
	return router
}

func setRoutes() {

	router.HandleFunc("/", getRoot).Methods("GET")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Algebra's blog API")
}
