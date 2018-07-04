package main

import (
	"sync"
	"github.com/aswinda/notifyme/controllers"
	"github.com/aswinda/notifyme/services"
	"github.com/aswinda/notifyme/repositories"
	"github.com/aswinda/notifyme/infrastructure"
	"database/sql"
)

type IServiceContainer interface {
	InjectUserController() controller.UserController
}

type kernel struct{}

func (k *kernel) InjectUserController() controller.UserController {
	
}