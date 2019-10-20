create table "user"
(
    id         serial not null
        constraint user_pk
            primary key,
    name       varchar(64),
    email      varchar(255),
    password   varchar(255),
    salt       varchar(255),
    enabled    boolean default false,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table "user-verification"
(
    id     serial not null
        constraint user_verification_pk
            primary key,
    userId integer not null
        references "user",
    code   varchar(64),
    activated boolean default false
);