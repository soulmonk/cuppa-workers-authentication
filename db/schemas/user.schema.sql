create table "user"
(
  id            BIGSERIAL
    constraint user_pk
      primary key,
  name          varchar(64)             not null,
  email         varchar(256)            not null,
  password      varchar(256)            not null,
  enabled       boolean   default false not null,
  created_at    timestamp default now() not null,
  updated_at    timestamp default now() not null,
  refresh_token uuid      default null
);
