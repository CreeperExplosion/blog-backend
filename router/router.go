package router

import (
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
