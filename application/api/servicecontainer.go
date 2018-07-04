package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aswinda/notifyme/application/api/controllers"
	"github.com/aswinda/notifyme/application/api/infrastructures"
	"github.com/aswinda/notifyme/application/api/repositories"
	"github.com/aswinda/notifyme/application/api/services"
	"github.com/joho/godotenv"
)

type IServiceContainer interface {
	InjectUserController() controllers.UserController
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func (k *kernel) InjectUserController() controllers.UserController {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDefaultDb := os.Getenv("MYSQL_DEFAULT_DB")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDefaultDb)

	mysqlConn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = mysqlConn

	userRepository := &repositories.UserRepository{mysqlHandler}
	userService := &services.UserService{&repositories.UserRepositoryWithCircuitBreaker{userRepository}}
	userController := controllers.UserController{userService}

	return userController
}

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
