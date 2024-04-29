create table if not exists users(
    id serial primary key,
    firstname character varying not null,
    lastname character varying not null,
    email character varying unique not null,
    hashed_password character varying not null,
    created_at timestamp not null
);