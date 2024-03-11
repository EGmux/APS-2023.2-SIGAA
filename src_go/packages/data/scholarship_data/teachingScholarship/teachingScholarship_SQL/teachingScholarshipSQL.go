package teachingscholarshipsql

import (
	"context"
	"fmt"
	"log"
	

	"github.com/jackc/pgx/v4/pgxpool"
	teachingscholarship "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship"
	singleton_db "sigaa.ufpe/packages/utils/singleton"
)

var db *pgxpool.Pool

func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship {
	db = singleton_db.Refer_DB()
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

func GetAllTeachingScholarships() []teachingscholarship.TeachingScholarship {
	db = singleton_db.Refer_DB()
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


func Get_Teaching_Scholarship_By_Id(Scholarship_Id string) []teachingscholarship.TeachingScholarship{
	db = singleton_db.Refer_DB()
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
	db = singleton_db.Refer_DB()

	_, err := db.Exec(context.Background(), "UPDATE teachingscholarship SET student = $1 WHERE identifier = $2", Student_Name, Scholarship_Id)

	if err != nil{
		log.Fatal("Error while handling the UPDATE Querry: ")
	}
}