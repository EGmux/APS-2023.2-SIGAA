create table curricuilum(
    idcourse integer references class(idclass),
    idclass integer references course(idcourse),
    constraint curriculum_pk PRIMARY KEY(idcourse, idclass) );
);

