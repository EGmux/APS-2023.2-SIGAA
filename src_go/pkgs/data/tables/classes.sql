create table classes(
  idclass serial primary key,
  capacity int,
  required int not null,
  timetable text, 
  assessments text[]
)

