package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	server := "server_address"  // server address
	port := 1433                //default port
	user := "username"          //your DB username
	password := "password"      //your DB password
	database := "database_name" // name of your Database

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	for true {
		rows, err := conn.Query("SELECT COUNT(*) FROM table_name") //change table_name to name of your users table
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var count sql.NullInt64
			if err := rows.Scan(&count); err != nil {
				log.Fatal(err)
			}

			fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf(" || User count: %d\n", count.Int64)

		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(8 * time.Second)
	}
}
