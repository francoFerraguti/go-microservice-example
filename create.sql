CREATE TABLE IF NOT EXISTS USERS (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    enabled tinyint(1) NOT NULL DEFAULT 1,
    dateCreated TIMESTAMP
);