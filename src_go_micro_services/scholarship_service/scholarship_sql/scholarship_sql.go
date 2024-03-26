package scholarshipsql

import (
	"context"
	"fmt"
	"log"
	"os"
	"scholarship/student"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)


var db *pgxpool.Pool

func Init_DB(){

	var err error
	db, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error trying to connect to the DB: ", err)
	}

	_, err = db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS teachingscholarship (value INTEGER, student TEXT, scholarship_type TEXT, scholarship_class TEXT, semester TEXT, professor TEXT)")
	if err != nil{
		fmt.Println(err)
		log.Fatal("Error during scholarship table creating: ", err)
	}

	_, err = db.Exec(context.Background(), "INSERT INTO teachingscholarship (value, student, scholarship_type, scholarship_class, semester, professor) VALUES (700, NULL, 'teaching', 'gdi', '2024.2', 'Robson'),(800, NULL, 'teaching', 'ess', '2024.2', 'Andre'), (650, NULL, 'teaching', 'plc', '2024.2', 'Andre');")
	if err != nil{
		fmt.Println(err)
		log.Fatal("Error during teaching scholarship table population: ", err)
	}

}

func GetUser(username string) []student.Student {
	rows, err := db.Query(context.Background(), "SELECT usr, pswrd, disciplines, deferral, enrolled FROM students WHERE usr = $1",username)
	if err != nil{
		log.Fatal("Student_SQL.go: It was not possible to realize the querry")
	}
	defer rows.Close()

	var users []student.Student
	for rows.Next() {
		var user student.Student
		var disciplines pq.StringArray
		err := rows.Scan(&user.User, &user.Password, &disciplines, &user.Deferral, &user.Enrolled, &user.Historic)
		if err != nil{
			fmt.Println(err)
			log.Fatal("Student_SQL.go: Error during row scan")
		}
		user.Disciplines = disciplines
		users = append(users, user)
	}
	return users
}