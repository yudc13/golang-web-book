CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE NOT NULL,
    pwd VARCHAR(36) NOT NULL,
    email VARCHAR(100) NOT NULL
)

create table books (
    id int primary key auto_increment,
    name varchar(100) not null,
    author varchar(50) not null,
    price double(11, 2) not null,
    sales int not null,
    stock int not null,
    cover varchar(100)
)