CREATE SEQUENCE client_id_seq
    INCREMENT 1
    START 1;

CREATE SEQUENCE task_id_seq
    INCREMENT 1
    START 1;

CREATE TABLE Clients (
    clientid int PRIMARY KEY,
    name varchar(255) DEFAULT '',
    surname varchar(255) DEFAULT '',
    patronymic varchar(255) DEFAULT '',
    age int DEFAULT 0,
    passportseries varchar(255) DEFAULT '',
    passportnumber varchar(255) DEFAULT '',
    address varchar(255) DEFAULT ''
);

CREATE TABLE Tasks (
    taskid int PRIMARY KEY,
    clientid int,
    taskname varchar(255) DEFAULT '',
    tasktimestart timestamp,
    tasktimeend timestamp,
    FOREIGN KEY (clientid) REFERENCES Clients(clientid)
);