package class

import (
	professor "sigaa.ufpe/packages/data/professor_data"
)

type Class struct {
	Name        string
	Capacity    int
	Mandatory   bool
	Equivalency []string
	Timetable   string
	Assesment   []string
	Professor   professor.Professor
	Students    *string
	Grades      string
	Semester    string
}
