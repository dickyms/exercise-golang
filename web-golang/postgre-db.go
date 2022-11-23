package main

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

type student struct {
	first_name   string
	last_name string
	email string
}

func main() {
	connStr := "postgres://postgres:dckyms@localhost/go-test-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	_, err1 := db.Begin()
	if err1 != nil {
		log.Fatal(err1)
	}

	rows, err2 := db.Query(`select * from person`) 
	if err2 != nil {
		log.Println("error query")
		log.Fatal(err2)
	}
	defer rows.Close()

	//var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.first_name, &each.last_name, &each.email)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        log.Printf("first_name : %s last_name : %s email : %s \n", each.first_name, each.last_name, each.email)
	}
}


