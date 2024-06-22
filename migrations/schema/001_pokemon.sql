-- +goose Up
CREATE TABLE pokemon (
    id INT NOT NULL UNIQUE PRIMARY KEY ,
    name TEXT NOT NULL ,
    height INT NOT NULL ,
    weight INT NOT NULL ,
    base_experience INT NOT NULL
);

-- +goose Down
DROP TABLE pokemon;