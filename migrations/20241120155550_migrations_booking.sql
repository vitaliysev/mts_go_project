-- +goose Up
create table booking (
                      id serial primary key,
                      title text not null,
                      period_use text not null,
                      created_at timestamp not null default now(),
                      updated_at timestamp
);

-- +goose Down
drop table booking;