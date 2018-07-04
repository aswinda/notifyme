package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/aswinda/notifyme/interfaces"
)

type MysqlHandler struct {
	Conn *sql.DB
}

type MysqlRow struct {
	Rows *sql.Rows
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string) (interface.IRow, err) {
	rows, err := handler.Conn.Query(statement)

	if err := nill {
		fmt.Println(err)
		return new(MysqlRow), err
	}

	row := new(MysqlRow)
	row.Rows = rows

	return row, nil
}

func (r MysqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MysqlRow) Next() bool {
	return r.Rows.Next()
}