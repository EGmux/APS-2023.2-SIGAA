create table transcript(
    idclass integer references grades(idgrade),
    idgrade integer references classes(idclass),
    constraint teaches_pk PRIMARY KEY(idclass, idgrade) );
);
