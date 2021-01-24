
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id int(10) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    age int(3),
    gender int(1),
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users
