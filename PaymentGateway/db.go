package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "dbname=mydatabase user=postgres password=mysecretpassword host=localhost sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Check the connection by pinging the database
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to database!")

	// Perform a query on the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Iterate over the rows returned by the query
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}
