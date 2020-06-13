package router

import (
	"fmt"
	"log"
	"net/http"
)

func setRoutes() {
	log.Println("initializing routes.....")

	router.HandleFunc("/", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}
