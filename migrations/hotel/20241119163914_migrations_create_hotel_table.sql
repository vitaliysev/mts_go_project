-- +goose Up
CREATE TABLE hotels
(
    id         SERIAL PRIMARY KEY,
    hotel_name VARCHAR(255) NOT NULL,
    location   VARCHAR(255) NOT NULL,
    price      INT          NOT NULL,
    username   text         NOT NULL
);
-- +goose Down
DROP TABLE  hotels;
