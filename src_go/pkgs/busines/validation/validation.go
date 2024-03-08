package validation

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

func IsValidUser(username string, password string) bool {
	var storedPassword string
	err := db.QueryRow(context.Background(), "SELECT pswrd FROM students WHERE usr = $1", username).
		Scan(&storedPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("DATABASE RETURNED FALSE")
			return false
		}
		log.Fatalf("Erro ao consultar o banco de dados: %v", err)
	}
	fmt.Println(password, storedPassword)
	return password == storedPassword
}
