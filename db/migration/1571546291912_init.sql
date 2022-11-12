create table "user"
(
  id         BIGSERIAL               not null
    constraint user_pk
      primary key,
  name       varchar(64)             not null,
  email      varchar(256)            not null,
  password   varchar(256)            not null,
  enabled    boolean   default false not null,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null
);

create table "user_verification"
(
  id        BIGSERIAL not null
    constraint user_verification_pk
      primary key,
  user_id   integer   not null
    references "user",
  code      varchar(64),
  activated boolean default false
);
