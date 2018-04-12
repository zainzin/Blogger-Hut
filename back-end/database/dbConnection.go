package database

import (
	"os"
	"encoding/json"
	"log"
	"sync"
	"../database/blog"
	"../database/author"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func init() {
	once.Do(func() {
		conf := readConfFile()
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			conf.Address, conf.Port, conf.Username, conf.Password, conf.Database)
		var err error
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(fmt.Sprintf("dbConnection - init: %s", err))
			panic(fmt.Sprintf("dbConnection - init: %s", err))
		}
		log.Println("CONNECTED TO DATABASE")

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		blog.CreateBlogTable(db)
		author.CreateAuthorTable(db)
	})
}

type DatabaseConfiguration struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func readConfFile() DatabaseConfiguration {
	file, _ := os.Open("db-conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := DatabaseConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(fmt.Sprintf("dbConnection - readConfFile: %s", err))
		panic(fmt.Sprintf("dbConnection - readConfFile: %s", err))
	}
	return configuration
}

func GetConnection() *sql.DB {
	return db
}

func CloseConnection() error {
	return db.Close()
}
