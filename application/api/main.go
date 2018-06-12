package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "fmt"
)

// main function to boot up everything
func main() { 
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    name := os.Getenv("NAME")
    fmt.Printf("%s %s", "Owl Notification", name)
}