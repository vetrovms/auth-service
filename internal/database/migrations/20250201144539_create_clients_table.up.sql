CREATE TABLE clients (
    id bigserial primary key,
    client_id varchar(255) not null constraint users_client_id_unique unique,
    client_secret varchar(255) not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    deleted_at timestamp(0)
);
