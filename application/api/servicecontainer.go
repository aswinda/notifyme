package main

import (
	"sync"
	"github.com/aswinda/notifyme/controllers"
	"github.com/aswinda/notifyme/services"
	"github.com/aswinda/notifyme/repositories"
	"github.com/aswinda/notifyme/infrastructure"
	"database/sql"
	"github.com/joho/godotenv"
    "log"
    "os"
)

type IServiceContainer interface {
	InjectUserController() controller.UserController
}

type kernel struct{}

func (k *kernel) InjectUserController() controller.UserController {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mysqlUsername := os.Getenv("MYSQL_USERNAME")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlDefaultDb := os.Getenv("MYSQL_DEFAULT_DB")
    mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:8082)/%s", user, password, dbname)
    var err error
    mysqlConn , err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
	}
	
	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = mysqlConn

	userRepository := &repositories.UserRepository{mysqlHandler}

}