package class

import (
	grades "sigaa.ufpe/packages/data/grades_data"
	professor "sigaa.ufpe/packages/data/professor_data"
	student "sigaa.ufpe/packages/data/students_data"
)

type Class struct {
	Capacity    int
	Mandatory   bool
	Equivalency []Class
	Timetable   string
	Assesment   string
	Professor   professor.Professor
	Students    []student.Student
	Grades      []grades.Grade
}

