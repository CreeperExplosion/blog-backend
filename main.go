package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"ghibran.xyz/blog-backend/database"
	"ghibran.xyz/blog-backend/router"
)

var (
	withSSL = flag.Bool("ssl", false, "served with TLS?")
	domain  = os.Getenv("DOMAIN")
)

func main() {
	flag.Parse()

	serve()
}

func init() {
	database.Init()
	router.Init()
}

func serve() {
	certFile := os.Getenv("CERT")
	keyFile := os.Getenv("KEY")

	router := router.GetRouter()

	if !*withSSL {
		log.Println("running on http://" + domain + "...")
		err := http.ListenAndServe(":80", router)

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	log.Println("running on https://" + domain + "...")
	go http.ListenAndServe(":80", http.HandlerFunc(toHTTPS))

	err := http.ListenAndServeTLS(":443", certFile, keyFile, router)

	if err != nil {
		log.Fatal(err)
	}

}

func toHTTPS(w http.ResponseWriter, r *http.Request) {

	target := "https://" + domain + r.URL.Path

	http.Redirect(w, r, target, http.StatusMovedPermanently)
}
