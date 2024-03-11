package singleton_db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)


var db *pgxpool.Pool

func InitDB(){
	var err error
	db, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Fail")
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
}


func Refer_DB() *pgxpool.Pool{
	return db
}

func CloseDB(){
	db.Close()
}