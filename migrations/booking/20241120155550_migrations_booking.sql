-- +goose Up
create table booking (
                         id serial primary key,
                         period_use text not null,
                         created_at timestamp not null default now(),
                         updated_at timestamp,
                         hotel_id serial,
                         username text not null
);

-- +goose Down
drop table booking;