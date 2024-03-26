package scholarshipsql

import (
	"context"
	"fmt"
	"log"
	"os"
	"scholarship/utils/student"
	"scholarship/utils/teachingscholarship"

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
		err := rows.Scan(&user.User, &user.Password, &disciplines, &user.Deferral, &user.Enrolled)
		if err != nil{
			fmt.Println(err)
			log.Fatal("Student_SQL.go: Error during row scan")
		}
		user.Disciplines = disciplines
		users = append(users, user)
	}
	return users
}

func GetAllTeachingScholarships() []teachingscholarship.TeachingScholarship {

	rows, err := db.Query(context.Background(), "SELECT identifier, value, scholarship_class, semester, professor FROM teachingscholarship")
	if err != nil {
		log.Fatal("TeachingScholarship.go : It was not possible to realize the querry")
	}
	defer rows.Close()
	var teachingScholarships []teachingscholarship.TeachingScholarship
	for rows.Next() {
		var teachingScholarship teachingscholarship.TeachingScholarship
		err := rows.Scan(&teachingScholarship.Scholarship.Id,&teachingScholarship.Scholarship.Value, &teachingScholarship.Class.Name, &teachingScholarship.Semester, &teachingScholarship.Professor.Name)
		if err != nil {
			log.Fatal("TeachingScholarship.go: Erros during row scan")
		}
		teachingScholarships = append(teachingScholarships, teachingScholarship)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("TeachingScholarship.go: Error during row scan")
	}

	fmt.Println(teachingScholarships)
	return teachingScholarships
}

func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship {

	rows, err := db.Query(context.Background(), "SELECT identifier, value, scholarship_class, semester, professor FROM teachingscholarship WHERE student IS NULL")
	if err != nil {
		log.Fatal("TeachingScholarship.go : It was not possible to realize the querry")
	}
	defer rows.Close()
	var teachingScholarships []teachingscholarship.TeachingScholarship
	for rows.Next() {
		var teachingScholarship teachingscholarship.TeachingScholarship
		err := rows.Scan(&teachingScholarship.Scholarship.Id,&teachingScholarship.Scholarship.Value, &teachingScholarship.Class.Name, &teachingScholarship.Semester, &teachingScholarship.Professor.Name)
		if err != nil {
			log.Fatal("TeachingScholarship.go: Erros during row scan")
		}
		teachingScholarships = append(teachingScholarships, teachingScholarship)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("TeachingScholarship.go: Error during row scan")
	}

	fmt.Println(teachingScholarships)
	return teachingScholarships
}

func Get_Teaching_Scholarship_By_Id(Scholarship_Id string) []teachingscholarship.TeachingScholarship{

	var teachingScholarships []teachingscholarship.TeachingScholarship
	fmt.Println(Scholarship_Id)
	row, err := db.Query(context.Background(), "SELECT identifier, value, scholarship_class, semester, professor FROM teachingscholarship WHERE identifier = $1",Scholarship_Id)
	fmt.Println(row.RawValues())
	if err != nil{
		log.Fatal("Error while handling the Query")
	}
	
	for row.Next() {
		var teachingScholarship teachingscholarship.TeachingScholarship
		err := row.Scan(&teachingScholarship.Scholarship.Id,&teachingScholarship.Scholarship.Value, &teachingScholarship.Class.Name, &teachingScholarship.Semester, &teachingScholarship.Professor.Name)
		if err != nil {
			log.Fatal("TeachingScholarship.go: Erros during row scan")
		}
		teachingScholarships = append(teachingScholarships, teachingScholarship)
	}

	fmt.Println(teachingScholarships)

	return teachingScholarships

}


func Associate_Student_To_Teaching_Scholarship(Scholarship_Id int, Student_Name string) {

	_, err := db.Exec(context.Background(), "UPDATE teachingscholarship SET student = $1 WHERE identifier = $2", Student_Name, Scholarship_Id)

	if err != nil{
		log.Fatal("Error while handling the UPDATE Querry: ")
	}
}