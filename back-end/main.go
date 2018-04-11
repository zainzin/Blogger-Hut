package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	blogRoute "./web/routes"
	db "./database"
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recent-blogs", blogRoute.BlogIndex)
	router.HandleFunc("/blog/{blogId}", blogRoute.OneBlog)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Println(http.ListenAndServe(":3000", router))

	go func() {
		runtime.SetFinalizer(router, func() {
			err := db.CloseConnection()
			if err != nil {
				log.Fatal("Connection Wrong")
			} else {
				log.Println("Connection to database was successfully closed")
			}
			runtime.GC()
		})
	}()

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("Signal %v", <-c)
	}()

	log.Println("Exit:", <-errc)
}