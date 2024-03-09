create table teaches(
    idclass integer references professors(idprofessor),
    idprofessor integer references classes(idclass),
    constraint teaches_pk PRIMARY KEY(idclass, idprofessor) );
);

