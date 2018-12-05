package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DB do aplicativo ...
type DB struct {
	DB *sql.DB
}

//Connection ...
func (a *DB) Connection() error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "catavento", "localhost", 3306, "ouvidoria")
	var err error

	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("[db/OpenConnection] - Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return err
	}
	return nil
}
