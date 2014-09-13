CREATE Table mycomputer;

USE mycomputer;

CREATE TABLE IF NOT EXISTS user (
    name varchar(256) PRIMARY KEY,
    email varchar(256),
    description varchar(256)
);

CREATE TABLE IF NOT EXISTS item (
   number int PRIMARY KEY,
   image varchar(256) NOT NULL,
   description varchar(256),
   username varchar(256) REFERENCES user(name)
);

CREATE TABLE IF NOT EXISTS comment (
       id int PRIMARY KEY AUTO_INCREMENT,
       content varchar(256)
)
