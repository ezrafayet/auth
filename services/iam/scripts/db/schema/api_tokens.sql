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