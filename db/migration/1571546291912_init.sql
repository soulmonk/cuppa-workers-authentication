create table "user"
(
  id         serial                  not null
    constraint user_pk
      primary key,
  name       varchar(64)             not null, -- todo unique
  email      varchar(256)            not null,
  password   varchar(256)            not null,
--     salt       varchar(256),
  enabled    boolean   default false not null,
  created_at timestamp default now() not null,
  updated_at timestamp default now() not null
);

create table "user-verification"
(
  id        serial  not null
    constraint user_verification_pk
      primary key,
  user_id   integer not null
    references "user",
  code      varchar(64),
  activated boolean default false
);

-- insert into "user" (name, email, password, salt, enabled, created_at, updated_at)
-- values ('admin', 'admin@example.com', '', '', false, now(), now());
