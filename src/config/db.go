package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    dataSourceName := "seu_user:sua_senha@tcp(localhost:3306)/todo_db"
    var err error
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Erro ao pingar o banco de dados: %v", err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso")
}

func CloseDB() {
    err := DB.Close()
    if err != nil {
        log.Fatalf("Erro ao fechar conexão com o banco de dados: %v", err)
    }

    fmt.Println("Conexão com o banco de dados fechada com sucesso")
}