-- +goose Up
create table auth (
                         login text primary key,
                         hashed_password text not null,
                         role text not null
);

-- +goose Down
drop table auth;