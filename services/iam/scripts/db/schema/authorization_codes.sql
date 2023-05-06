create table authorization_codes
(
    user_id    uuid                     not null,
    created_at timestamp with time zone not null,
    expires_at timestamp with time zone not null,
    code       char(44)                 not null
);