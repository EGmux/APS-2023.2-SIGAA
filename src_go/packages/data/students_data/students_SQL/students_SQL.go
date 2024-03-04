package data_studentsdata_studentsSQL

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	student "sigaa.ufpe/packages/data/students_data"
)

var db *pgxpool.Pool

func InitDB() {
    var err error
    db, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }
}

func CloseDB() {
    db.Close()
}

func IsValidUser(username string, password string) bool {
    var storedPassword string
    err := db.QueryRow(context.Background(), "SELECT pswrd FROM students WHERE usr = $1", username).Scan(&storedPassword)
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

func InsertUser(username string, password string) {
	if !IsValidUser(username, password){
		db.Exec(context.Background(), "INSERT INTO students (usr, pswrd) VALUES ($1, $2)", username, password)
	}
	 
}

func GetAllUsers() ([]student.Student) {
	rows, err := db.Query(context.Background(), "SELECT usr, pswrd FROM students")
	if err != nil {
		log.Fatal("Student_SQL.go : It was not possible to realize the querry")
	}
	defer rows.Close()

	var users []student.Student
	for rows.Next() {
		var user student.Student
		err := rows.Scan(&user.User, &user.Password)
		if err != nil {
			log.Fatal("Student_SQL.go: Erros during row scan")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Student_SQL.go: Error during row scan")
	}

	return users
}