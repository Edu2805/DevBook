CREATE DATABASE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id INTEGER generated always as IDENTITY primary key,
    user_name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    user_password varchar(255) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
