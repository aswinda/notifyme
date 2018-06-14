package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "fmt"
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) { }

func (a *App) Run(addr string) { }

// main function to boot up everything
func main() { 
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mysqlUsername := os.Getenv("MYSQL_USERNAME")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlDefaultDb := os.Getenv("MYSQL_DEFAULT_DB")

    fmt.Printf("%s", mysqlUsername)

    a := App{} 
    a.Initialize(mysqlUsername, mysqlPassword, mysqlDefaultDb)

    a.Run(":8080")
}