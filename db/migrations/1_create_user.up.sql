-- +migrate Up
CREATE TABLE users (
    id serial NOT NULL,
    firstname  text NOT NULL,
    lastname  text NOT NULL,
    PRIMARY KEY (id)
);
