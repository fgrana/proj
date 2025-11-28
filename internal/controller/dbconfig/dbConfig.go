package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
)

func DBconfig(db *sql.DB, dsn string, dbName string) bool {

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	fmt.Println("Database 'users' ensured.")

	// Connect to the newly created database
	dbWithDB, err := sql.Open("mysql", dsn+dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbWithDB.Close()

	// Create the table 'user' if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = dbWithDB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table 'user' ensured.")

	stmt := `USE users;`

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
		return false
	}
	

	return true
}
