CREATE DATABASE simple_user_db;
USE simple_user_db;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);

insert into users (name, age) values ('Matthew', 22);
insert into users (name, age) values ('Albert', 27);
insert into users (name, age) values ('Daniel', 32);
insert into users (name, age) values ('Criag', 28);