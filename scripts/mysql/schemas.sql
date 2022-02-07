CREATE DATABASE IF NOT EXISTS users_db CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS users_db.users (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    internal_code BIGINT(20) NOT NULL UNIQUE,
    first_name VARCHAR(45) NULL,
    last_name VARCHAR(45) NULL,
    email VARCHAR(45) NOT NULL,
    status VARCHAR(45) NOT NULL,
    password VARCHAR(36) NOT NULL,
    date_created DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY(id),
    UNIQUE INDEX email_UNIQUE (email ASC));
