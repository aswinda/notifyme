package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "fmt"
    "."
)

// main function to boot up everything
func main() { 
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mysqlUsername := os.Getenv("MYSQL_USERNAME")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlDefaultDb := os.Getenv("MYSQL_DEFAULT_DB")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlPort := os.Getenv("MYSQL_PORT")

    fmt.Printf("%s", mysqlUsername)

    a := App{} 
    a.Initialize(mysqlUsername, mysqlPassword, mysqlDefaultDb, mysqlHost, mysqlPort)

    a.Run(":8080")
}
