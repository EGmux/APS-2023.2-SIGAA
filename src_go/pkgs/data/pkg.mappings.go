package data

type Direction int

// ENUMS
const (
	SQLTOSTRUCT Direction = iota
	STRUCTTOSQL
)

// Abs value for int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Inplace, pass &, Convert a list of struct pointer to list of struct SQL, assuming both are equivalent
func ConvertSQLToList[U SQLTablesPtrs, V TablePtrs](
	listSrcType *[]*U,
	listDestType *[]*V,
	interval ...[]int,
) error {
	// NOTE: the subsequent functions expect empty listDestType
	if interval != nil {
		for _, pos := range interval {
			for _, tpos := range pos {
				_ = mapSQL_Struct(listSrcType, listDestType, tpos, SQLTOSTRUCT)
			}
		}
		return nil
	}
	for pos := range *listSrcType {
		_ = mapSQL_Struct(listSrcType, listDestType, pos, SQLTOSTRUCT)
	}
	return nil
}

// Inplace, pass &, Convert a list of  SQL  to list of struct, assuming both are equivalent
func ConvertStructToSQL[U SQLTablesPtrs, V TablePtrs](
	listSrcType *[]*U,
	listDestType *[]*V,
	interval ...[]int,
) error {
	if interval != nil {
		for _, pos := range interval {
			for _, index := range pos {
				_ = mapSQL_Struct(listSrcType, listDestType, index, STRUCTTOSQL)
			}
		}
		return nil
	}
	for pos := range *listSrcType {
		_ = mapSQL_Struct(listSrcType, listDestType, pos, STRUCTTOSQL)
	}
	return nil
}

// map SQL to struct or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mapSQL_Struct[U SQLTablesPtrs, V TablePtrs](
	s *[]*U,
	d *[]*V,
	pos int,
	dir Direction,
) error {
	switch t1 := any(s).(type) {
	case *[]*CredentialsSQL:
		t2 := any(d).(*[]*Credentials)
		AddRows(t2)
		_ = mappingCredentials((*t2)[pos], (*t1)[pos], dir)
	case *[]*ProfessorSQL:
		t2 := any(d).(*[]*Professor)
		AddRows(t2)
		_ = mappingProfessor((*t2)[pos], (*t1)[pos], dir)
	case *[]*PROAESSQL:
		t2 := any(d).(*[]*PROAES)
		AddRows(t2)
		_ = mappingPROAES((*t2)[pos], (*t1)[pos], dir)
	case *[]*EnrollmentSQL:
		t2 := any(d).(*[]*Enrollment)
		AddRows(t2)
		_ = mappingEnrollment((*t2)[pos], (*t1)[pos], dir)
	case *[]*TeachingScholarshipSQL:
		t2 := any(d).(*[]*TeachingScholarship)
		AddRows(t2)
		_ = mappingTeachingScholarship((*t2)[pos], (*t1)[pos], dir)
	}
	return nil
}

// map Credentials struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingCredentials(d *Credentials, s *CredentialsSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Password = s.Passwd
		d.User = s.Username
		d.Logged = s.Logged
		d.Id = s.Id
	} else if dir == STRUCTTOSQL {
		s.Passwd = d.Password
		s.Username = d.User
		s.Logged = d.Logged
		s.Id = d.Id
	}
	return nil
}

// map Professor struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingProfessor(d *Professor, s *ProfessorSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Id = s.Id
		d.Name = s.Name
		d.Email = s.Email
		d.CPF = s.CPF
	} else if dir == STRUCTTOSQL {
		s.Id = d.Id
		s.Name = d.Name
		s.Email = d.Email
		s.CPF = d.CPF
	}
	return nil
}

// map PROAES struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingPROAES(d *PROAES, s *PROAESSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Email = s.Email
		d.Id = s.Id
	} else if dir == STRUCTTOSQL {
		s.Email = d.Email
		s.Id = d.Id
	}
	return nil
}

// map Enrollment struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingEnrollment(d *Enrollment, s *EnrollmentSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Id = s.Id
		d.CPF = s.CPF
		d.Name = s.Name
		for p, e := range s.Transcript {
			mappingClass(d.Transcript[p], e, dir)
		}
	} else if dir == STRUCTTOSQL {
		s.Id = d.Id
		s.CPF = d.CPF
		s.Name = d.Name
		for p, e := range s.Transcript {
			mappingClass(d.Transcript[p], e, STRUCTTOSQL)
		}
	}
	return nil
}

// map Student struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingStudent(d *Student, s *StudentSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Name = s.Name
		d.CPF = s.CPF
	} else if dir == STRUCTTOSQL {
		s.Name = d.Name
		s.CPF = d.CPF
	}
	return nil
}

// map Class struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingClass(d *Class, s *ClassSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Id = s.Id
		d.Assesment = s.Assesment
		d.Timetable = s.Timetable
		d.Mandatory = s.Required
		d.Capacity = s.Capacity
		for p, e := range s.Professor {
			mappingProfessor(d.Professor[p], e, dir)
		}
	} else if dir == STRUCTTOSQL {
		s.Id = d.Id
		s.Assesment = d.Assesment
		s.Timetable = d.Timetable
		s.Required = d.Mandatory
		s.Capacity = d.Capacity
		for p, e := range s.Professor {
			mappingProfessor(d.Professor[p], e, STRUCTTOSQL)
		}
	}
	return nil
}

// map Grades struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingGrades(d *Grade, s *GradeSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Grade = s.Grade
		d.Id = s.Id
	} else if dir == STRUCTTOSQL {
		s.Grade = d.Grade
		s.Id = d.Id
	}
	return nil
}

// map Courses struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingCourses(d *Course, s *CourseSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		for p, e := range s.Classes {
			mappingClass(d.Classes[p], e, dir)
		}
		d.Id = s.Id
		d.Department = s.Department
		d.Curriculum = s.Curriculum
		d.Name = s.Name
	} else if dir == STRUCTTOSQL {
		for p, e := range s.Classes {
			mappingClass(d.Classes[p], e, STRUCTTOSQL)
		}
		s.Id = d.Id
		s.Department = d.Department
		s.Curriculum = d.Curriculum
		s.Name = d.Name
	}
	return nil
}

// map Scholarship struct to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingScholarship(d *Scholarship, s *ScholarshipSQL, dir Direction) error {
	if dir == SQLTOSTRUCT {
		d.Id = s.Id
		d.Value = s.Value
		mappingStudent(&d.Student, &s.Student, dir)
	} else if dir == STRUCTTOSQL {
		s.Id = d.Id
		s.Value = d.Value
		mappingStudent(&d.Student, &s.Student, STRUCTTOSQL)
	}
	return nil
}

// map mappingTeachingScholarship to SQL or vice-versa according to dir ENUM
// STRUCTTOSQL and STQLTOSTRUCT
func mappingTeachingScholarship(
	d *TeachingScholarship,
	s *TeachingScholarshipSQL,
	dir Direction,
) error {
	if dir == SQLTOSTRUCT {
		mappingScholarship(&d.Scholarship, &s.Scholarship, dir)
		mappingClass(&d.Class, &s.Class, dir)
		d.Id = s.Id
		mappingProfessor(&d.Professor, &s.Professor, dir)
		d.Semester = s.Semester
	} else if dir == STRUCTTOSQL {
		mappingScholarship(&d.Scholarship, &s.Scholarship, STRUCTTOSQL)
		mappingClass(&d.Class, &s.Class, STRUCTTOSQL)
		s.Id = d.Id
		mappingProfessor(&d.Professor, &s.Professor, STRUCTTOSQL)
		s.Semester = d.Semester
	}
	return nil
}
