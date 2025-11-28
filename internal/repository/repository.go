package repository

import (
	"database/sql"
	"fgrana/auth-project/internal/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	driver *sql.DB
}

func NewRepository(driver *sql.DB) *DB {
	db := &DB{
		driver: driver,
	}
	return db
}

var dbMoked = map[string]string{
	"foo":  "bar",
	"manu": "123",
}

func (db *DB) AddUser(user string) (string, bool) {

	stmt, err := db.driver.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user, "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User" + user + " inserted.")

	return "ok", true
}

func (db *DB) GetUser(c *gin.Context) (string, bool) {
	user := c.Params.ByName("name")
	if dbMoked[user] == "" {
		return "", false
	}
	return dbMoked[user], true
}
func (db *DB) GetAllUsers(c *gin.Context) (model.Users, bool) {
	rows, err := db.driver.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("All users:")
	var allUsers model.Users
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
			return allUsers, false
		}
		allUsers = append(allUsers, u)
		fmt.Printf("  %+v\n", u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return allUsers, false
	}
	return allUsers, true
}
