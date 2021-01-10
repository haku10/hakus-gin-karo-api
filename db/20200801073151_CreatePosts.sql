-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id int(10) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    age int(3),
    gender int(1),
    create_at datetime,
    update_at datetime,
    PRIMARY KEY(id)
);
