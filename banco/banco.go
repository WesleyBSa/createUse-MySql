package banco

import (
	"database/sql"

	_ "log"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:DevBook2024*@tcp(localhost:3006)/host"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
