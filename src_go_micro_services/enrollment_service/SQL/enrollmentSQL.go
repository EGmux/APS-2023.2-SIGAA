package sql

import (
	"context"
	"fmt"
	"log"
	"main/utils/class"
	"main/utils/student"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
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


func GetAllClasses() []class.Class {

	rows, err := db.Query(context.Background(), "SELECT name, capacity, mandatory, equivalency, timetable, assesment, professor, semester FROM classes")
	if err != nil {
		log.Fatal("class_SQL.go: It was not possible to realize the querry")
	}
	defer rows.Close()

	var classes []class.Class
	for rows.Next() {
		var class class.Class
		var equivalency pq.StringArray
		var assesment pq.StringArray
		err := rows.Scan(&class.Name, &class.Capacity, &class.Mandatory, &equivalency, &class.Timetable, &assesment, &class.Professor.Name, &class.Semester)
		if err != nil {
			fmt.Println(err)
			log.Fatal("class_SQL.go: Error guring row scan")
		}
		class.Equivalency = equivalency
		class.Assesment = assesment
		classes = append(classes, class)
	}
	fmt.Println(classes)
	return classes
}


func GetSelectedClasses(selectedClasses []string) []class.Class{

	query := "SELECT name, capacity, mandatory, equivalency, timetable, assesment, professor, semester FROM classes WHERE name IN ("
	for i, class := range selectedClasses{

		query += "'"+class+"'"
		if i < len(selectedClasses)-1{
			query += ", "
		}
	}
	query += ")"
	fmt.Println(query)
	rows, err := db.Query(context.Background(), query)
	if err != nil{
		fmt.Println(err)
		log.Fatal("class_SQL.go: Error during query")
	}

	defer rows.Close()

	var classes []class.Class
	for rows.Next() {
		var class class.Class
		var equivalency pq.StringArray
		var assesment pq.StringArray
		err := rows.Scan(&class.Name, &class.Capacity, &class.Mandatory, &equivalency, &class.Timetable, &assesment, &class.Professor.Name, &class.Semester)
		if err != nil {
			fmt.Println(err)
			log.Fatal("class_SQL.go: Error guring row scan")
		}
		class.Equivalency = equivalency
		class.Assesment = assesment
		classes = append(classes, class)
	}
	fmt.Println(classes)
	return classes
}

func Update_Student_Enrollment(username string, classes []class.Class){
	
	query := "Update students SET "

	var class_enrolled pq.StringArray

	query += "disciplines = ARRAY["
	for i, class := range classes{

		fmt.Println(class)
		query += "'" +class.Name +"'"
		if i < len(classes)-1{
			query += ", "
		}

	} 
	query += "] "

	query += ", enrolled = true"

	fmt.Println(class_enrolled)
	query += " WHERE usr ='" +username + "'"
	fmt.Println(query)
	
	_, err := db.Exec(context.Background(), query)
	if err != nil{
		fmt.Println(err)
		log.Fatal("Student_SQL.go: Error during Exec query")
	}

	Insert_Student_Grades(username, classes)

}

func Insert_Student_Grades(username string, classes []class.Class){

	for _, class := range(classes){

		query := "INSERT INTO grades (class, assignment, student, grade) VALUES ('"+class.Name+"',"

		for _, assignment := range(class.Assesment){
			query_2 := ""
			query_2 += "'"+assignment+"','"+ username+"',"+"NULL);"

			_,err := db.Exec(context.Background(), query+query_2)
			if err != nil{
				fmt.Println(query)
				fmt.Println(err)
				log.Fatal("student_SQL.go: Error during Insert into Grades")
			}
		}
	}

}

func Update_Student_Deferral(username string){

	query := "UPDATE students  SET deferral=true , enrolled=false , disciplines=NULL WHERE usr='"+username+"'"
	_, err := db.Exec(context.Background(), query)
	if err != nil{
		fmt.Println(err)
		log.Fatal("student_SQL.go: Error duting UPDATE Query")
	}

}

func Retake_Student_Enrollment(username string){

	query := "UPDATE students SET deferral=false, enrolled=false WHERE usr='"+username+"'"
	_, err := db.Exec(context.Background(), query)
	if err != nil{
		fmt.Println(err)
		log.Fatal("student_SQL.go: Error during UPDATE Query")
	}
}