package blog

import (
	"time"
	"database/sql"
	"log"
)

type Blog struct {
	ID    int    `sql:"id,pk"`
	Title string `sql:"title"`
	Body  string `sql:"body"`
	// Author Author `sql:author`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

func (b *Blog) Save(db *sql.DB) bool {
	sqlStatement := `INSERT INTO blog (title, body) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, b.Title, b.Body)
	return err != nil
}
// SELECT|people|age,name|age=?"
// SELECT|blog|id, title, body, created_at, updated_at|id = ?
func FetchBlogById(db *sql.DB, id int, blog *Blog) {
	err := db.QueryRow(`SELECT id, title, body, created_at, updated_at FROM blog WHERE id = $1`, id).Scan(&blog.ID,
		&blog.Title, &blog.Body, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}
}

func FetchRecentTenBlogs(db *sql.DB) ([]Blog, error) {
	sqlQuery := `SELECT id, title, body, created_at, updated_at FROM blog ORDER BY updated_at LIMIT 10`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var blogs []Blog
	var id int
	var title string
	var body string
	var createdAt time.Time
	var updatedAt time.Time
	for rows.Next() {
		err = rows.Scan(&id, &title, &body, &createdAt, &updatedAt)
		if err != nil {
			panic(err)
		}
		b := Blog{id, title, body, createdAt, updatedAt}
		blogs = append(blogs, b)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return blogs, err
}

func CreateBlogTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS blog 
	(id SERIAL NOT NULL, title VARCHAR(255), body TEXT, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW())`)
	if err != nil {
		log.Fatal(err)
	}
}
