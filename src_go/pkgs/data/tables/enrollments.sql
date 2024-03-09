create table enrollments(
  idenrollment serial primary key,
  name text, 
  status bool,
  cpf text not null,
  /*transcript classes_table*/
)

