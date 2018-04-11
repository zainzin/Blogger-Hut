package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	blogRoute "./web/routes"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recent-blogs", blogRoute.BlogIndex)
	router.HandleFunc("/blog/{blogId}", blogRoute.OneBlog)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":3000", router))
}