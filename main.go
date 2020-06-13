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

	router := router.GetRouter()
	port := os.Getenv("PORT")

	if !*withSSL {
		if port == "" {
			port = "80"
		}
		log.Println("running on http://" + domain + ":" + port + "...")
		err := http.ListenAndServe(":"+port, router)

		if err != nil {
			log.Fatal(err)
		}
		return
	}

	certFile := os.Getenv("CERT")
	keyFile := os.Getenv("KEY")

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
