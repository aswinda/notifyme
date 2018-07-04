package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/aswinda/notifyme/interfaces"
)

type MysqlHandler struct {
	Conn *sql.DB
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) 