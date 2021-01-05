CREATE TABLE users (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       pwd VARCHAR(36) NOT NULL,
                       email VARCHAR(100) NOT NULL
)