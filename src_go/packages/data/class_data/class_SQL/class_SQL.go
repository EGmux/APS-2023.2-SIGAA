package data_classdata_classsql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	class "sigaa.ufpe/packages/data/class_data"
	singleton_db "sigaa.ufpe/packages/utils/singleton"
)

var db *pgxpool.Pool

func GetAllClasses() []class.Class {
	db = singleton_db.Refer_DB()
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
	db = singleton_db.Refer_DB()

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