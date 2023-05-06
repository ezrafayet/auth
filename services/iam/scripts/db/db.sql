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

create table verification_codes
(
    user_id    uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    code       char(44)                 not null
);

create table authorization_codes
(
    user_id    uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    code       char(44)                 not null
);

create table refresh_tokens
(
    user_id    uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    token      char(44)                 not null
);

create table api_tokens
(
    owner_id   uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    name       varchar(100)             not null,
    token      char(44)                 not null,
    blocked    boolean default false    not null,
    revoked    boolean default false    not null,
    revoked_at timestamp with time zone
);

create table api_token_owners
(
    owner_id uuid not null,
    user_id  uuid not null
);