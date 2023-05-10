create table refresh_tokens
(
    user_id    uuid      not null,
    created_at timestamp not null,
    expires_at timestamp not null,
    token      char(44)  not null
);