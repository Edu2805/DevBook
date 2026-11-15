CREATE DATABASE devbook;

CREATE TABLE IF NOT EXISTS users(
    id INTEGER generated always as IDENTITY primary key,
    user_name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    user_password varchar(255) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS followers (
    user_id INTEGER NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    follower_id INTEGER NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id)
);
