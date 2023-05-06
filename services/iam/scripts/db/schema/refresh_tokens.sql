create table refresh_tokens
(
    user_id    uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    token      char(44)                 not null
);