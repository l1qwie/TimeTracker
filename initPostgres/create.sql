CREATE SEQUENCE client_id_seq
    INCREMENT 1
    START 1;

CREATE TABLE Clients (
    clientid int PRIMARY KEY,
    name varchar(255) DEFAULT '',
    surname varchar(255) DEFAULT '',
    patronymic varchar(255) DEFAULT '',
    age int DEFAULT 0,
    passportnumber varchar(255) DEFAULT '',
    passportseries varchar(255) DEFAULT '',
    address varchar(255) DEFAULT ''
)