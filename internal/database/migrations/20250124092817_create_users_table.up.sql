CREATE TABLE users (
    id bigserial primary key,
    email varchar(255) not null constraint users_email_unique unique,
    password varchar(255) not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    deleted_at timestamp(0)
);

alter table users owner to postgres;
