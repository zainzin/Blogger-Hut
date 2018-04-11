package routes

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	author "../../database/author"
	db "../../database"
)

func AuthorIndex(w http.ResponseWriter, r *http.Request) {
	a := author.Author{
		Name: "test author",
	}
	a.Save(db.GetConnection())
	skip, _ := strconv.Atoi(mux.Vars(r)["skip"])
	limit, _ := strconv.Atoi(mux.Vars(r)["limit"])
	blogs, err := author.FetchAuthors(db.GetConnection(), limit, skip)
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			error bool
		}{error: true})
	}
	json.NewEncoder(w).Encode(blogs)
}