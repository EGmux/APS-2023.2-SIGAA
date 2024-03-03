package credentials

import (
    "context"
    "log"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB() {
    var err error
    db, err = pgxpool.Connect(context.Background(), "postgresql://postgres:279135@localhost:5433/postgres")
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
}

func closeDB() {
    db.Close()
}

func IsValidUser(username, password string) bool {
    var storedPassword string
    err := db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
    if err != nil {
        if err == pgx.ErrNoRows {
            return false
        }
        log.Fatalf("Erro ao consultar o banco de dados: %v", err)
    }
    return password == storedPassword
}