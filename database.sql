/* Make MySQL support utf-8
root@ emacs /etc/mysql/my.cnf
[client]
default-character-set=utf8
[mysqld]
init_connect=’SET collation_connection = utf8_unicode_ci’
init_connect=’SET NAMES utf8′
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
*/

CREATE DATABASE mycomputer;

USE mycomputer;

CREATE TABLE IF NOT EXISTS user (
    name varchar(255) PRIMARY KEY,
    email varchar(255),
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS item (
   number int PRIMARY KEY,
   image varchar(255),
   description varchar(255),
   username varchar(255) REFERENCES user(name)
);

CREATE TABLE IF NOT EXISTS comment (
  id int PRIMARY KEY AUTO_INCREMENT,
  content varchar(255)
)
