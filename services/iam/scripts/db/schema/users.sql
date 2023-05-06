create table users
(
    id                   uuid                     not null,
    created_at           timestamp with time zone not null,
    username             varchar(50)              not null,
    username_fingerprint bytea                    not null,
    email                varchar(100)             not null,
    email_verified       boolean default false    not null,
    email_verified_at    timestamp with time zone,
    blocked              boolean default false    not null,
    deleted              boolean default false    not null,
    deleted_at           timestamp with time zone
);