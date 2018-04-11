package routes

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	blog "../../database/blog"
	db "../../database"
)

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

func OneBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId, _ := strconv.Atoi(vars["blogId"])
	var b blog.Blog
	blog.FetchBlogById(db.GetConnection(), blogId, &b)
	json.NewEncoder(w).Encode(b)
}