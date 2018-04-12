package routes

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"../../database/blog"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func OneBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId, _ := strconv.Atoi(vars["blogId"])
	var b blog.Blog
	blog.FetchBlogById(db.GetConnection(), blogId, &b)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func AllBlogs(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.Atoi(r.FormValue("skip"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	if limit == 0 {
		limit = 10
	}

	if skip <= 1 {
		skip = 1
	}

	blogs, err, totalRows := blog.FetchBlogs(db.GetConnection(), (skip - 1) * 10, limit)

	if err != nil {
		json.NewEncoder(w).Encode(struct {
			error bool
		}{error: true})
	}

	resp := struct {
		TotalRows int
		CurrentPage int
		Blogs []blog.Blog
	}{
		TotalRows:totalRows,
		CurrentPage: skip,
		Blogs:blogs,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}