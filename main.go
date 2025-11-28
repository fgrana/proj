package main

import (
	"database/sql"
	"fgrana/auth-project/internal/controller"
	"fgrana/auth-project/internal/controller/dbconfig"
	"fgrana/auth-project/internal/repository"
	"fgrana/auth-project/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func setupRouter() *gin.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("MYSQL_USER"),
		os.Getenv("ROOT_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"))
	dbName := os.Getenv("DATABASE")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	
	ok := dbconfig.DBconfig(db, dsn, dbName)
	if !ok {
		log.Panic("could not initialize tables")
	}

	r := gin.Default()
	r.Use(controller.SetUserMiddleware())

	repo := repository.NewRepository(db)
	s := service.NewService(*repo)
	controller.NewHandler(*s, r)

	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := setupRouter()
	r.Run(os.Getenv("SERVER_PORT"))
}
