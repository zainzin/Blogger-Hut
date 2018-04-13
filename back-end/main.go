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
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recent-blogs", blogRoute.BlogIndex)
	router.HandleFunc("/blog/{blogId}", blogRoute.OneBlog)
	router.HandleFunc("/blogs", blogRoute.AllBlogs)
	router.HandleFunc("/create-blog", blogRoute.CreateBlog).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

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