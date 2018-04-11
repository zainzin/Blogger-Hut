package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	db "./database"
	"./database/blog"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recent-blogs", BlogIndex)
	router.HandleFunc("/blog/{blogId}", TodoShow)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":3000", router))
}

func BlogIndex(w http.ResponseWriter, r *http.Request) {
	b := blog.Blog{
		Title: "test",
		Body:  "tester",
	}
	b.Save(db.GetConnection())
	blogs, err := blog.FetchRecentTenBlogs(db.GetConnection())
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			error bool
		}{error: true})
	}
	json.NewEncoder(w).Encode(blogs)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId, _ := strconv.Atoi(vars["blogId"])
	var b blog.Blog
	blog.FetchBlogById(db.GetConnection(), blogId, &b)
	json.NewEncoder(w).Encode(b)
}
