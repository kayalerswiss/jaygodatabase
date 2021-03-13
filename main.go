package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("auto run")
}

func main() {
	//db, err := sql.Open("postgres", "root:password@tcp(127.0.0.1:3306)/tododb")
	url := os.Getenv("DATABASE_URL")
	fmt.Println("url: ", url)
	//db, err := sql.Open("postgres", "postgres://wanyjbxi:mh5jZwUhnMww5d_kqWp0rgLvKdqH7dwP@john.db.elephantsql.com:5432/wanyjbxi")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        title TEXT,
        status TEXT
    );
    `
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")
	fmt.Println("okay")
}
