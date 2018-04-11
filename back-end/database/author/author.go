package author

import (
	"time"
	"database/sql"
	"log"
)

type Author struct {
	ID        int       `sql:"id,pk"`
	Name      string    `sql:"name"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

func (a *Author) Save(db *sql.DB) bool {
	sqlStatement := `INSERT INTO blog (title, body) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, a.Name)
	return err != nil
}

func FetchAuthorById(db *sql.DB, id int, author *Author) {
	err := db.QueryRow(`SELECT id, name, created_at, updated_at FROM author WHERE id = $1`, id).Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}
}

func FetchAuthors(db *sql.DB, skip int, limit int) []Author {
	var authors []Author
	rows, err := db.Query(`SELECT id, name FROM author ORDER BY name LIMIT $1 OFFSET $2`, limit, skip)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var author Author
		rows.Scan(&author.ID, &author.Name)
		authors = append(authors, author)
	}
	return authors
}

func CreateAuthorTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS author 
	(id SERIAL NOT NULL, name VARCHAR(50), created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW())`)
	if err != nil {
		log.Fatal(err)
	}
}
