package gradessql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	grades "sigaa.ufpe/packages/data/grades_data"
	singleton_db "sigaa.ufpe/packages/utils/singleton"
)

var db *pgxpool.Pool

func Get_All_Grades_From_Student(username string) []grades.Grade{

	db = singleton_db.Refer_DB()

	query := "SELECT class, assignment, student, grade FROM grades WHERE student='"+username+"'"

	rows, err := db.Query(context.Background(), query)
	if err != nil{
		fmt.Println(query)
		fmt.Println(err)
		log.Fatal("grades_SQL.go: Error during SELECT query")
	}

	defer rows.Close()

	var gs []grades.Grade
	var g grades.Grade

	for rows.Next(){
		var grd pq.StringArray

		err := rows.Scan(&g.Class, &g.Assignment, &g.Student, &grd)
		if err != nil{
			fmt.Println(err)
			log.Fatal("grades_SQL.go: Error during ROW SCAN")
		}
		g.Grade = grd
		gs = append(gs, g)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		log.Fatal("grades_SQL.go: Error during row scan")
	}

	return gs
}