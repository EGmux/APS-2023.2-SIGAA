package data

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
